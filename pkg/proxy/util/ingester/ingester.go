package ingester

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"sync"
	"time"

	"github.com/Azure/azure-kusto-go/kusto"
	"github.com/Azure/azure-kusto-go/kusto/data/table"
	azstorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/util/deployment"
	"github.com/Azure/ARO-RP/pkg/util/recover"
)

type Ingester interface {
	NewUploader(context.Context, string) (Uploader, error)
}

type ingester struct {
	log            *logrus.Entry
	deploymentMode deployment.Mode
	mu             sync.RWMutex

	cli      *kusto.Client
	database string

	ingestionResources map[string][]*url.URL
	kustoIdentityToken string
}

func NewIngester(ctx context.Context, log *logrus.Entry, deploymentMode deployment.Mode, url, database string) (Ingester, error) {
	authorizer, err := auth.NewAuthorizerFromEnvironmentWithResource(url)
	if err != nil {
		return nil, err
	}

	cli, err := kusto.New(url, kusto.Authorization{
		Authorizer: authorizer,
	})
	if err != nil {
		return nil, err
	}

	i := &ingester{
		log:      log,
		cli:      cli,
		database: database,
	}

	err = i.refreshOne(ctx)
	if err != nil {
		return nil, err
	}

	err = i.logIngestFailures(ctx)
	if err != nil {
		return nil, err
	}

	go i.refresh(ctx)

	return i, nil
}

func (i *ingester) refresh(ctx context.Context) {
	defer recover.Panic(i.log)

	t := time.NewTicker(time.Minute)
	defer t.Stop()

	for {
		<-t.C

		err := i.refreshOne(ctx)
		if err != nil {
			i.log.Error(err)
		}
	}
}

func (i *ingester) refreshOne(ctx context.Context) error {
	token, err := i.getKustoIdentityToken(ctx)
	if err != nil {
		return err
	}

	resources, err := i.getIngestionResources(ctx)
	if err != nil {
		return err
	}

	i.mu.Lock()
	i.ingestionResources = resources
	i.kustoIdentityToken = token
	i.mu.Unlock()

	return nil
}

func (i *ingester) getIngestionResources(ctx context.Context) (map[string][]*url.URL, error) {
	rows, err := i.cli.Mgmt(ctx, "NetDefaultDB", kusto.NewStmt(".get ingestion resources"), kusto.IngestionEndpoint())
	if err != nil {
		return nil, err
	}

	resources := map[string][]*url.URL{}

	err = rows.Do(
		func(r *table.Row) error {
			url, err := url.Parse(r.Values[1].String())
			if err != nil {
				return err
			}

			resources[r.Values[0].String()] = append(resources[r.Values[0].String()], url)

			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (i *ingester) getIngestionResource(ctx context.Context, resourceName string) (string, string, string, error) {
	i.mu.RLock()
	resources := i.ingestionResources[resourceName]
	i.mu.RUnlock()

	if len(resources) == 0 {
		return "", "", "", fmt.Errorf("resource %s not found", resourceName)
	}

	resource := resources[rand.Intn(len(resources))]

	return resource.String(), resource.RawQuery, resource.Path[1:], nil
}

func (i *ingester) getKustoIdentityToken(ctx context.Context) (string, error) {
	rows, err := i.cli.Mgmt(ctx, "NetDefaultDB", kusto.NewStmt(".get kusto identity token"), kusto.IngestionEndpoint())
	if err != nil {
		return "", err
	}

	var token string
	err = rows.Do(
		func(r *table.Row) error {
			token = r.Values[0].String()
			return nil
		},
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (i *ingester) logIngestFailures(ctx context.Context) error {
	endpoint, sasToken, queueName, err := i.getIngestionResource(ctx, "FailedIngestionsQueue")
	if err != nil {
		return err
	}

	cli, err := azstorage.NewAccountSASClientFromEndpointToken(endpoint, sasToken)
	if err != nil {
		return err
	}

	queueService := cli.GetQueueService()
	queue := queueService.GetQueueReference(queueName)

	go func() {
		defer recover.Panic(i.log)

		t := time.NewTicker(time.Second)
		defer t.Stop()

		for {
			msgs, err := queue.GetMessages(&azstorage.GetMessagesOptions{
				NumOfMessages: 32,
			})
			if err != nil {
				i.log.Error(err)
			}

			for _, msg := range msgs {
				b, err := base64.StdEncoding.DecodeString(msg.Text)
				if err != nil {
					i.log.Error(err)
					continue
				}

				i.log.Warn(string(b))

				err = msg.Delete(nil)
				if err != nil {
					i.log.Error(err)
				}
			}

			if len(msgs) == 0 {
				<-t.C
			}
		}
	}()

	return nil
}
