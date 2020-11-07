package clientauthorizer

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"crypto/tls"
	"errors"
	"testing"

	"github.com/sirupsen/logrus"

	utiltls "github.com/Azure/ARO-RP/pkg/util/tls"
)

func TestAdminClientAuthorizer(t *testing.T) {
	caBundlePath := "/fake/path/to/ca/cert.pem"
	log := logrus.NewEntry(logrus.StandardLogger())

	validCaKey, validCaCerts, err := utiltls.GenerateKeyAndCertificate([]string{"validca"}, nil, nil, true, false)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range []struct {
		name string
		cs   func() (*tls.ConnectionState, error)
		want bool
	}{
		{
			name: "allow: single valid client certificate",
			want: true,
			cs: func() (*tls.ConnectionState, error) {
				_, validSingleClientCert, err := utiltls.GenerateKeyAndCertificate([]string{"validclient"}, validCaKey, validCaCerts[0], false, true)
				if err != nil {
					return nil, err
				}

				return &tls.ConnectionState{
					PeerCertificates: validSingleClientCert,
				}, nil
			},
		},
		{
			name: "allow: valid client certificate with intermediates",
			want: true,
			cs: func() (*tls.ConnectionState, error) {
				validIntermediateCaKey, validIntermediateCaCerts, err := utiltls.GenerateKeyAndCertificate([]string{"valid-intermediate-ca"}, validCaKey, validCaCerts[0], true, false)
				if err != nil {
					return nil, err
				}

				_, validCertWithIntermediates, err := utiltls.GenerateKeyAndCertificate([]string{"validclient"}, validIntermediateCaKey, validIntermediateCaCerts[0], false, true)
				if err != nil {
					return nil, err
				}
				validCertWithIntermediates = append(validCertWithIntermediates, validIntermediateCaCerts...)

				return &tls.ConnectionState{
					PeerCertificates: validCertWithIntermediates,
				}, nil
			},
		},
		{
			name: "deny: valid certificate with unexpected common name",
			cs: func() (*tls.ConnectionState, error) {
				_, invalidCommonNameClientCert, err := utiltls.GenerateKeyAndCertificate([]string{"invalidclient"}, validCaKey, validCaCerts[0], false, true)
				if err != nil {
					return nil, err
				}

				return &tls.ConnectionState{
					PeerCertificates: invalidCommonNameClientCert,
				}, nil
			},
		},
		{
			name: "deny: certificate with unexpected key usage",
			cs: func() (*tls.ConnectionState, error) {
				_, invalidKeyUsagesCert, err := utiltls.GenerateKeyAndCertificate([]string{"validclient"}, validCaKey, validCaCerts[0], false, false)
				if err != nil {
					return nil, err
				}

				return &tls.ConnectionState{
					PeerCertificates: invalidKeyUsagesCert,
				}, nil
			},
		},
		{
			name: "deny: matching common name, but unexpected ca",
			cs: func() (*tls.ConnectionState, error) {
				invalidCaKey, invalidCaCerts, err := utiltls.GenerateKeyAndCertificate([]string{"invalidca"}, nil, nil, true, false)
				if err != nil {
					return nil, err
				}

				_, invalidSigningCa, err := utiltls.GenerateKeyAndCertificate([]string{"validclient"}, invalidCaKey, invalidCaCerts[0], false, true)
				if err != nil {
					return nil, err
				}

				return &tls.ConnectionState{
					PeerCertificates: invalidSigningCa,
				}, nil
			},
		},
		{
			name: "deny: connection without client certificates",
			cs: func() (*tls.ConnectionState, error) {
				return &tls.ConnectionState{}, nil
			},
		},
		{
			name: "deny: nil connection state",
			cs: func() (*tls.ConnectionState, error) {
				return nil, nil
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			adminAuthorizer := &admin{
				clientCertCommonName: "validclient",

				log: log,
				readFile: func(filename string) ([]byte, error) {
					if filename != caBundlePath {
						t.Fatal(filename)
						return nil, errors.New("noop")
					}

					return utiltls.CertAsBytes(validCaCerts...)
				},
			}
			err := adminAuthorizer.readCABundle(caBundlePath)
			if err != nil {
				t.Fatal(err)
			}

			cs, err := tt.cs()
			if err != nil {
				t.Error(err)
			}

			result := adminAuthorizer.IsAuthorized(cs)
			if result != tt.want {
				t.Error(result)
			}
		})
	}
}

func TestAdminClientAuthorizerReadCABundle(t *testing.T) {
	validCaKey, validCaCerts, err := utiltls.GenerateKeyAndCertificate([]string{"validca"}, nil, nil, true, false)
	if err != nil {
		t.Fatal(err)
	}

	_, validClientCert, err := utiltls.GenerateKeyAndCertificate([]string{"validclient"}, validCaKey, validCaCerts[0], false, true)
	if err != nil {
		t.Fatal(err)
	}

	cs := &tls.ConnectionState{PeerCertificates: validClientCert}

	for _, tt := range []struct {
		name     string
		readFile func(string) ([]byte, error)
		want     bool
	}{
		{
			name: "valid ca cert",
			readFile: func(string) ([]byte, error) {
				return utiltls.CertAsBytes(validCaCerts...)
			},
			want: true,
		},
		{
			name: "error reading ca cert file",
			readFile: func(string) ([]byte, error) {
				return nil, errors.New("noop")
			},
		},
		{
			name: "error decoding ca cert file",
			readFile: func(string) ([]byte, error) {
				return []byte("invalid-ca-cert"), nil
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			adminAuthorizer := &admin{
				clientCertCommonName: "validclient",

				log:      logrus.NewEntry(logrus.StandardLogger()),
				readFile: tt.readFile,
			}

			if adminAuthorizer.IsAuthorized(cs) {
				t.Error("expected deny before the readCABundle call")
			}

			readCABundleErr := adminAuthorizer.readCABundle("/fake/path/to/ca/cert.pem")
			IsAuthorized := adminAuthorizer.IsAuthorized(cs)

			if tt.want {
				if readCABundleErr != nil {
					t.Error(readCABundleErr)
				}
				if !IsAuthorized {
					t.Error("expected to allow")
				}
			} else {
				if readCABundleErr == nil {
					t.Error("expected an error")
				}
				if IsAuthorized {
					t.Error("expected deny")
				}
			}
		})
	}
}
