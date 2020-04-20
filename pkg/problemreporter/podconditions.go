package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"sort"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var podConditionsExpected = map[v1.PodConditionType]v1.ConditionStatus{
	v1.ContainersReady: v1.ConditionTrue,
	v1.PodInitialized:  v1.ConditionTrue,
	v1.PodScheduled:    v1.ConditionTrue,
	v1.PodReady:        v1.ConditionTrue,
}

type podConditions struct{}

func (podConditions) Header() []string {
	return []string{
		"id",
		"namespace",
		"name",
		"type",
		"status",
		"message",
	}
}

func (podConditions) Report(r *reporter) (recs [][]string, err error) {
	ps, err := r.cli.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	sort.SliceStable(ps.Items, func(i, j int) bool { return ps.Items[i].Name < ps.Items[j].Name })
	sort.SliceStable(ps.Items, func(i, j int) bool { return ps.Items[i].Namespace < ps.Items[j].Namespace })
	for _, p := range ps.Items {
		if !isOpenShiftNamespace(p.Namespace) {
			continue
		}

		if p.Status.Phase == v1.PodSucceeded {
			continue
		}

		sort.Slice(p.Status.Conditions, func(i, j int) bool { return p.Status.Conditions[i].Type < p.Status.Conditions[j].Type })
		for _, c := range p.Status.Conditions {
			if c.Status == podConditionsExpected[c.Type] {
				continue
			}

			recs = append(recs, []string{r.oc.ID, p.Namespace, p.Name, string(c.Type), string(c.Status), c.Message})
		}
	}

	return recs, err
}
