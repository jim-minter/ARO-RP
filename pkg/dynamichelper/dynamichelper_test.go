package dynamichelper

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/fake"
	ktesting "k8s.io/client-go/testing"

	"github.com/Azure/ARO-RP/pkg/api"
	utillog "github.com/Azure/ARO-RP/pkg/util/log"
)

func TestFindGVR(t *testing.T) {
	tests := []struct {
		name      string
		resources []*metav1.APIResourceList
		kind      string
		want      *schema.GroupVersionResource
		wantErr   error
	}{
		{
			name: "find one",
			resources: []*metav1.APIResourceList{
				{
					GroupVersion: "v1",
					APIResources: []metav1.APIResource{
						{
							Name: "configmaps",
							Kind: "ConfigMap",
						},
					},
				},
			},
			kind: "configmap",
			want: &schema.GroupVersionResource{Group: "", Version: "v1", Resource: "configmaps"},
		},
		{
			name: "find best version",
			resources: []*metav1.APIResourceList{
				{
					GroupVersion: "v1",
					APIResources: []metav1.APIResource{
						{
							Name: "configmaps",
							Kind: "ConfigMap",
						},
					},
				},
				{
					GroupVersion: "v1beta1",
					APIResources: []metav1.APIResource{
						{
							Name: "configmaps",
							Kind: "ConfigMap",
						},
					},
				},
			},
			kind: "configmap",
			want: &schema.GroupVersionResource{Group: "", Version: "v1", Resource: "configmaps"},
		},
		{
			name: "find full group.resource",
			resources: []*metav1.APIResourceList{
				{
					GroupVersion: "metal3.io/v1alpha1",
					APIResources: []metav1.APIResource{
						{
							Name: "baremetalhosts",
							Kind: "BareMetalHost",
						},
					},
				},
			},
			kind: "baremetalhost.metal3.io",
			want: &schema.GroupVersionResource{Group: "metal3.io", Version: "v1alpha1", Resource: "baremetalhosts"},
		},
		{
			name: "no sub.resources",
			resources: []*metav1.APIResourceList{
				{
					GroupVersion: "metal3.io/v1alpha1",
					APIResources: []metav1.APIResource{
						{
							Name: "baremetalhosts/status",
							Kind: "BareMetalHost",
						},
					},
				},
			},
			kind: "baremetalhost/status",
			wantErr: api.NewCloudError(
				http.StatusBadRequest, api.CloudErrorCodeNotFound,
				"", "The groupKind '%s' was not found.", "baremetalhost/status"),
		},
		{
			name: "empty resources",
			kind: "configmap",
			wantErr: api.NewCloudError(
				http.StatusBadRequest, api.CloudErrorCodeNotFound,
				"", "The groupKind '%s' was not found.", "configmap"),
		},
		{
			name: "find all kinds",
			resources: []*metav1.APIResourceList{
				{
					GroupVersion: "metal3.io/v1alpha1",
					APIResources: []metav1.APIResource{
						{
							Name: "baremetalhosts",
							Kind: "BareMetalHost",
						},
					},
				},
				{
					GroupVersion: "plastic.io/v1alpha1",
					APIResources: []metav1.APIResource{
						{
							Name: "plastichosts",
							Kind: "BareMetalHost",
						},
					},
				},
			},
			kind: "baremetalhost",
			want: nil,
			wantErr: api.NewCloudError(
				http.StatusBadRequest, api.CloudErrorCodeInvalidParameter,
				"", "The groupKind '%s' matched multiple groupKinds (baremetalhost.metal3.io, baremetalhost.plastic.io).", "baremetalhost"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ka := &dynamicHelper{apiresources: tt.resources}

			got, err := ka.findGVR(tt.kind, "")
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Error(got)
			}
		})
	}
}

