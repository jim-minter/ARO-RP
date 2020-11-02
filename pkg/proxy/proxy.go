package proxy

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/proxy/acr"
	"github.com/Azure/ARO-RP/pkg/proxy/logs"
	"github.com/Azure/ARO-RP/pkg/proxy/metrics"
	"github.com/Azure/ARO-RP/pkg/proxy/util/ingester"
	"github.com/Azure/ARO-RP/pkg/util/recover"
)

type proxy struct {
	log *logrus.Entry

	key   *rsa.PrivateKey
	certs []*x509.Certificate

	ingester ingester.Ingester
}

func New(log *logrus.Entry, key *rsa.PrivateKey, certs []*x509.Certificate, ingester ingester.Ingester) *proxy {
	return &proxy{
		log: log,

		key:   key,
		certs: certs,

		ingester: ingester,
	}
}

func (p *proxy) Run(ctx context.Context) error {
	config := &tls.Config{
		Certificates: []tls.Certificate{
			{
				PrivateKey: p.key,
			},
		},
	}

	for _, cert := range p.certs {
		config.Certificates[0].Certificate = append(config.Certificates[0].Certificate, cert.Raw)
	}

	l, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		return err
	}

	p.log.Print("listening")
	//return p.serveACR(ctx, l)
	//return p.serveLogs(ctx, l)
	return p.serveMetrics(ctx, l)
}

func (p *proxy) serveACR(ctx context.Context, l net.Listener) error {
	// podman pull --tls-verify=false localhost:8443/aro:latest

	acr, _ /* TODO */ := acr.New(p.log, "arosvcint.azurecr.io")
	return http.Serve(l, acr)
}

func (p *proxy) serveLogs(ctx context.Context, l net.Listener) error {
	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}

		go func() {
			defer recover.Panic(p.log)

			err = logs.Serve(ctx, p.log, p.ingester, c)
			if err != nil {
				p.log.Error(err)
			}
		}()
	}
}

func (p *proxy) serveMetrics(ctx context.Context, l net.Listener) error {
	return http.Serve(l, metrics.New(ctx, p.log, p.ingester))
}
