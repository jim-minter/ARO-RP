package controllers

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"sort"

	securityclient "github.com/openshift/client-go/security/clientset/versioned"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/ARO-RP/pkg/dynamichelper"
	"github.com/Azure/ARO-RP/pkg/genevalogging"
	aro "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	aroclient "github.com/Azure/ARO-RP/pkg/util/aro-operator-client/clientset/versioned/typed/aro.openshift.io/v1alpha1"
)

// GenevaloggingReconciler reconciles a Cluster object
type GenevaloggingReconciler struct {
	kubernetescli kubernetes.Interface
	securitycli   securityclient.Interface
	arocli        aroclient.AroV1alpha1Interface
	restConfig    *rest.Config
	log           *logrus.Entry
	scheme        *runtime.Scheme
}

func NewGenevaloggingReconciler(log *logrus.Entry, kubernetescli kubernetes.Interface, securitycli securityclient.Interface, arocli aroclient.AroV1alpha1Interface, restConfig *rest.Config, scheme *runtime.Scheme) *GenevaloggingReconciler {
	return &GenevaloggingReconciler{
		securitycli:   securitycli,
		kubernetescli: kubernetescli,
		arocli:        arocli,
		restConfig:    restConfig,
		log:           log,
		scheme:        scheme,
	}
}

// Reconcile the genevalogging deployment.
func (r *GenevaloggingReconciler) Reconcile(request ctrl.Request) (ctrl.Result, error) {
	if request.Name != aro.SingletonClusterName {
		return reconcile.Result{}, nil
	}

	ctx := context.TODO()
	instance, err := r.arocli.Clusters().Get(request.Name, metav1.GetOptions{})
	if err != nil {
		// Error reading the object or not found - requeue the request.
		return reconcile.Result{}, err
	}

	newCert, err := r.certificatesSecret(instance)
	if err != nil {
		r.log.Error(err)
		return reconcile.Result{}, err
	}
	dh, err := dynamichelper.New(r.log, r.restConfig, dynamichelper.UpdatePolicy{
		IgnoreDefaults:  true,
		LogChanges:      true,
		RetryOnConflict: true,
	})
	if err != nil {
		r.log.Error(err)
		return reconcile.Result{}, err
	}
	gl := genevalogging.New(r.log, &instance.Spec, r.securitycli, newCert)

	resources, err := gl.Resources(ctx)
	if err != nil {
		r.log.Error(err)
		return reconcile.Result{}, err
	}

	var objects []*unstructured.Unstructured
	for _, res := range resources {
		var un *unstructured.Unstructured
		un, err = dynamichelper.ToUnstructured(res)
		if err != nil {
			r.log.Error(err)
			return reconcile.Result{}, err
		}

		// This sets the reference on all objects that we create
		// to our cluster instance. This causes the Owns() below to work and
		// to get Reconcile events when anything happens to our objects.
		err = controllerutil.SetControllerReference(instance, un, r.scheme)
		if err != nil {
			r.log.Errorf("SetControllerReference %s/%s: %v", instance.Kind, instance.Name, err)
			return reconcile.Result{}, err
		}
		objects = append(objects, un)
	}
	dynamichelper.HashWorkloadConfigs(objects)

	sort.Slice(objects, func(i, j int) bool {
		return dynamichelper.KindLess(objects[i].GetKind(), objects[j].GetKind())
	})

	for _, un := range objects {
		err = dh.CreateOrUpdate(ctx, un)
		if err != nil {
			r.log.Error(err)
			return reconcile.Result{}, err
		}
	}

	// watching should catch all changes, but double check later..
	return ReconcileResultRequeueLong, nil
}

func (r *GenevaloggingReconciler) certificatesSecret(instance *aro.Cluster) (*corev1.Secret, error) {
	newCert, err := r.kubernetescli.CoreV1().Secrets(genevalogging.KubeNamespace).Get(genevalogging.CertificatesSecretName, metav1.GetOptions{})
	if err != nil && apierrors.IsNotFound(err) {
		// copy the certificates from our namespace into the genevalogging one.
		certs, err := r.kubernetescli.CoreV1().Secrets(OperatorNamespace).Get(genevalogging.CertificatesSecretName, metav1.GetOptions{})
		if err != nil {
			r.log.Errorf("Error reading the certificates secret: %v", err)
			return nil, err
		}

		newCert = &corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      genevalogging.CertificatesSecretName,
				Namespace: genevalogging.KubeNamespace,
			},
			Data: certs.Data,
		}
	} else if err != nil {
		return nil, err
	}
	return newCert, nil
}

func genevaloggingRelatedObjects() []corev1.ObjectReference {
	return []corev1.ObjectReference{
		{Kind: "Namespace", Name: genevalogging.KubeNamespace},
		{Kind: "SecurityContextConstraints", Name: "privileged-genevalogging"},
		{Kind: "Secret", Name: genevalogging.CertificatesSecretName, Namespace: genevalogging.KubeNamespace},
		{Kind: "ConfigMap", Name: "fluent-config", Namespace: genevalogging.KubeNamespace},
		{Kind: "ServiceAccount", Name: "geneva", Namespace: genevalogging.KubeNamespace},
		{Kind: "DaemonSet", Name: "mdsd", Namespace: genevalogging.KubeNamespace},
	}
}

// SetupWithManager setup our mananger
func (r *GenevaloggingReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&aro.Cluster{}).Owns(&appsv1.DaemonSet{}).Owns(&corev1.Secret{}).Owns(&corev1.ConfigMap{}).Named(GenevaLoggingControllerName).
		Complete(r)
}
