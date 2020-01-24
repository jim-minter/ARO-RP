package cosmodb

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/metrics"
)

var _ http.RoundTripper = (*tracerRoundTripper)(nil)

type tracerRoundTripper struct {
	log *logrus.Entry
	m   metrics.Interface
	tr  http.RoundTripper
}

func New(log *logrus.Entry, tr *http.Transport, m metrics.Interface) *tracerRoundTripper {
	return &tracerRoundTripper{
		log: log,
		m:   m,
		tr:  tr,
	}
}

func (t *tracerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := t.tr.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	var ru float64
	// Sometimes we get request-charge="" because pkranges API is free
	// We log this on debug mode only and ignore
	requestCharge := strings.Trim(resp.Header.Get("x-ms-request-charge"), `"`)
	if requestCharge != "" {
		ru, err = strconv.ParseFloat(requestCharge, 64)
		if err != nil {
			// we don't want to kill all DB calls if this fails
			t.log.Error(err)
		}
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) >= 2 && parts[len(parts)-2] == "docs" {
		parts[len(parts)-1] = "{id}"
	}
	path := strings.Join(parts, "/")

	t.m.EmitGauge("client.cosmosdb.count", 1, map[string]string{
		"code": strconv.Itoa(resp.StatusCode),
		"verb": r.Method,
		"path": path,
	})

	t.m.EmitFloat("client.cosmosdb.duration", time.Now().Sub(start).Seconds(), map[string]string{
		"code": strconv.Itoa(resp.StatusCode),
		"verb": r.Method,
		"path": path,
	})

	t.m.EmitFloat("client.cosmosdb.requestunits", ru, map[string]string{
		"code": strconv.Itoa(resp.StatusCode),
		"verb": r.Method,
		"path": path,
	})

	return resp, err

}
