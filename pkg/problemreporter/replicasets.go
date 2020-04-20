package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"sort"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type replicasets struct{}

func (replicasets) Header() []string {
	return []string{
		"id",
		"namespace",
		"name",
		"replicas",
		"availableReplicas",
	}
}

func (replicasets) Report(r *reporter) (recs [][]string, err error) {
	rss, err := r.cli.AppsV1().ReplicaSets("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	sort.SliceStable(rss.Items, func(i, j int) bool { return rss.Items[i].Name < rss.Items[j].Name })
	sort.SliceStable(rss.Items, func(i, j int) bool { return rss.Items[i].Namespace < rss.Items[j].Namespace })
	for _, rs := range rss.Items {
		if !isOpenShiftNamespace(rs.Namespace) {
			continue
		}

		if rs.Status.Replicas == rs.Status.AvailableReplicas {
			continue
		}

		recs = append(recs, []string{r.oc.ID, rs.Namespace, rs.Name, strconv.Itoa(int(rs.Status.Replicas)), strconv.Itoa(int(rs.Status.AvailableReplicas))})
	}

	return recs, err
}
