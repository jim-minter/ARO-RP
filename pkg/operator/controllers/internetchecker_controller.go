package controllers

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/operator-framework/operator-sdk/pkg/status"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	aro "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	aroclient "github.com/Azure/ARO-RP/pkg/util/aro-operator-client/clientset/versioned/typed/aro.openshift.io/v1alpha1"
)

// InternetChecker reconciles a Cluster object
type InternetChecker struct {
	kubernetescli kubernetes.Interface
	arocli        aroclient.AroV1alpha1Interface
	log           *logrus.Entry
	scheme        *runtime.Scheme
	sr            *StatusReporter
	placement     string
}

func NewInternetChecker(log *logrus.Entry, kubernetescli kubernetes.Interface, arocli aroclient.AroV1alpha1Interface, scheme *runtime.Scheme, placement string) *InternetChecker {
	return &InternetChecker{
		kubernetescli: kubernetescli,
		arocli:        arocli,
		log:           log,
		scheme:        scheme,
		placement:     placement,
	}
}

// SimpleHTTPClient to aid in mocking
type SimpleHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// This is the permissions that this controller needs to work.
// "make generate" will run kubebuilder and cause operator/deploy/staticresources/*/role.yaml to be updated
// from the annotation below.
// +kubebuilder:rbac:groups=aro.openshift.io,resources=clusters,verbs=get;list;watch
// +kubebuilder:rbac:groups=aro.openshift.io,resources=clusters/status,verbs=get;update;patch

// Reconcile will keep checking that the cluster can connect to essential services.
func (r *InternetChecker) Reconcile(request ctrl.Request) (ctrl.Result, error) {
	if request.Name != aro.SingletonClusterName {
		return reconcile.Result{}, nil
	}
	ctx := context.TODO()
	instance, err := r.arocli.Clusters().Get(request.Name, v1.GetOptions{})
	if err != nil {
		return reconcile.Result{}, err
	}
	if len(instance.Spec.InternetChecker.URLs) == 0 {
		return ReconcileResultRequeueShort, nil
	}

	sitesNotAvailable := map[string]string{}
	for _, testurl := range instance.Spec.InternetChecker.URLs {
		checkErr := r.check(&http.Client{}, testurl)
		if checkErr != nil {
			sitesNotAvailable[testurl] = checkErr.Error()
		}
	}

	cTypeMap := map[string]status.ConditionType{
		"master": aro.InternetReachableFromMaster,
		"worker": aro.InternetReachableFromWorker,
	}
	if len(sitesNotAvailable) > 0 {
		msg := ""
		for k, v := range sitesNotAvailable {
			msg += "[" + k + "] " + v + "\n"
		}
		err = r.sr.SetConditionFalse(ctx, cTypeMap[r.placement], msg)
	} else {
		err = r.sr.SetConditionTrue(ctx, cTypeMap[r.placement], "Outgoing connection successful.")
	}
	if err != nil {
		r.log.Errorf("StatusReporter request:%v err:%v", request, err)
		return reconcile.Result{}, err
	}

	return ReconcileResultRequeueShort, nil
}

func (r *InternetChecker) check(client SimpleHTTPClient, testurl string) error {
	ctx := context.TODO()
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequest("GET", testurl, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req.WithContext(timeoutCtx))
	if err != nil {
		return err
	}
	if resp.StatusCode >= http.StatusInternalServerError {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		r.log.Warnf("check failed (%s) status:%s body:%s", testurl, resp.Status, string(b))
		return fmt.Errorf("check failed %s bad status:%s", testurl, resp.Status)
	}
	return nil
}

// SetupWithManager setup our mananger
func (r *InternetChecker) SetupWithManager(mgr ctrl.Manager) error {
	r.sr = NewStatusReporter(r.log, r.arocli, aro.SingletonClusterName)

	return ctrl.NewControllerManagedBy(mgr).
		For(&aro.Cluster{}).Named(InternetCheckerControllerName).
		Complete(r)
}
