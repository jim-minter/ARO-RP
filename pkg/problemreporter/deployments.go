package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"sort"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type deployments struct{}

func (deployments) Header() []string {
	return []string{
		"id",
		"namespace",
		"name",
		"replicas",
		"avaliableReplicas",
	}
}

func (deployments) Report(r *reporter) (recs [][]string, err error) {
	ds, err := r.cli.AppsV1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	sort.SliceStable(ds.Items, func(i, j int) bool { return ds.Items[i].Name < ds.Items[j].Name })
	sort.SliceStable(ds.Items, func(i, j int) bool { return ds.Items[i].Namespace < ds.Items[j].Namespace })
	for _, d := range ds.Items {
		if !isOpenShiftNamespace(d.Namespace) {
			continue
		}

		if d.Status.Replicas == d.Status.AvailableReplicas {
			continue
		}

		recs = append(recs, []string{r.oc.ID, d.Namespace, d.Name, strconv.Itoa(int(d.Status.Replicas)), strconv.Itoa(int(d.Status.AvailableReplicas))})
	}

	return recs, err
}
