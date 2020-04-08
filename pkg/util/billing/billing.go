package billing

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	azstorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/storage"
	"github.com/Azure/ARO-RP/pkg/util/feature"
)

const (
	tenantIDMSFT string = "72f988bf-86f1-41af-91ab-2d7cd011db47"
	tenantIDAME  string = "33e01921-4d64-4f8c-a055-5bdaffd5e33d"
)

type Manager interface {
	Create(context.Context, *api.OpenShiftClusterDocument) error
	Delete(context.Context, string) error
}

type manager struct {
	storageClient *azstorage.Client
	billingDB     database.Billing
	subDB         database.Subscriptions
	log           *logrus.Entry
}

func NewManager(_env env.Interface, billing database.Billing, sub database.Subscriptions, log *logrus.Entry) (Manager, error) {
	fpAuthorizer, err := _env.FPAuthorizer(_env.TenantID(), azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		return nil, err
	}

	e2estorage := storage.NewAccountsClient(_env.E2EStorageAccountSubID(), fpAuthorizer)
	if err != nil {
		return nil, err
	}

	keys, err := e2estorage.ListKeys(context.Background(), "billing-global", _env.E2EStorageAccountName(), "")
	if err != nil {
		return nil, err
	}
	key := *(*keys.Keys)[0].Value

	client, err := azstorage.NewBasicClient(_env.E2EStorageAccountName(), key)
	if err != nil {
		return nil, err
	}
	return &manager{
		storageClient: &client,
		subDB:         sub,
		billingDB:     billing,
		log:           log,
	}, nil
}

func (m *manager) Create(ctx context.Context, doc *api.OpenShiftClusterDocument) error {
	billingDoc, err := m.billingDB.Create(ctx, &api.BillingDocument{
		ID:                        doc.ID,
		Key:                       doc.Key,
		ClusterResourceGroupIDKey: doc.ClusterResourceGroupIDKey,
		Billing: &api.Billing{
			TenantID: doc.OpenShiftCluster.Properties.ServicePrincipalProfile.TenantID,
			Location: doc.OpenShiftCluster.Location,
		},
	})
	if err != nil {
		if err, ok := err.(*cosmosdb.Error); ok && err.StatusCode == http.StatusConflict {
			m.log.Printf("billing record already present id DB")
			return nil
		}
		return err
	}

	if e2eErr := m.createOrUpdateE2EBlob(ctx, billingDoc); e2eErr != nil {
		m.log.Warnf("CreateOrUpdateE2EBlob failed : %s", e2eErr.Error())
	}

	return nil
}

func (m *manager) Delete(ctx context.Context, id string) error {
	m.log.Printf("updating billing record with deletion time")
	billingDoc, err := m.billingDB.MarkForDeletion(ctx, id)
	if cosmosdb.IsErrorStatusCode(err, http.StatusNotFound) {
		return nil
	}

	if e2eErr := m.createOrUpdateE2EBlob(ctx, billingDoc); e2eErr != nil {
		// We are not failing the operation if we cannot write to e2e storage account, just warning
		m.log.Warnf("CreateOrUpdateE2EBlob failed : %s", e2eErr.Error())
	}

	return nil
}

// IsSubscriptionRegisteredToE2E returns true if the subscription is having the
// "Microsoft.ContainerService/SaveAROTestConfig" feature registered
func IsSubscriptionRegisteredToE2E(sub *api.SubscriptionProperties) bool {
	if sub.TenantID == tenantIDMSFT || sub.TenantID == tenantIDAME {
		return feature.IsRegisteredForFeature(sub, api.FeatureSaveAROTestConfig)
	}
	return false
}

// createOrUpdateE2Eblob create a copy of the billing document into the e2e storage account.
// This is used later on for E2E Billing
func (m *manager) createOrUpdateE2EBlob(ctx context.Context, doc *api.BillingDocument) error {
	// Validate if E2E Feature is registred
	resource, err := azure.ParseResourceID(doc.Key)
	if err != nil {
		return err
	}

	subDocument, err := m.subDB.Get(ctx, resource.SubscriptionID)
	if err != nil {
		return err
	}

	if !IsSubscriptionRegisteredToE2E(subDocument.Subscription.Properties) {
		return nil
	}

	blobclient := m.storageClient.GetBlobService()

	containerName := strings.ToLower("bill-" + doc.Billing.Location + "-" + resource.ResourceGroup + "-" + resource.ResourceName)
	if len(containerName) > 63 {
		containerName = containerName[:62] // maximum characters allowed is 63
	}

	containerRef := blobclient.GetContainerReference(containerName)
	_, err = containerRef.CreateIfNotExists(nil)
	if err != nil {
		return fmt.Errorf("Error creating container : %s", err.Error())
	}

	blobRef := containerRef.GetBlobReference("billingentity")
	b, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	return blobRef.CreateBlockBlobFromReader(bytes.NewReader(b), nil)
}
