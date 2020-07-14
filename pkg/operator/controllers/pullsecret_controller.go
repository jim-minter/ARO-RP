package controllers

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/ARO-RP/pkg/operator/deploy"
	aroclient "github.com/Azure/ARO-RP/pkg/util/aro-operator-client/clientset/versioned/typed/aro.openshift.io/v1alpha1"
	"github.com/Azure/ARO-RP/pkg/util/pullsecret"
)

var pullSecretName = types.NamespacedName{Name: "pull-secret", Namespace: "openshift-config"}

// PullsecretReconciler reconciles a Cluster object
type PullsecretReconciler struct {
	kubernetescli kubernetes.Interface
	arocli        aroclient.AroV1alpha1Interface
	log           *logrus.Entry
	scheme        *runtime.Scheme
}

func NewPullsecretReconciler(log *logrus.Entry, kubernetescli kubernetes.Interface, arocli aroclient.AroV1alpha1Interface, scheme *runtime.Scheme) *PullsecretReconciler {
	return &PullsecretReconciler{
		log:           log,
		kubernetescli: kubernetescli,
		arocli:        arocli,
		scheme:        scheme,
	}
}

// Reconcile will make sure that the ACR part of the pull secret is correct
func (r *PullsecretReconciler) Reconcile(request ctrl.Request) (ctrl.Result, error) {
	if request.NamespacedName != pullSecretName {
		// filter out other secrets.
		return reconcile.Result{}, nil
	}
	repoTokens, err := r.requiredRepoTokens()
	if err != nil || len(repoTokens) == 0 {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, retry.RetryOnConflict(retry.DefaultRetry, func() error {
		ps, isCreate, err := r.pullsecret(request)
		if err != nil {
			return err
		}

		// validate
		if !json.Valid(ps.Data[v1.DockerConfigJsonKey]) {
			r.log.Info("Pull Secret is not valid json - recreating")
			delete(ps.Data, v1.DockerConfigJsonKey)
		}
		if ps.Data == nil {
			ps.Data = map[string][]byte{}
		}

		// repair data
		newPS, changed, err := pullsecret.Replace(ps.Data[corev1.DockerConfigJsonKey], repoTokens)
		if err != nil {
			return err
		} else if len(changed) > 0 {
			r.log.Info("Repaired the following repositories ", changed)
		}

		// repair Secret type
		if ps.Type != v1.SecretTypeDockerConfigJson {
			ps = &v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pull-secret",
					Namespace: "openshift-config",
				},
				Type: v1.SecretTypeDockerConfigJson,
				Data: map[string][]byte{},
			}
			isCreate = true
			r.log.Info("Pull Secret has the wrong secret type - recreating")

			// unfortunately the type field is immutable.
			err = r.kubernetescli.CoreV1().Secrets(ps.Namespace).Delete(ps.Name, nil)
			if err != nil {
				return err
			}

			// there is a small risk of crashing here: if that happens, we will
			// restart, create a new pull secret, and will have dropped the rest
			// of the customer's pull secret on the floor :-(
		}
		if !isCreate && len(changed) == 0 {
			return nil
		}

		ps.Data[corev1.DockerConfigJsonKey] = newPS

		if isCreate {
			r.log.Info("Re-creating the Pull Secret")
			_, err = r.kubernetescli.CoreV1().Secrets("openshift-config").Create(ps)
		} else {
			r.log.Info("Updating the Pull Secret")
			_, err = r.kubernetescli.CoreV1().Secrets("openshift-config").Update(ps)
		}
		return err
	})
}

func (r *PullsecretReconciler) pullsecret(request ctrl.Request) (*v1.Secret, bool, error) {
	var isCreate bool
	ps, err := r.kubernetescli.CoreV1().Secrets(request.Namespace).Get(request.Name, metav1.GetOptions{})
	switch {
	case apierrors.IsNotFound(err):
		ps = &v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      request.Name,
				Namespace: request.Namespace,
			},
			Type: v1.SecretTypeDockerConfigJson,
		}
		isCreate = true
	case err != nil:
		return nil, false, err
	}
	return ps, isCreate, nil
}

func (r *PullsecretReconciler) requiredRepoTokens() (map[string]string, error) {
	repoTokensSecret, err := r.kubernetescli.CoreV1().Secrets(OperatorNamespace).Get(deploy.ACRPullSecretName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("Error reading the repoToken secret: %v", err)
	}

	result := map[string]string{}
	for k, v := range repoTokensSecret.Data {
		result[k] = string(v)
	}

	return result, nil
}

func pullsecretRelatedObjects() []corev1.ObjectReference {
	return []corev1.ObjectReference{
		{Kind: "Secret", Name: pullSecretName.Name, Namespace: pullSecretName.Namespace},
	}
}

func triggerReconcile(secret *corev1.Secret) bool {
	return secret.Name == pullSecretName.Name && secret.Namespace == pullSecretName.Namespace
}

// SetupWithManager setup our mananger
func (r *PullsecretReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// The pull secret may already be deleted when controller starts
	initialRequest := ctrl.Request{
		NamespacedName: types.NamespacedName{
			Namespace: pullSecretName.Namespace,
			Name:      pullSecretName.Name,
		},
	}
	_, isCreate, err := r.pullsecret(initialRequest)
	if err == nil && isCreate {
		r.Reconcile(initialRequest)
	}

	isPullSecret := predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			oldSecret, ok := e.ObjectOld.(*corev1.Secret)
			if !ok {
				return false
			}
			newSecret, ok := e.ObjectNew.(*corev1.Secret)
			if !ok {
				return false
			}
			return (triggerReconcile(oldSecret) || triggerReconcile(newSecret))
		},
		CreateFunc: func(e event.CreateEvent) bool {
			secret, ok := e.Object.(*corev1.Secret)
			if !ok {
				return false
			}
			return triggerReconcile(secret)
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			secret, ok := e.Object.(*corev1.Secret)
			if !ok {
				return false
			}
			return triggerReconcile(secret)
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.Secret{}).WithEventFilter(isPullSecret).Named(PullSecretControllerName).
		Complete(r)
}
