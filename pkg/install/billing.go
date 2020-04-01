package install

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

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	"github.com/Azure/ARO-RP/pkg/env"
)

const (
	tenantIDMSFT string = "72f988bf-86f1-41af-91ab-2d7cd011db47"
	tenantIDAME  string = "33e01921-4d64-4f8c-a055-5bdaffd5e33d"
)

func (i *Installer) createBillingRecord(ctx context.Context) error {
	billingDoc, err := i.billing.Create(ctx, &api.BillingDocument{
		ID:                        i.doc.ID,
		Key:                       i.doc.Key,
		ClusterResourceGroupIDKey: i.doc.ClusterResourceGroupIDKey,
		Billing: &api.Billing{
			TenantID: i.doc.OpenShiftCluster.Properties.ServicePrincipalProfile.TenantID,
			Location: i.doc.OpenShiftCluster.Location,
		},
	})
	if err != nil {
		if err, ok := err.(*cosmosdb.Error); ok && err.StatusCode == http.StatusConflict {
			i.log.Printf("billing record already present id DB")
			return nil
		}
		return err
	}
	// Validate if E2E Feature is registred
	resource, err := azure.ParseResourceID(billingDoc.Key)
	if err != nil {
		return err
	}

	subDocument, err := i.sub.Get(ctx, resource.SubscriptionID)
	if err != nil {
		return err
	}

	if isSubscriptionRegisteredToE2E(subDocument.Subscription.Properties) {
		// We are not failing the operation if we cannot write to e2e storage account, just warning
		// The blob can only be created when the billing record is also written for the first time
		errBlob := createE2Eblob(ctx, i.env, *billingDoc, resource)
		if errBlob != nil {
			i.log.Warnf("createE2Eblob failed : %s", errBlob.Error())
		}
	}
	return nil
}

func createE2Eblob(ctx context.Context, env env.Interface, doc api.BillingDocument, resource azure.Resource) error {
	key, err := env.E2EStorageAccountKey(ctx)
	if err != nil {
		return err
	}
	client, err := azstorage.NewBasicClient(env.E2EStorageAccountName(), key)
	if err != nil {
		return err
	}
	blobclient := client.GetBlobService()

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

// isSubscriptionRegisteredToE2E returns true if the subscription is having the
// "Microsoft.ContainerService/SaveAROTestConfig" feature registered
func isSubscriptionRegisteredToE2E(sub *api.SubscriptionProperties) bool {
	if sub.TenantID == tenantIDMSFT || sub.TenantID == tenantIDAME {
		for _, feature := range sub.RegisteredFeatures {
			if feature.Name == api.SubscriptionFeatureE2E && feature.State == "Registered" {
				return true
			}
		}
	}
	return false
}
