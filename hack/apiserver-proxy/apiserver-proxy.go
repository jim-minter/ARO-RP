package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"flag"

	"github.com/Azure/ARO-RP/pkg/util/apiserverproxy"
	utillog "github.com/Azure/ARO-RP/pkg/util/log"
	"github.com/Azure/ARO-RP/pkg/util/version"
)

var (
	certFile       = flag.String("certFile", "secrets/apiserver-proxy.crt", "file containing server certificate")
	keyFile        = flag.String("keyFile", "secrets/apiserver-proxy.key", "file containing server key")
	clientCertFile = flag.String("clientCertFile", "secrets/apiserver-proxy-client.crt", "file containing client certificate")
	subnet         = flag.String("subnet", "10.0.0.0/8", "allowed subnet")
)

func main() {
	log := utillog.GetLogger()

	log.Printf("starting, git commit %s", version.GitCommit)

	flag.Parse()

	s := &apiserverproxy.Server{
		CertFile:       *certFile,
		KeyFile:        *keyFile,
		ClientCertFile: *clientCertFile,
		Subnet:         *subnet,
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
