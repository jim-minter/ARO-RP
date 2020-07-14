package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"flag"
	"fmt"
	"strings"

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

var (
	scheme = runtime.NewScheme()
)

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
	flag.Parse()
	role := strings.ToLower(flag.Arg(1))
	ctrl.SetLogger(utillog.GetRLogger(log))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: "0", // disabled
		Port:               8443,
		LeaderElection:     false, // disabled
	})
	if err != nil {
		return err
	}
	restConfig, err := ctrl.GetConfig()
	if err != nil {
		return err
	}

	kubernetescli, err := kubernetes.NewForConfig(mgr.GetConfig())
	if err != nil {
		return err
	}
	securitycli, err := securityclient.NewForConfig(mgr.GetConfig())
	if err != nil {
		return err
	}
	arocli, err := aroclient.NewForConfig(mgr.GetConfig())
	if err != nil {
		return err
	}

	if role == "master" {
		if err = (controllers.NewGenevaloggingReconciler(
			log.WithField("controller", controllers.GenevaLoggingControllerName),
			kubernetescli, securitycli, arocli,
			restConfig,
			mgr.GetScheme())).SetupWithManager(mgr); err != nil {
			return fmt.Errorf("unable to create controller Genevalogging: %v", err)
		}
		if err = (controllers.NewPullsecretReconciler(
			log.WithField("controller", controllers.PullSecretControllerName),
			kubernetescli, arocli,
			mgr.GetScheme())).SetupWithManager(mgr); err != nil {
			return fmt.Errorf("unable to create controller PullSecret: %v", err)
		}
		if err = (controllers.NewAlertWebhookReconciler(
			log.WithField("controller", controllers.AlertwebhookControllerName),
			kubernetescli,
			mgr.GetScheme())).SetupWithManager(mgr); err != nil {
			return fmt.Errorf("unable to create controller AlertWebhook: %v", err)
		}
	}

	if err = (controllers.NewInternetChecker(
		log.WithField("controller", controllers.InternetCheckerControllerName),
		kubernetescli, arocli,
		mgr.GetScheme(),
		role,
	)).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("unable to create controller InternetChecker: %v", err)
	}
	// +kubebuilder:scaffold:builder

	log.Info("starting manager")
	return mgr.Start(ctrl.SetupSignalHandler())
}
