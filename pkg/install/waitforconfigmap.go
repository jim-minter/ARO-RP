package install

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"time"

	configv1 "github.com/openshift/api/config/v1"
	configclient "github.com/openshift/client-go/config/clientset/versioned"
	operatorclient "github.com/openshift/client-go/operator/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"

	"github.com/Azure/ARO-RP/pkg/util/restconfig"
)

func (i *Installer) waitForClusterVersion(ctx context.Context) error {
	restConfig, err := restconfig.RestConfig(ctx, i.env, i.doc.OpenShiftCluster)
	if err != nil {
		return err
	}
	cli, err := configclient.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Minute)
	defer cancel()
	return wait.PollImmediateUntil(10*time.Second, func() (bool, error) {
		cv, err := cli.ConfigV1().ClusterVersions().Get("version", metav1.GetOptions{})
		if err == nil {
			for _, cond := range cv.Status.Conditions {
				if cond.Type == configv1.OperatorAvailable && cond.Status == configv1.ConditionTrue {
					return true, nil
				}
			}
		}
		return false, nil

	}, timeoutCtx.Done())
}

func (i *Installer) initializeKubernetesClients(ctx context.Context) error {
	// call this at some point at which the cluster is basically up?

	restConfig, err := restconfig.RestConfig(ctx, i.env, i.doc.OpenShiftCluster)
	if err != nil {
		return err
	}

	i.kubernetescli, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	i.operatorcli, err = operatorclient.NewForConfig(restConfig)
	return err
}

func (i *Installer) waitForBootstrapConfigmap(ctx context.Context) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Minute)
	defer cancel()

	return wait.PollImmediateUntil(10*time.Second, func() (bool, error) {
		cm, err := i.kubernetescli.CoreV1().ConfigMaps("kube-system").Get("bootstrap", metav1.GetOptions{})
		return err == nil && cm.Data["status"] == "complete", nil

	}, timeoutCtx.Done())
}
