package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/env"
	pkgproxy "github.com/Azure/ARO-RP/pkg/proxy"
	"github.com/Azure/ARO-RP/pkg/proxy/util/ingester"
)

func proxy(ctx context.Context, log *logrus.Entry) error {
	_env, err := env.NewCore(ctx, log)
	if err != nil {
		return err
	}

	key, certs, err := _env.GetCertificateSecret(ctx, env.ProxyServerSecretName)
	if err != nil {
		return err
	}

	ingester, err := ingester.NewIngester(ctx, log, _env.DeploymentMode(), "https://jminter.eastus.kusto.windows.net", "test")
	if err != nil {
		return err
	}

	p := pkgproxy.New(log, key, certs, ingester)
	return p.Run(ctx)
}
