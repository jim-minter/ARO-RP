package database

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"time"

	"github.com/Azure/ARO-RP/pkg/util/recover"
)

func (db *Database) emitMetrics(ctx context.Context) {
	defer recover.Panic(db.log)
	t := time.NewTicker(time.Minute)
	defer t.Stop()

	for {
		i, err := db.OpenShiftClusters.QueueLength(ctx)
		if err != nil {
			db.log.Error(err)
		} else {
			db.m.EmitGauge("database.queue.openshiftclusters.length", int64(i), nil)
		}

		select {
		case <-t.C:
			return
		}
	}

}
