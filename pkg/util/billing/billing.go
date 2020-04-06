package billing

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	azstorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/storage"
	"github.com/Azure/ARO-RP/pkg/util/feature"
)

const (
	tenantIDMSFT string = "72f988bf-86f1-41af-91ab-2d7cd011db47"
	tenantIDAME  string = "33e01921-4d64-4f8c-a055-5bdaffd5e33d"
)

type E2EManager interface {
	CreateOrUpdateE2EBlob(ctx context.Context, env env.Interface, sub database.Subscriptions, billingDoc *api.BillingDocument) error
}

type e2eManager struct {
	storageClient *azstorage.Client
}

func NewE2EManager(ctx context.Context, env env.Interface) (E2EManager, error) {
	fpAuthorizer, err := env.FPAuthorizer(env.TenantID(), azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		return nil, err
	}

	e2estorage := storage.NewAccountsClient(env.E2EStorageAccountSubID(), fpAuthorizer)
	if err != nil {
		return nil, err
	}

	keys, err := e2estorage.ListKeys(ctx, "billing-global", env.E2EStorageAccountName(), "")
	if err != nil {
		return nil, err
	}
	key := *(*keys.Keys)[0].Value

	client, err := azstorage.NewBasicClient(env.E2EStorageAccountName(), key)
	if err != nil {
		return nil, err
	}
	return &e2eManager{
		storageClient: &client,
	}, nil
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
func (m *e2eManager) createOrUpdateE2Eblob(ctx context.Context, env env.Interface, doc api.BillingDocument, resource azure.Resource) error {
	blobclient := m.storageClient.GetBlobService()

	containerName := strings.ToLower("bill-" + doc.Billing.Location + "-" + resource.ResourceGroup + "-" + resource.ResourceName)
	if len(containerName) > 63 {
		containerName = containerName[:62] // maximum characters allowed is 63
	}

	containerRef := blobclient.GetContainerReference(containerName)
	_, err := containerRef.CreateIfNotExists(nil)
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

// CreateOrUpdateE2EBlob determines if the document belongs to an e2e subscription.
// If so create a copy of the billing document into the e2e storage account.
// This is used later on for E2E Billing
func (m *e2eManager) CreateOrUpdateE2EBlob(ctx context.Context, env env.Interface, sub database.Subscriptions, billingDoc *api.BillingDocument) error {
	// Validate if E2E Feature is registred
	resource, err := azure.ParseResourceID(billingDoc.Key)
	if err != nil {
		return err
	}

	subDocument, err := sub.Get(ctx, resource.SubscriptionID)
	if err != nil {
		return err
	}

	if IsSubscriptionRegisteredToE2E(subDocument.Subscription.Properties) {
		errBlob := m.createOrUpdateE2Eblob(ctx, env, *billingDoc, resource)
		if errBlob != nil {
			return errBlob
		}
	}

	return nil
}
