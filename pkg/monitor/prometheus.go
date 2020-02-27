package monitor

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/util/portforward"
)

func (mon *monitor) checkPrometheus(ctx context.Context, oc *api.OpenShiftCluster) error {
	hc := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
				_, port, err := net.SplitHostPort(address)
				if err != nil {
					return nil, err
				}

				return portforward.DialContext(ctx, mon.env, oc, "openshift-monitoring", "alertmanager-main-0", port)
			},
		},
	}

	// TODO: try other pods if -0 isn't available?
	req, err := http.NewRequest(http.MethodGet, "http://alertmanager-main.openshift-monitoring.svc:9093/api/v2/alerts", nil)
	if err != nil {
		return err
	}

	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	// TODO: this probably exists in a library which we can import
	var alerts []struct {
		Labels   map[string]string `json:"labels,omitempty"`
		StartsAt time.Time         `json:"startsAt,omitempty"`
		EndsAt   time.Time         `json:"endsAt,omitempty"`
	}
	err = json.NewDecoder(resp.Body).Decode(&alerts)
	if err != nil {
		return err
	}

	for _, alert := range alerts {
		// TODO: check StartsAt / EndsAt?
		mon.clusterm.EmitGauge("monitoring.prometheus.alert", 1, map[string]string{
			"resource": oc.ID,
			"alert":    alert.Labels["alertname"],
		})
	}

	return nil
}
