package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"sort"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type podContainerConditions struct{}

func (podContainerConditions) Header() []string {
	return []string{
		"id",
		"namespace",
		"name",
		"reason",
		"message",
	}
}

func (podContainerConditions) Report(r *reporter) (recs [][]string, err error) {
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

		sort.Slice(p.Status.ContainerStatuses, func(i, j int) bool {
			return p.Status.ContainerStatuses[i].Name < p.Status.ContainerStatuses[j].Name
		})
		for _, cs := range p.Status.ContainerStatuses {
			if cs.State.Waiting == nil {
				continue
			}

			recs = append(recs, []string{r.oc.ID, p.Namespace, p.Name, cs.State.Waiting.Reason, cs.State.Waiting.Message})
		}
	}

	return recs, err
}
