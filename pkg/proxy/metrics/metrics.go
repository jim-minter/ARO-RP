package metrics

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"

	"github.com/Azure/ARO-RP/pkg/proxy/util/ingester"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/prometheus/prompb"
	"github.com/sirupsen/logrus"
)

func New(ctx context.Context, log *logrus.Entry, i ingester.Ingester) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		b, err = snappy.Decode(nil, b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var req prompb.WriteRequest
		err = proto.Unmarshal(b, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		uploader, err := i.NewUploader(ctx, "prometheus")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, timeseries := range req.Timeseries {
			labels := make(map[string]interface{}, len(timeseries.Labels))
			for _, label := range timeseries.Labels {
				labels[label.Name] = label.Value
			}

			for _, sample := range timeseries.Samples {
				if math.IsNaN(sample.Value) {
					continue
				}

				o := map[string]interface{}{
					"timestamp": sample.Timestamp,
					"labels":    labels,
					"value":     sample.Value,
				}

				err = json.NewEncoder(uploader).Encode(o)
				if err != nil {
					uploader.Close()
					uploader.Delete()
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}

		err = uploader.Close()
		if err != nil {
			uploader.Delete()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = uploader.Enqueue(ctx, []*ingester.Mapping{
			{
				Column:   "timestamp",
				Datatype: "datetime",
				Properties: map[string]interface{}{
					"path":      "$.timestamp",
					"transform": "DateTimeFromUnixMilliseconds",
				},
			},
			{
				Column:   "labels",
				Datatype: "dynamic",
				Properties: map[string]interface{}{
					"path": "$.labels",
				},
			},
			{
				Column:   "value",
				Datatype: "real",
				Properties: map[string]interface{}{
					"path": "$.value",
				},
			},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
