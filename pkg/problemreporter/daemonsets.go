package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"sort"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type daemonsets struct{}

func (daemonsets) Header() []string {
	return []string{
		"id",
		"namespace",
		"name",
		"desiredNumberScheduled",
		"numberAvailable",
	}
}

func (daemonsets) Report(r *reporter) (recs [][]string, err error) {
	dss, err := r.cli.AppsV1().DaemonSets("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	sort.SliceStable(dss.Items, func(i, j int) bool { return dss.Items[i].Name < dss.Items[j].Name })
	sort.SliceStable(dss.Items, func(i, j int) bool { return dss.Items[i].Namespace < dss.Items[j].Namespace })
	for _, ds := range dss.Items {
		if !isOpenShiftNamespace(ds.Namespace) {
			continue
		}

		if ds.Status.DesiredNumberScheduled == ds.Status.NumberAvailable {
			continue
		}

		recs = append(recs, []string{r.oc.ID, ds.Namespace, ds.Name, strconv.Itoa(int(ds.Status.DesiredNumberScheduled)), strconv.Itoa(int(ds.Status.NumberAvailable))})
	}

	return recs, err
}
