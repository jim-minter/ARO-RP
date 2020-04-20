package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"sort"

	configv1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type clusterOperatorConditionsIgnoreStruct struct {
	Name   string
	Type   configv1.ClusterStatusConditionType
	Status configv1.ConditionStatus
}

var clusterOperatorConditionsExpected = map[configv1.ClusterStatusConditionType]configv1.ConditionStatus{
	configv1.OperatorAvailable:   configv1.ConditionTrue,
	configv1.OperatorDegraded:    configv1.ConditionFalse,
	configv1.OperatorProgressing: configv1.ConditionFalse,
	configv1.OperatorUpgradeable: configv1.ConditionTrue,
}

var clusterOperatorConditionsIgnore = map[clusterOperatorConditionsIgnoreStruct]struct{}{
	{"insights", "Disabled", configv1.ConditionFalse}:                                         {},
	{"insights", "Disabled", configv1.ConditionTrue}:                                          {},
	{"openshift-controller-manager", configv1.OperatorUpgradeable, configv1.ConditionUnknown}: {},
	{"service-ca", configv1.OperatorUpgradeable, configv1.ConditionUnknown}:                   {},
	{"service-catalog-apiserver", configv1.OperatorUpgradeable, configv1.ConditionUnknown}:    {},
}

type clusterOperatorConditions struct{}

func (clusterOperatorConditions) Header() []string {
	return []string{
		"id",
		"name",
		"type",
		"status",
		"message",
	}
}

func (clusterOperatorConditions) Report(r *reporter) (recs [][]string, err error) {
	cos, err := r.configcli.ConfigV1().ClusterOperators().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	sort.Slice(cos.Items, func(i, j int) bool { return cos.Items[i].Name < cos.Items[j].Name })
	for _, co := range cos.Items {
		sort.Slice(co.Status.Conditions, func(i, j int) bool { return co.Status.Conditions[i].Type < co.Status.Conditions[j].Type })
		for _, c := range co.Status.Conditions {
			if clusterOperatorConditionIsExpected(&co, &c) {
				continue
			}

			recs = append(recs, []string{r.oc.ID, co.Name, string(c.Type), string(c.Status), c.Message})
		}
	}

	return recs, err
}

func clusterOperatorConditionIsExpected(co *configv1.ClusterOperator, c *configv1.ClusterOperatorStatusCondition) bool {
	if _, ok := clusterOperatorConditionsIgnore[clusterOperatorConditionsIgnoreStruct{
		Name:   co.Name,
		Type:   c.Type,
		Status: c.Status,
	}]; ok {
		return true
	}

	return c.Status == clusterOperatorConditionsExpected[c.Type]
}
