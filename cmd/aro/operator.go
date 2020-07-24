package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"flag"
	"fmt"

	securityclient "github.com/openshift/client-go/security/clientset/versioned"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"

	aro "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	"github.com/Azure/ARO-RP/pkg/operator/controllers"
	aroclient "github.com/Azure/ARO-RP/pkg/util/aro-operator-client/clientset/versioned/typed/aro.openshift.io/v1alpha1"
	utillog "github.com/Azure/ARO-RP/pkg/util/log"
	// +kubebuilder:scaffold:imports
)

const (
	roleMaster = "master"
	roleWorker = "worker"
)

var scheme = runtime.NewScheme()

func init() {
	err := clientgoscheme.AddToScheme(scheme)
	if err != nil {
		panic(err)
	}

	err = aro.AddToScheme(scheme)
	if err != nil {
		panic(err)
	}
	// +kubebuilder:scaffold:scheme
}

func operator(ctx context.Context, log *logrus.Entry) error {
	role := flag.Arg(1)
	switch role {
	case roleMaster, roleWorker:
	default:
		return fmt.Errorf("invalid role %s", role)
	}

	ctrl.SetLogger(utillog.GetRLogger(log))

	restConfig, err := ctrl.GetConfig()
	if err != nil {
		return err
	}

	mgr, err := ctrl.NewManager(restConfig, ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: "0", // disabled
		Port:               8443,
	})
	if err != nil {
		return err
	}

	kubernetescli, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}
	securitycli, err := securityclient.NewForConfig(restConfig)
	if err != nil {
		return err
	}
	arocli, err := aroclient.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	if role == roleMaster {
		if err = (controllers.NewGenevaloggingReconciler(
			log.WithField("controller", controllers.GenevaLoggingControllerName),
			kubernetescli, securitycli, arocli,
			restConfig,
			scheme)).SetupWithManager(mgr); err != nil {
			return fmt.Errorf("unable to create controller Genevalogging: %v", err)
		}
		if err = (controllers.NewPullsecretReconciler(
			log.WithField("controller", controllers.PullSecretControllerName),
			kubernetescli, arocli,
			scheme)).SetupWithManager(mgr); err != nil {
			return fmt.Errorf("unable to create controller PullSecret: %v", err)
		}
		if err = (controllers.NewAlertWebhookReconciler(
			log.WithField("controller", controllers.AlertwebhookControllerName),
			kubernetescli,
			scheme)).SetupWithManager(mgr); err != nil {
			return fmt.Errorf("unable to create controller AlertWebhook: %v", err)
		}
	}

	if err = (controllers.NewInternetChecker(
		log.WithField("controller", controllers.InternetCheckerControllerName),
		kubernetescli, arocli,
		scheme,
		role,
	)).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("unable to create controller InternetChecker: %v", err)
	}
	// +kubebuilder:scaffold:builder

	log.Info("starting manager")
	return mgr.Start(ctrl.SetupSignalHandler())
}
