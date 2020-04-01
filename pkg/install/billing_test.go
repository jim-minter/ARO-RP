package install

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	mock_database "github.com/Azure/ARO-RP/pkg/util/mocks/database"
)

func TestCreateBillingEntry(t *testing.T) {
	ctx := context.Background()
	mockSubID := "11111111-1111-1111-1111-111111111111"
	mockTenantID := mockSubID
	location := "eastus"

	type test struct {
		name         string
		openshiftdoc *api.OpenShiftClusterDocument
		mocks        func(*test, *mock_database.MockBilling, *mock_database.MockSubscriptions)
		wantError    error
	}

	for _, tt := range []*test{
		{
			name: "create a new billing entry",
			openshiftdoc: &api.OpenShiftClusterDocument{
				Key:                       fmt.Sprintf("/subscriptions/%s/resourcegroups/rgName/providers/microsoft.redhatopenshift/openshiftclusters/clusterName", mockSubID),
				ClusterResourceGroupIDKey: fmt.Sprintf("/subscriptions/%s/resourcegroups/rgName", mockSubID),
				ID:                        mockSubID,
				OpenShiftCluster: &api.OpenShiftCluster{
					Properties: api.OpenShiftClusterProperties{
						ServicePrincipalProfile: api.ServicePrincipalProfile{
							TenantID: mockTenantID,
						},
					},
					Location: location,
				},
			},
			mocks: func(tt *test, billing *mock_database.MockBilling, subscription *mock_database.MockSubscriptions) {
				billingDoc := &api.BillingDocument{
					Key:                       tt.openshiftdoc.Key,
					ClusterResourceGroupIDKey: tt.openshiftdoc.ClusterResourceGroupIDKey,
					ID:                        mockSubID,
					Billing: &api.Billing{
						TenantID: tt.openshiftdoc.OpenShiftCluster.Properties.ServicePrincipalProfile.TenantID,
						Location: tt.openshiftdoc.OpenShiftCluster.Location,
					},
				}

				billing.EXPECT().
					Create(gomock.Any(), billingDoc).
					Return(billingDoc, nil)

				subscription.EXPECT().
					Get(gomock.Any(), mockSubID).
					Return(&api.SubscriptionDocument{
						Subscription: &api.Subscription{
							Properties: &api.SubscriptionProperties{
								RegisteredFeatures: []api.RegisteredFeatureProfile{},
							},
						},
					}, nil)
			},
		},
		{
			name: "error on create a new billing entry",
			openshiftdoc: &api.OpenShiftClusterDocument{
				Key:                       fmt.Sprintf("/subscriptions/%s/resourcegroups/rgName/providers/microsoft.redhatopenshift/openshiftclusters/clusterName", mockSubID),
				ClusterResourceGroupIDKey: fmt.Sprintf("/subscriptions/%s/resourcegroups/rgName", mockSubID),
				ID:                        mockSubID,
				OpenShiftCluster: &api.OpenShiftCluster{
					Properties: api.OpenShiftClusterProperties{
						ServicePrincipalProfile: api.ServicePrincipalProfile{
							TenantID: mockTenantID,
						},
					},
					Location: location,
				},
			},
			mocks: func(tt *test, billing *mock_database.MockBilling, _ *mock_database.MockSubscriptions) {
				billingDoc := &api.BillingDocument{
					Key:                       tt.openshiftdoc.Key,
					ClusterResourceGroupIDKey: tt.openshiftdoc.ClusterResourceGroupIDKey,
					ID:                        mockSubID,
					Billing: &api.Billing{
						TenantID: tt.openshiftdoc.OpenShiftCluster.Properties.ServicePrincipalProfile.TenantID,
						Location: tt.openshiftdoc.OpenShiftCluster.Location,
					},
				}

				billing.EXPECT().
					Create(gomock.Any(), billingDoc).
					Return(nil, tt.wantError)
			},
			wantError: fmt.Errorf("Error creating document"),
		},
		{
			name: "billing document already existing on DB on create",
			openshiftdoc: &api.OpenShiftClusterDocument{
				Key:                       fmt.Sprintf("/subscriptions/%s/resourcegroups/rgName/providers/microsoft.redhatopenshift/openshiftclusters/clusterName", mockSubID),
				ClusterResourceGroupIDKey: fmt.Sprintf("/subscriptions/%s/resourcegroups/rgName", mockSubID),
				ID:                        mockSubID,
				OpenShiftCluster: &api.OpenShiftCluster{
					Properties: api.OpenShiftClusterProperties{
						ServicePrincipalProfile: api.ServicePrincipalProfile{
							TenantID: mockTenantID,
						},
					},
					Location: location,
				},
			},
			mocks: func(tt *test, billing *mock_database.MockBilling, _ *mock_database.MockSubscriptions) {
				billingDoc := &api.BillingDocument{
					Key:                       fmt.Sprintf("/subscriptions/%s/resourcegroups/rgName/providers/microsoft.redhatopenshift/openshiftclusters/clusterName", mockSubID),
					ClusterResourceGroupIDKey: tt.openshiftdoc.ClusterResourceGroupIDKey,
					ID:                        mockSubID,
					Billing: &api.Billing{
						TenantID: tt.openshiftdoc.OpenShiftCluster.Properties.ServicePrincipalProfile.TenantID,
						Location: tt.openshiftdoc.OpenShiftCluster.Location,
					},
				}

				billing.EXPECT().
					Create(gomock.Any(), billingDoc).
					Return(nil, &cosmosdb.Error{
						StatusCode: http.StatusConflict,
					})
			},
			wantError: nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			billing := mock_database.NewMockBilling(controller)
			subscription := mock_database.NewMockSubscriptions(controller)
			log := logrus.NewEntry(logrus.StandardLogger())
			tt.mocks(tt, billing, subscription)
			i := &Installer{
				log:     log,
				doc:     tt.openshiftdoc,
				billing: billing,
				sub:     subscription,
			}

			err := i.createBillingRecord(ctx)
			if err != nil {
				if tt.wantError != err {
					t.Errorf("Error want (%s), having (%s)", tt.wantError.Error(), err.Error())
				}
			}
		})
	}
}

