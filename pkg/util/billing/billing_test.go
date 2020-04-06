package billing

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"testing"

	"github.com/Azure/ARO-RP/pkg/api"
)

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
						Name:  api.FeatureSaveAROTestConfig,
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
						Name:  api.FeatureSaveAROTestConfig,
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
						Name:  api.FeatureSaveAROTestConfig,
						State: "Registered",
					},
				},
			},
			want: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSubscriptionRegisteredToE2E(tt.sub)
			if result != tt.want {
				t.Errorf("result is :%t, want %t", result, tt.want)
			}

		})
	}
}
