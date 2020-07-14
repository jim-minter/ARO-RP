package install

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/Azure/ARO-RP/pkg/dynamichelper"
	"github.com/Azure/ARO-RP/pkg/operator/deploy"
	"github.com/Azure/ARO-RP/pkg/util/restconfig"
)

func (i *Installer) readyToDeployAROOperator() (bool, error) {
	restConfig, err := restconfig.RestConfig(i.env, i.doc.OpenShiftCluster)
	if err != nil {
		return false, err
	}
	dh, err := dynamichelper.New(i.log, restConfig, dynamichelper.UpdatePolicy{})
	if err != nil {
		return false, nil
	}
	_, err = dh.Get(context.TODO(), "SecurityContextConstraints", "", "privileged")
	return err == nil, nil
}

func (i *Installer) ensureAROOperator(ctx context.Context) error {
	dep, err := deploy.New(i.log, i.env, i.doc.OpenShiftCluster, i.kubernetescli, i.securitycli, i.arocli)
	if err != nil {
		return err
	}
	return dep.CreateOrUpdate(ctx, i.env)
}

func (i *Installer) aroDeploymentReady() (bool, error) {
	dep, err := deploy.New(i.log, i.env, i.doc.OpenShiftCluster, i.kubernetescli, i.securitycli, i.arocli)
	if err != nil {
		return false, err
	}
	return dep.IsReady()
}