func TestIsSubscriptionRegisteredToE2E(t *testing.T) {
	mockSubID := "11111111-1111-1111-1111-111111111111"
	for _, tt := range []struct {
		name string
		sub  *api.SubscriptionProperties
		want bool
	}{
		{
			name: "empty sub",
			sub:  &api.SubscriptionProperties{},
			want: false,
		},
		{
			name: "sub wihout feature the flag registered and not internal tenant",
			sub: &api.SubscriptionProperties{
				TenantID: mockSubID,
				RegisteredFeatures: []api.RegisteredFeatureProfile{
					{
						Name:  "RandomFeature",
						State: "Registered",
					},
				},
			},
			want: false,
		},
		{
			name: "sub with feature the flag registered and not internal tenant",
			sub: &api.SubscriptionProperties{
				TenantID: mockSubID,
				RegisteredFeatures: []api.RegisteredFeatureProfile{
					{
						Name:  api.SubscriptionFeatureE2E,
						State: "Registered",
					},
				},
			},
			want: false,
		},
		{
			name: "AME internal tenant and feature flag not registered",
			sub: &api.SubscriptionProperties{
				TenantID: tenantIDAME,
				RegisteredFeatures: []api.RegisteredFeatureProfile{
					{
						Name:  "RandomFeature",
						State: "Registered",
					},
				},
			},
			want: false,
		},
		{
			name: "MSFT internal tenant and feature flag not registered",
			sub: &api.SubscriptionProperties{
				TenantID: tenantIDMSFT,
				RegisteredFeatures: []api.RegisteredFeatureProfile{
					{
						Name:  "RandomFeature",
						State: "Registered",
					},
				},
			},
			want: false,
		},
		{
			name: "AME internal tenant and feature flag registered",
			sub: &api.SubscriptionProperties{
				TenantID: tenantIDAME,
				RegisteredFeatures: []api.RegisteredFeatureProfile{
					{
						Name:  api.SubscriptionFeatureE2E,
						State: "Registered",
					},
				},
			},
			want: true,
		},
		{
			name: "MSFT internal tenant and feature flag registered",
			sub: &api.SubscriptionProperties{
				TenantID: tenantIDMSFT,
				RegisteredFeatures: []api.RegisteredFeatureProfile{
					{
						Name:  api.SubscriptionFeatureE2E,
						State: "Registered",
					},
				},
			},
			want: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result := isSubscriptionRegisteredToE2E(tt.sub)
			if result != tt.want {
				t.Errorf("result is :%t, want %t", result, tt.want)
			}

		})
	}
}
