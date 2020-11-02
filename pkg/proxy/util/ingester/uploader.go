package ingester

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"time"

	azstorage "github.com/Azure/azure-sdk-for-go/storage"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/util/deployment"
	"github.com/Azure/ARO-RP/pkg/util/recover"
)

type Uploader interface {
	io.WriteCloser
	Enqueue(context.Context, []*Mapping) error
	Delete() error
}

type uploader struct {
	log            *logrus.Entry
	deploymentMode deployment.Mode
	ingester       *ingester
	table          string

	uuid     string
	sasToken string
	blob     *azstorage.Blob

	r  *io.PipeReader
	w  *io.PipeWriter
	gz *gzip.Writer

	size  int
	errch chan error
}

func (i *ingester) NewUploader(ctx context.Context, table string) (Uploader, error) {
	uuid := uuid.NewV4().String()

	endpoint, sasToken, containerName, err := i.getIngestionResource(ctx, "TempStorage")
	if err != nil {
		return nil, err
	}

	cli, err := azstorage.NewAccountSASClientFromEndpointToken(endpoint, sasToken)
	if err != nil {
		return nil, err
	}

	blobService := cli.GetBlobService()
	container := blobService.GetContainerReference(containerName)

	r, w := io.Pipe()

	u := &uploader{
		log:            i.log,
		deploymentMode: i.deploymentMode,
		ingester:       i,
		table:          table,

		uuid:     uuid,
		sasToken: sasToken,
		blob:     container.GetBlobReference(time.Now().UTC().Format("20060102150405-") + uuid + ".json.gz"),

		r:  r,
		w:  w,
		gz: gzip.NewWriter(w),

		errch: make(chan error, 1),
	}

	go u.upload()

	return u, nil
}

func (u *uploader) upload() {
	defer recover.Panic(u.log)
	u.errch <- u.blob.CreateBlockBlobFromReader(u.r, nil)
}

func (u *uploader) Write(b []byte) (int, error) {
	n, err := u.gz.Write(b)
	u.size += n
	return n, err
}

func (u *uploader) Close() error {
	err := u.gz.Close()
	if err != nil {
		return err
	}

	err = u.w.Close()
	if err != nil {
		return err
	}

	return <-u.errch
}

func (u *uploader) Delete() error {
	return u.blob.Delete(nil)
}

type enqueueMessage struct {
	ID                   string                 `json:"Id,omitempty"`
	BlobPath             string                 `json:"BlobPath,omitempty"`
	RawDataSize          int                    `json:"RawDataSize,omitempty"`
	DatabaseName         string                 `json:"DatabaseName,omitempty"`
	TableName            string                 `json:"TableName,omitempty"`
	RetainBlobOnSuccess  bool                   `json:"RetainBlobOnSuccess,omitempty"`
	FlushImmediately     bool                   `json:"FlushImmediately,omitempty"`
	ReportLevel          int                    `json:"ReportLevel,omitempty"`
	ReportMethod         int                    `json:"ReportMethod,omitempty"`
	AdditionalProperties map[string]interface{} `json:"AdditionalProperties,omitempty"`
}

func (u *uploader) Enqueue(ctx context.Context, mapping []*Mapping) error {
	endpoint, sasToken, queueName, err := u.ingester.getIngestionResource(ctx, "SecuredReadyForAggregationQueue")
	if err != nil {
		return err
	}

	cli, err := azstorage.NewAccountSASClientFromEndpointToken(endpoint, sasToken)
	if err != nil {
		return err
	}

	queueService := cli.GetQueueService()
	queue := queueService.GetQueueReference(queueName)

	b, err := json.Marshal(mapping)
	if err != nil {
		return err
	}

	m := &enqueueMessage{
		ID:               u.uuid,
		BlobPath:         u.blob.GetURL() + "?" + u.sasToken,
		RawDataSize:      u.size,
		DatabaseName:     u.ingester.database,
		TableName:        u.table,
		FlushImmediately: u.deploymentMode == deployment.Development,
		AdditionalProperties: map[string]interface{}{
			"ingestionMapping":     string(b),
			"ingestionMappingType": "json",
			"extend_schema":        true,
		},
	}

	u.ingester.mu.RLock()
	m.AdditionalProperties["authorizationContext"] = u.ingester.kustoIdentityToken
	u.ingester.mu.RUnlock()

	b, err = json.Marshal(m)
	if err != nil {
		return err
	}

	u.log.Printf("enqueuing %s", u.blob.GetURL()+"?"+u.sasToken)

	message := queue.GetMessageReference(base64.StdEncoding.EncodeToString(b))

	return message.Put(nil)
}
