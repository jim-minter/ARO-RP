package acr

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"net/http"
	"net/http/httputil"

	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/util/pullsecret"
)

type acr struct {
	log           *logrus.Entry
	hostname      string
	authorization string
}

func New(log *logrus.Entry, hostname string) (http.Handler, error) {
	authorization, err := pullsecret.Authorization(hostname)
	if err != nil {
		return nil, err
	}

	a := &acr{
		log:           log,
		hostname:      hostname,
		authorization: authorization,
	}

	return &httputil.ReverseProxy{
		Director:  a.director,
		Transport: a,
	}, nil
}

func (a *acr) director(r *http.Request) {
	r.URL.Scheme = "https"
	r.URL.Host = a.hostname
	r.Host = ""
	r.RequestURI = ""

	if a.authorization != "" {
		r.Header.Set("Authorization", "Basic "+a.authorization)
	}
}

func (a *acr) RoundTrip(r *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(r)
}
