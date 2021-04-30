package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/database"
	pkgdbtoken "github.com/Azure/ARO-RP/pkg/dbtoken"
	"github.com/Azure/ARO-RP/pkg/env"
	pkggateway "github.com/Azure/ARO-RP/pkg/gateway"
	"github.com/Azure/ARO-RP/pkg/metrics/noop"
	utilnet "github.com/Azure/ARO-RP/pkg/util/net"
)

func gateway(ctx context.Context, log *logrus.Entry) error {
	_env, err := env.NewCore(ctx, log)
	if err != nil {
		return err
	}

	dbc, err := database.NewDatabaseClient(log.WithField("component", "database"), _env, nil, &noop.Noop{}, nil)
	if err != nil {
		return err
	}

	msiRefresherAuthorizer, err := _env.NewMSIAuthorizer(env.MSIContextGateway, pkgdbtoken.Resource)
	if err != nil {
		return err
	}

	// TODO: refactor this poor man's feature flag
	insecureSkipVerify := _env.IsLocalDevelopmentMode()
	for _, feature := range strings.Split(os.Getenv("GATEWAY_FEATURES"), ",") {
		if feature == "InsecureSkipVerifyDBTokenCertificate" {
			insecureSkipVerify = true
			break
		}
	}

	dbRefresher, err := pkgdbtoken.NewRefresher(log, _env, msiRefresherAuthorizer, insecureSkipVerify, dbc, "gateway")
	if err != nil {
		return err
	}

	dbGateway, err := database.NewGateway(ctx, _env.IsLocalDevelopmentMode(), dbc)
	if err != nil {
		return err
	}

	go func() {
		_ = dbRefresher.Run(ctx)
	}()

	log.Print("waiting for database token")
	for !dbRefresher.Ready() {
		time.Sleep(time.Second)
	}

	httpl, err := utilnet.Listen("tcp", ":8080", pkggateway.SocketSize)
	if err != nil {
		return err
	}

	httpsl, err := utilnet.Listen("tcp", ":8443", pkggateway.SocketSize)
	if err != nil {
		return err
	}

	log.Print("listening")

	p, err := pkggateway.NewGateway(ctx, _env, log.WithField("component", "gateway"), log.WithField("component", "gateway-access"), dbGateway, httpsl, httpl, os.Getenv("ACR_RESOURCE_ID"), os.Getenv("GATEWAY_DOMAINS"))
	if err != nil {
		return err
	}

	sigterm := make(chan os.Signal, 1)
	cancelCtx, cancel := context.WithCancel(ctx)
	done := make(chan struct{})
	signal.Notify(sigterm, syscall.SIGTERM)

	go p.Run(cancelCtx, done)

	<-sigterm
	log.Print("received SIGTERM")
	cancel()
	<-done

	return nil
}