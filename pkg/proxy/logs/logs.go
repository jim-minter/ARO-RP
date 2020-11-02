package logs

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"encoding/json"
	"net"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/proxy/util/fluent"
	"github.com/Azure/ARO-RP/pkg/proxy/util/ingester"
	"github.com/Azure/ARO-RP/pkg/util/errors"
	"github.com/Azure/ARO-RP/pkg/util/recover"
)

type logs struct {
	log      *logrus.Entry
	ingester ingester.Ingester
	c        net.Conn

	uploaders map[string]*uploader
}

type uploader struct {
	uploader ingester.Uploader
	fields   map[string]string
}

func Serve(ctx context.Context, log *logrus.Entry, ingester ingester.Ingester, c net.Conn) error {
	defer recover.Panic(log)
	defer c.Close()

	l := &logs{
		log:      log,
		ingester: ingester,
		c:        c,

		uploaders: map[string]*uploader{},
	}

	return l.serve(ctx)
}

func (l *logs) serve(ctx context.Context) error {
	msgs := fluent.Messages(l.log, l.c)

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case msg, ok := <-msgs:
			if !ok {
				return l.closeAndEnqueueAll(ctx)
			}

			u, err := l.getUploader(ctx, msg.Tag)
			if err != nil {
				err2 := l.closeAndEnqueueAll(ctx)
				if err2 != nil {
					l.log.Error(err2)
				}
				return err
			}

			for k, v := range msg.Record {
				if _, ok := u.fields[k]; !ok {
					switch v.(type) {
					case bool:
						u.fields[k] = "bool"
					case float64:
						u.fields[k] = "real"
					case string:
						u.fields[k] = "string"
					}
				}
			}

			msg.Record["timestamp"] = msg.Time.UnixNano()

			err = json.NewEncoder(u.uploader).Encode(msg.Record)
			if err != nil {
				err2 := l.closeAndEnqueueAll(ctx)
				if err2 != nil {
					l.log.Error(err2)
				}
				return err
			}

		case <-ticker.C:
			err := l.closeAndEnqueueAll(ctx)
			if err != nil {
				return err
			}
		}
	}
}

func (l *logs) getUploader(ctx context.Context, table string) (*uploader, error) {
	table = strings.ToLower(table)

	if l.uploaders[table] == nil {
		u, err := l.ingester.NewUploader(ctx, table)
		if err != nil {
			return nil, err
		}

		l.uploaders[table] = &uploader{
			uploader: u,
			fields:   map[string]string{},
		}
	}

	return l.uploaders[table], nil
}

func (l *logs) closeAndEnqueueAll(ctx context.Context) error {
	var errs errors.Errors

	for _, u := range l.uploaders {
		err := l.closeAndEnqueueOne(ctx, u)
		if err != nil {
			errs = append(errs, err)
		}
	}

	l.uploaders = map[string]*uploader{}

	return errs.AsError()
}

func (l *logs) closeAndEnqueueOne(ctx context.Context, u *uploader) error {
	err := u.uploader.Close()
	if err != nil {
		l.log.Error(err)
	}

	m := []*ingester.Mapping{
		{
			Column:   "timestamp",
			Datatype: "datetime",
			Properties: map[string]interface{}{
				"path":      "$.timestamp",
				"transform": "DateTimeFromUnixNanoseconds",
			},
		},
	}

	for f, datatype := range u.fields {
		m = append(m, &ingester.Mapping{
			Column:   f,
			Datatype: datatype,
			Properties: map[string]interface{}{
				"path": "$." + f,
			},
		})
	}

	sort.Slice(m[1:], func(i, j int) bool { return strings.Compare(m[i+1].Column, m[j+1].Column) < 0 })

	err = u.uploader.Enqueue(ctx, m)
	if err != nil {
		err2 := u.uploader.Delete()
		if err != nil {
			l.log.Error(err2)
		}
		return err
	}

	return nil
}