func TestCreateOrUpdate(t *testing.T) {
	testlog := utillog.GetLogger()
	ctx := context.TODO()
	v1configmap := []*metav1.APIResourceList{
		{
			GroupVersion: "v1",
			APIResources: []metav1.APIResource{
				{
					Name: "configmaps",
					Kind: "ConfigMap",
				},
			},
		},
	}
	tests := []struct {
		name         string
		existing     []runtime.Object
		new          *unstructured.Unstructured
		apiresources []*metav1.APIResourceList
		wantCreate   bool
		wantUpdate   bool
		wantErr      bool
	}{
		{
			name:     "create",
			existing: []runtime.Object{},
			new: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind": "ConfigMap",
					"metadata": map[string]interface{}{
						"namespace": "openshift-azure-logging",
						"name":      "config",
					},
				},
			},
			wantCreate:   true,
			apiresources: v1configmap,
		},
		{
			name: "update",
			existing: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"kind":       "ConfigMap",
						"apiVersion": "v1",
						"metadata": map[string]interface{}{
							"namespace": "openshift-azure-logging",
							"name":      "config",
						},
						"data": map[string]interface{}{
							"audit.conf": "1",
						},
					},
				},
			},
			new: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "ConfigMap",
					"apiVersion": "v1",
					"metadata": map[string]interface{}{
						"namespace": "openshift-azure-logging",
						"name":      "config",
					},
					"data": map[string]interface{}{
						"audit.conf": "2",
					},
				},
			},
			wantUpdate:   true,
			apiresources: v1configmap,
		},
		{
			name: "no update needed",
			existing: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"kind":       "ConfigMap",
						"apiVersion": "v1",
						"metadata": map[string]interface{}{
							"namespace": "openshift-azure-logging",
							"name":      "config",
						},
						"data": map[string]interface{}{
							"audit.conf": "2",
						},
					},
				},
			},
			new: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "ConfigMap",
					"apiVersion": "v1",
					"metadata": map[string]interface{}{
						"namespace": "openshift-azure-logging",
						"name":      "config",
					},
					"data": map[string]interface{}{
						"audit.conf": "2",
					},
				},
			},
			apiresources: v1configmap,
		},
		{
			name: "no update needed with IgnoreDefaults",
			existing: []runtime.Object{
				&unstructured.Unstructured{
					Object: map[string]interface{}{
						"kind":       "ConfigMap",
						"apiVersion": "v1",
						"metadata": map[string]interface{}{
							"namespace":  "openshift-azure-logging",
							"name":       "config",
							"generation": "4", // <- should be stripped out by clean()
						},
						"data": map[string]interface{}{
							"audit.conf": "2",
						},
					},
				},
			},
			new: &unstructured.Unstructured{
				Object: map[string]interface{}{
					"kind":       "ConfigMap",
					"apiVersion": "v1",
					"metadata": map[string]interface{}{
						"namespace": "openshift-azure-logging",
						"name":      "config",
					},
					"data": map[string]interface{}{
						"audit.conf": "2",
					},
				},
			},
			apiresources: v1configmap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var created, updated bool
			fakeDyn := fake.NewSimpleDynamicClient(runtime.NewScheme(), tt.existing...)

			fakeDyn.PrependReactor("create", "configmaps", func(action ktesting.Action) (handled bool, ret runtime.Object, err error) {
				created = true
				return false, nil, nil
			})

			fakeDyn.PrependReactor("update", "configmaps", func(action ktesting.Action) (handled bool, ret runtime.Object, err error) {
				updated = true
				return false, nil, nil
			})

			dh := &dynamicHelper{
				log: testlog,
				updatePolicy: UpdatePolicy{
					LogChanges:     true,
					IgnoreDefaults: true,
				},
				dyn:          fakeDyn,
				apiresources: tt.apiresources,
			}
			if err := dh.CreateOrUpdate(ctx, tt.new); (err != nil) != tt.wantErr {
				t.Errorf("dynamicHelper.CreateOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantCreate != created {
				t.Errorf("dynamicHelper.CreateOrUpdate create should be %v but was %v", tt.wantCreate, created)
			}
			if tt.wantUpdate != updated {
				t.Errorf("dynamicHelper.CreateOrUpdate update should be %v but was %v", tt.wantUpdate, updated)
			}
		})
	}
}
