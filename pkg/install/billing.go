package install

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	"github.com/Azure/ARO-RP/pkg/util/billing"
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

	if billing.IsSubscriptionRegisteredToE2E(subDocument.Subscription.Properties) {
		// We are not failing the operation if we cannot write to e2e storage account, just warning
		// The blob can only be created when the billing record is also written for the first time
		errBlob := billing.CreateOrUpdateE2Eblob(ctx, i.env, *billingDoc, resource)
		if errBlob != nil {
			i.log.Warnf("CreateOrUpdateE2Eblob failed : %s", errBlob.Error())
		}
	}
	return nil
}
