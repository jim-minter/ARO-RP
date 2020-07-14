package deploy

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"os"
	"sort"

	securityclient "github.com/openshift/client-go/security/clientset/versioned"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/yaml"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/dynamichelper"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/genevalogging"
	aro "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
	aroclient "github.com/Azure/ARO-RP/pkg/util/aro-operator-client/clientset/versioned/typed/aro.openshift.io/v1alpha1"
	"github.com/Azure/ARO-RP/pkg/util/jsonpath"
	"github.com/Azure/ARO-RP/pkg/util/pullsecret"
	"github.com/Azure/ARO-RP/pkg/util/ready"
	"github.com/Azure/ARO-RP/pkg/util/restconfig"
	"github.com/Azure/ARO-RP/pkg/util/tls"
)

const (
	KubeNamespace     = "openshift-azure-operator"
	ACRPullSecretName = "acr-pullsecret-tokens"
)

type Operator interface {
	CreateOrUpdate(ctx context.Context, _env env.Interface) error
	IsReady() (bool, error)
}

type operator struct {
	log *logrus.Entry

	regTokens map[string]string

	clusterSpec *aro.ClusterSpec

	dh     dynamichelper.DynamicHelper
	cli    kubernetes.Interface
	seccli securityclient.Interface
	arocli aroclient.AroV1alpha1Interface
}

func New(log *logrus.Entry, _env env.Interface, oc *api.OpenShiftCluster, cli kubernetes.Interface, seccli securityclient.Interface, arocli aroclient.AroV1alpha1Interface) (Operator, error) {
	restConfig, err := restconfig.RestConfig(_env, oc)
	if err != nil {
		return nil, err
	}
	dh, err := dynamichelper.New(log, restConfig, dynamichelper.UpdatePolicy{
		IgnoreDefaults:                true,
		LogChanges:                    true,
		RetryOnConflict:               true,
		RefreshAPIResourcesOnNotFound: true,
	})
	if err != nil {
		return nil, err
	}

	o := &operator{
		log: log,

		regTokens: map[string]string{},

		clusterSpec: &aro.ClusterSpec{
			ResourceID: oc.ID,
			ACRName:    _env.ACRName(),
			Location:   _env.Location(),
			GenevaLogging: aro.GenevaLoggingSpec{
				ConfigVersion:            _env.ClustersGenevaLoggingConfigVersion(),
				MonitoringGCSEnvironment: _env.ClustersGenevaLoggingEnvironment(),
			},
			InternetChecker: aro.InternetCheckerSpec{
				URLs: []string{
					"https://registry.redhat.io/",
					"https://quay.io/",
					"https://sso.redhat.com/",
					"https://mirror.openshift.com/",
					"https://api.openshift.com/",
					"https://management.azure.com/",
				},
			},
		},

		dh:     dh,
		cli:    cli,
		seccli: seccli,
		arocli: arocli,
	}

	regName := _env.ACRName() + ".azurecr.io"
	for _, reg := range oc.Properties.RegistryProfiles {
		if reg.Name == regName && string(reg.Password) != "" {
			o.regTokens[regName] = reg.Username + ":" + string(reg.Password)
		}
	}
	if _, ok := _env.(env.Dev); ok {
		auths, err := pullsecret.Auths([]byte(os.Getenv("PULL_SECRET")))
		if err != nil {
			return nil, err
		}
		o.regTokens[regName] = auths[regName]["auth"].(string)
	}
	return o, nil
}

func (o *operator) resources(ctx context.Context, _env env.Interface) ([]runtime.Object, error) {
	// first static resources from Assets
	results := []runtime.Object{}
	for _, assetName := range AssetNames() {
		b, err := Asset(assetName)
		if err != nil {
			return nil, err
		}
		obj := &unstructured.Unstructured{}
		err = yaml.Unmarshal(b, obj)
		if err != nil {
			return nil, err
		}

		// set the image for the deployments
		if obj.GroupVersionKind().GroupKind().String() == "Deployment.apps" {
			for _, podSpec := range jsonpath.MustCompile("$.spec.template.spec").Get(obj.Object) {
				for _, contaner := range jsonpath.MustCompile("$.containers.*").Get(podSpec.(map[string]interface{})) {
					jsonpath.MustCompile("$.image").Set(contaner.(map[string]interface{}), _env.AROOperatorImage())
				}
			}
		}

		results = append(results, obj)
	}
	// then dynamic resources
	key, cert := _env.ClustersGenevaLoggingSecret()
	gcsKeyBytes, err := tls.PrivateKeyAsBytes(key)
	if err != nil {
		return nil, err
	}

	gcsCertBytes, err := tls.CertAsBytes(cert)
	if err != nil {
		return nil, err
	}

	// create a secret here for genevalogging, later we will copy it to
	// the genevalogging namespace.
	return append(results,
		&corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      genevalogging.CertificatesSecretName,
				Namespace: KubeNamespace,
			},
			StringData: map[string]string{
				"gcscert.pem": string(gcsCertBytes),
				"gcskey.pem":  string(gcsKeyBytes),
			},
		},
		&corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      ACRPullSecretName,
				Namespace: KubeNamespace,
			},
			Type:       corev1.SecretTypeOpaque,
			StringData: o.regTokens,
		},
		&aro.Cluster{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Cluster",
				APIVersion: "aro.openshift.io/v1alpha1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: aro.SingletonClusterName,
			},
			Spec: *o.clusterSpec,
		}), nil
}

func (o *operator) CreateOrUpdate(ctx context.Context, _env env.Interface) error {
	resources, err := o.resources(ctx, _env)
	if err != nil {
		return err
	}

	objects := []*unstructured.Unstructured{}
	for _, res := range resources {
		var un *unstructured.Unstructured
		un, err = dynamichelper.ToUnstructured(res)
		if err != nil {
			return err
		}
		objects = append(objects, un)
	}

	sort.Slice(objects, func(i, j int) bool {
		return dynamichelper.KindLess(objects[i].GetKind(), objects[j].GetKind())
	})
	for _, un := range objects {
		err = o.dh.CreateOrUpdate(ctx, un)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *operator) IsReady() (bool, error) {
	ok, err := ready.CheckDeploymentIsReady(o.cli.AppsV1().Deployments(KubeNamespace), "aro-operator-master")()
	if !ok || err != nil {
		return ok, err
	}
	return ready.CheckDeploymentIsReady(o.cli.AppsV1().Deployments(KubeNamespace), "aro-operator-worker")()
}
