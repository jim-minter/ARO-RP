package install

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	mock_billing "github.com/Azure/ARO-RP/pkg/util/mocks/billing"
)

func TestCreateBillingEntry(t *testing.T) {
	ctx := context.Background()

	type test struct {
		openshiftdoc *api.OpenShiftClusterDocument
		name         string
		mocks        func(*test, *mock_billing.MockManager)
		wantError    error
	}

	for _, tt := range []*test{
		{
			openshiftdoc: &api.OpenShiftClusterDocument{},
			name:         "manager create is called and doesn't return an error when create doesn't return an error",
			mocks: func(tt *test, billing *mock_billing.MockManager) {
				billing.EXPECT().
					Create(gomock.Any(), tt.openshiftdoc).
					Return(nil)
			},
			wantError: nil,
		},
		{
			openshiftdoc: &api.OpenShiftClusterDocument{},
			name:         "manager create is called and returns an error on create returning an error",
			mocks: func(tt *test, billing *mock_billing.MockManager) {
				billing.EXPECT().
					Create(gomock.Any(), tt.openshiftdoc).
					Return(tt.wantError)
			},
			wantError: fmt.Errorf("Error"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()
			billing := mock_billing.NewMockManager(controller)
			log := logrus.NewEntry(logrus.StandardLogger())
			tt.mocks(tt, billing)
			i := &Installer{
				log:     log,
				doc:     tt.openshiftdoc,
				billing: billing,
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
