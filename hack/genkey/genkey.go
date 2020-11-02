package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Azure/ARO-RP/pkg/util/tls"
)

var (
	client   = flag.Bool("client", false, "generate client certificate")
	ca       = flag.Bool("ca", false, "generate ca certificate")
	keyFile  = flag.String("keyFile", "", `file containing signing key in der format (default "" - self-signed)`)
	certFile = flag.String("certFile", "", `file containing signing certificate in der format (default "" - self-signed)`)
)

func run() error {
	var signingKey *rsa.PrivateKey
	var signingCert *x509.Certificate

	if *keyFile != "" {
		b, err := ioutil.ReadFile(*keyFile)
		if err != nil {
			return err
		}

		signingKey, err = x509.ParsePKCS1PrivateKey(b)
		if err != nil {
			return err
		}
	}

	if *certFile != "" {
		b, err := ioutil.ReadFile(*certFile)
		if err != nil {
			return err
		}

		signingCert, err = x509.ParseCertificate(b)
		if err != nil {
			return err
		}
	}

	key, cert, err := tls.GenerateKeyAndCertificate(flag.Args(), signingKey, signingCert, *ca, *client)
	if err != nil {
		return err
	}

	// key in der format
	err = ioutil.WriteFile(flag.Arg(0)+".key", x509.MarshalPKCS1PrivateKey(key), 0600)
	if err != nil {
		return err
	}

	// cert in der format
	err = ioutil.WriteFile(flag.Arg(0)+".crt", cert[0].Raw, 0666)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	b, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return err
	}

	err = pem.Encode(buf, &pem.Block{Type: "PRIVATE KEY", Bytes: b})
	if err != nil {
		return err
	}

	err = pem.Encode(buf, &pem.Block{Type: "CERTIFICATE", Bytes: cert[0].Raw})
	if err != nil {
		return err
	}

	// key and cert in PKCS#8 PEM format for Azure Key Vault.
	return ioutil.WriteFile(flag.Arg(0)+".pem", buf.Bytes(), 0600)
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "usage: %s commonName...\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	if err := run(); err != nil {
		panic(err)
	}
}
