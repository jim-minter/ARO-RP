package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"sort"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type statefulsets struct{}

func (statefulsets) Header() []string {
	return []string{
		"id",
		"namespace",
		"name",
		"replicas",
		"readyReplicas",
	}
}

func (statefulsets) Report(r *reporter) (recs [][]string, err error) {
	sss, err := r.cli.AppsV1().StatefulSets("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	sort.SliceStable(sss.Items, func(i, j int) bool { return sss.Items[i].Name < sss.Items[j].Name })
	sort.SliceStable(sss.Items, func(i, j int) bool { return sss.Items[i].Namespace < sss.Items[j].Namespace })
	for _, ss := range sss.Items {
		if !isOpenShiftNamespace(ss.Namespace) {
			continue
		}

		if ss.Status.Replicas == ss.Status.ReadyReplicas {
			continue
		}

		recs = append(recs, []string{r.oc.ID, ss.Namespace, ss.Name, strconv.Itoa(int(ss.Status.Replicas)), strconv.Itoa(int(ss.Status.ReadyReplicas))})
	}

	return recs, err
}
