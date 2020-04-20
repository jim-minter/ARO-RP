package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"sort"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var nodeConditionsExpected = map[v1.NodeConditionType]v1.ConditionStatus{
	v1.NodeDiskPressure:   v1.ConditionFalse,
	v1.NodeMemoryPressure: v1.ConditionFalse,
	v1.NodePIDPressure:    v1.ConditionFalse,
	v1.NodeReady:          v1.ConditionTrue,
}

type nodeConditions struct{}

func (nodeConditions) Header() []string {
	return []string{
		"id",
		"name",
		"type",
		"status",
		"message",
	}
}

func (nodeConditions) Report(r *reporter) (recs [][]string, err error) {
	ns, err := r.cli.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	sort.Slice(ns.Items, func(i, j int) bool { return ns.Items[i].Name < ns.Items[j].Name })
	for _, n := range ns.Items {
		sort.Slice(n.Status.Conditions, func(i, j int) bool { return n.Status.Conditions[i].Type < n.Status.Conditions[j].Type })
		for _, c := range n.Status.Conditions {
			if c.Status == nodeConditionsExpected[c.Type] {
				continue
			}

			recs = append(recs, []string{r.oc.ID, n.Name, string(c.Type), string(c.Status), c.Message})
		}
	}

	return recs, err
}
