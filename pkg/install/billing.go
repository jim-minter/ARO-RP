package install

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"net/http"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
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

	if e2eErr := i.e2e.CreateOrUpdateE2EBlob(ctx, i.env, i.sub, billingDoc); e2eErr != nil {
		i.log.Warnf("CreateOrUpdateE2EBlob failed : %s", e2eErr.Error())
	}

	return nil
}
