package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"sort"

	"github.com/Azure/ARO-RP/pkg/util/version"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type clusterOperatorVersions struct{}

func (clusterOperatorVersions) Header() []string {
	return []string{
		"id",
		"name",
		"version",
	}
}

func (clusterOperatorVersions) Report(r *reporter) (recs [][]string, err error) {
	cos, err := r.configcli.ConfigV1().ClusterOperators().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	sort.Slice(cos.Items, func(i, j int) bool { return cos.Items[i].Name < cos.Items[j].Name })
	for _, co := range cos.Items {
		for _, v := range co.Status.Versions {
			if v.Name != "operator" {
				continue
			}

			if v.Version == version.OpenShiftVersion {
				continue
			}

			recs = append(recs, []string{r.oc.ID, co.Name, v.Version})
		}
	}

	return recs, err
}
