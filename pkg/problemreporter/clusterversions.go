package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"github.com/Azure/ARO-RP/pkg/util/version"
	configv1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type clusterVersions struct{}

func (clusterVersions) Header() []string {
	return []string{
		"id",
		"desiredVersion",
		"actualVersion",
	}
}

func (clusterVersions) Report(r *reporter) ([][]string, error) {
	cv, err := r.configcli.ConfigV1().ClusterVersions().Get("version", metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	desiredVersion := cv.Status.Desired.Version
	if cv.Spec.DesiredUpdate != nil &&
		cv.Spec.DesiredUpdate.Version != "" {
		desiredVersion = cv.Spec.DesiredUpdate.Version
	}

	// Find the actual current cluster state. The history is ordered by most
	// recent first, so find the latest "Completed" status to get current
	// cluster version
	var actualVersion string
	for _, history := range cv.Status.History {
		if history.State == configv1.CompletedUpdate {
			actualVersion = history.Version
			break
		}
	}

	if actualVersion == version.OpenShiftVersion {
		return nil, nil
	}

	return [][]string{
		{
			r.oc.ID,
			desiredVersion,
			actualVersion,
		},
	}, nil
}
