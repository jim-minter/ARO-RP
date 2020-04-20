package problemreporter

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/csv"
	"io"
	"reflect"
	"sort"
	"strings"
	"time"

	configclient "github.com/openshift/client-go/config/clientset/versioned"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/util/restconfig"
)

type Reporter interface {
	Header() []string
	Report(r *reporter) (recs [][]string, err error)
}

type reporter struct {
	env env.Interface
	oc  *api.OpenShiftCluster

	cli       kubernetes.Interface
	configcli configclient.Interface
}

func newReporter(env env.Interface, oc *api.OpenShiftCluster) (*reporter, error) {
	restConfig, err := restconfig.RestConfig(env, oc)
	if err != nil {
		return nil, err
	}

	cli, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	configcli, err := configclient.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	return &reporter{
		env: env,
		oc:  oc,

		cli:       cli,
		configcli: configcli,
	}, nil
}

func Report(ctx context.Context, log *logrus.Entry, env env.Interface, w io.Writer, db database.OpenShiftClusters) error {
	reporters := []struct {
		Reporter
		buf *bytes.Buffer
		w   *csv.Writer
	}{
		{Reporter: &clusterOperatorConditions{}},
		{Reporter: &clusterOperatorVersions{}},
		{Reporter: &clusterVersions{}},
		{Reporter: &daemonsets{}},
		{Reporter: &deployments{}},
		{Reporter: &nodeConditions{}},
		{Reporter: &podConditions{}},
		{Reporter: &podContainerConditions{}},
		{Reporter: &replicasets{}},
		{Reporter: &statefulsets{}},
	}

	for i := range reporters {
		reporters[i].buf = &bytes.Buffer{}
		reporters[i].w = csv.NewWriter(reporters[i].buf)

		err := reporters[i].w.Write(reporters[i].Header())
		if err != nil {
			return err
		}
	}

	docs, err := db.ListAll(ctx)
	if err != nil {
		return err
	}

	sort.Slice(docs.OpenShiftClusterDocuments, func(i, j int) bool {
		return docs.OpenShiftClusterDocuments[i].OpenShiftCluster.ID < docs.OpenShiftClusterDocuments[j].OpenShiftCluster.ID
	})

next:
	for _, doc := range docs.OpenShiftClusterDocuments {
		ps := doc.OpenShiftCluster.Properties.ProvisioningState
		fps := doc.OpenShiftCluster.Properties.FailedProvisioningState

		switch {
		case ps == api.ProvisioningStateCreating,
			ps == api.ProvisioningStateDeleting,
			ps == api.ProvisioningStateFailed &&
				(fps == api.ProvisioningStateCreating ||
					fps == api.ProvisioningStateDeleting):
			continue next
		}

		log.Info(doc.OpenShiftCluster.ID)

		reporter, err := newReporter(env, doc.OpenShiftCluster)
		if err != nil {
			return err
		}

		for _, r := range reporters {
			recs, err := r.Report(reporter)
			if err != nil {
				log.Error(err)
				continue
			}

			err = r.w.WriteAll(recs)
			if err != nil {
				return err
			}
		}
	}

	gw := gzip.NewWriter(w)
	tw := tar.NewWriter(gw)

	now := time.Now()
	for _, r := range reporters {
		r.w.Flush()
		err = r.w.Error()
		if err != nil {
			return err
		}

		err = tw.WriteHeader(&tar.Header{
			Name:    reflect.TypeOf(r.Reporter).Elem().Name() + ".csv",
			Size:    int64(r.buf.Len()),
			Mode:    0644,
			ModTime: now,
		})
		if err != nil {
			return err
		}

		_, err = tw.Write(r.buf.Bytes())
		if err != nil {
			return err
		}
	}

	err = tw.Close()
	if err != nil {
		return err
	}

	return gw.Close()
}

func isOpenShiftNamespace(namespace string) bool {
	if namespace == "openshift-operators" {
		return false
	}

	return namespace == "" ||
		namespace == "default" ||
		namespace == "openshift" ||
		strings.HasPrefix(string(namespace), "kube-") ||
		strings.HasPrefix(string(namespace), "openshift-")
}
