package main

import (
	securityv1 "github.com/openshift/api/security/v1"
	azureproviderv1beta1 "github.com/openshift/cluster-api-provider-azure/pkg/apis/azureprovider/v1beta1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"

	arov1alpha1 "github.com/Azure/ARO-RP/pkg/operator/apis/aro.openshift.io/v1alpha1"
)

func init() {
	runtime.Must(apiextensions.AddToScheme(scheme.Scheme))
	runtime.Must(securityv1.AddToScheme(scheme.Scheme))
	runtime.Must(arov1alpha1.AddToScheme(scheme.Scheme))
	runtime.Must(azureproviderv1beta1.SchemeBuilder.AddToScheme(scheme.Scheme))
}
