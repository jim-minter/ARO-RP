package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"bytes"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/database"
	"github.com/Azure/ARO-RP/pkg/database/cosmosdb"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/util/clientauthorizer"
	"github.com/Azure/ARO-RP/pkg/util/deployment"
	mock_env "github.com/Azure/ARO-RP/pkg/util/mocks/env"
	utiltls "github.com/Azure/ARO-RP/pkg/util/tls"
	testdatabase "github.com/Azure/ARO-RP/test/database"
	testclusterdata "github.com/Azure/ARO-RP/test/util/clusterdata"
	"github.com/Azure/ARO-RP/test/util/listener"
)

var (
	serverkey, clientkey     *rsa.PrivateKey
	servercerts, clientcerts []*x509.Certificate
)

func init() {
	var err error

	clientkey, clientcerts, err = utiltls.GenerateKeyAndCertificate([]string{"client"}, nil, nil, false, true)
	if err != nil {
		panic(err)
	}

	serverkey, servercerts, err = utiltls.GenerateKeyAndCertificate([]string{"server"}, nil, nil, false, false)
	if err != nil {
		panic(err)
	}
}

type testInfra struct {
	env        env.Interface
	controller *gomock.Controller
	l          net.Listener
	cli        *http.Client
	enricher   testclusterdata.TestEnricher
	log        *logrus.Entry
	fixture    *testdatabase.Fixture
	checker    *testdatabase.Checker

	openShiftClustersClient   *cosmosdb.FakeOpenShiftClusterDocumentClient
	openShiftClustersDatabase database.OpenShiftClusters
	asyncOperationsClient     *cosmosdb.FakeAsyncOperationDocumentClient
	asyncOperationsDatabase   database.AsyncOperations
	billingClient             *cosmosdb.FakeBillingDocumentClient
	billingDatabase           database.Billing
	subscriptionsClient       *cosmosdb.FakeSubscriptionDocumentClient
	subscriptionsDatabase     database.Subscriptions
}

func newTestInfra(t *testing.T) *testInfra {
	pool := x509.NewCertPool()
	pool.AddCert(servercerts[0])

	l := listener.NewListener()

	controller := gomock.NewController(t)

	_env := mock_env.NewMockInterface(controller)
	_env.EXPECT().DeploymentMode().AnyTimes().Return(deployment.Production)
	_env.EXPECT().Location().AnyTimes().Return("eastus")
	_env.EXPECT().GetCertificateSecret(gomock.Any(), env.RPServerSecretName).AnyTimes().Return(serverkey, servercerts, nil)
	_env.EXPECT().ArmClientAuthorizer().AnyTimes().Return(clientauthorizer.NewOne(clientcerts[0].Raw))
	_env.EXPECT().AdminClientAuthorizer().AnyTimes().Return(clientauthorizer.NewOne(clientcerts[0].Raw))
	_env.EXPECT().Domain().AnyTimes().Return("")
	_env.EXPECT().Listen().AnyTimes().Return(l, nil)

	log := logrus.NewEntry(logrus.StandardLogger())

	fixture := testdatabase.NewFixture()
	checker := testdatabase.NewChecker()

	return &testInfra{
		env:        _env,
		controller: controller,
		l:          l,
		enricher:   testclusterdata.NewTestEnricher(),
		fixture:    fixture,
		checker:    checker,
		log:        log,
		cli: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: pool,
					Certificates: []tls.Certificate{
						{
							Certificate: [][]byte{clientcerts[0].Raw},
							PrivateKey:  clientkey,
						},
					},
				},
				Dial: l.Dial,
			},
		},
	}
}

func (ti *testInfra) WithOpenShiftClusters() *testInfra {
	ti.openShiftClustersDatabase, ti.openShiftClustersClient = testdatabase.NewFakeOpenShiftClusters()
	ti.fixture.WithOpenShiftClusters(ti.openShiftClustersDatabase)
	return ti
}

func (ti *testInfra) WithBilling() *testInfra {
	ti.billingDatabase, ti.billingClient = testdatabase.NewFakeBilling()
	ti.fixture.WithBilling(ti.billingDatabase)
	return ti
}

func (ti *testInfra) WithSubscriptions() *testInfra {
	ti.subscriptionsDatabase, ti.subscriptionsClient = testdatabase.NewFakeSubscriptions()
	ti.fixture.WithSubscriptions(ti.subscriptionsDatabase)
	return ti
}

func (ti *testInfra) WithAsyncOperations() *testInfra {
	ti.asyncOperationsDatabase, ti.asyncOperationsClient = testdatabase.NewFakeAsyncOperations()
	ti.fixture.WithAsyncOperations(ti.asyncOperationsDatabase)
	return ti
}

func (ti *testInfra) done() error {
	ti.controller.Finish()
	ti.cli.CloseIdleConnections()
	return ti.l.Close()
}

func (ti *testInfra) buildFixtures(fixtures func(*testdatabase.Fixture)) error {
	if fixtures != nil {
		fixtures(ti.fixture)
	}
	return ti.fixture.Create()
}

func (ti *testInfra) request(method, url string, header http.Header, in interface{}) (*http.Response, []byte, error) {
	var b []byte

	if in != nil {
		var err error
		b, err = json.Marshal(in)
		if err != nil {
			return nil, nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return nil, nil, err
	}

	req.Header = header

	resp, err := ti.cli.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, b, nil
}

func validateResponse(resp *http.Response, b []byte, wantStatusCode int, wantError string, wantResponse interface{}) error {
	if resp.StatusCode != wantStatusCode {
		return fmt.Errorf("unexpected status code %d, wanted %d", resp.StatusCode, wantStatusCode)
	}

	if wantError != "" {
		cloudErr := &api.CloudError{StatusCode: resp.StatusCode}
		err := json.Unmarshal(b, &cloudErr)
		if err != nil {
			return err
		}

		if cloudErr.Error() != wantError {
			return fmt.Errorf("unexpected error %s, wanted %s", cloudErr.Error(), wantError)
		}

		return nil
	}

	if wantResponse == nil || reflect.ValueOf(wantResponse).IsZero() {
		if len(b) != 0 {
			return fmt.Errorf("unexpected response %s, wanted no content", string(b))
		}
		return nil
	}

	if wantResponse, ok := wantResponse.([]byte); ok {
		if !bytes.Equal(b, wantResponse) {
			return fmt.Errorf("unexpected response %s, wanted %s", string(b), string(wantResponse))
		}
		return nil
	}

	v := reflect.New(reflect.TypeOf(wantResponse).Elem()).Interface()
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	if diff := deep.Equal(v, wantResponse); diff != nil {
		return fmt.Errorf("unexpected response %s, wanted to match %#v (%s)", string(b), wantResponse, diff)
	}

	return nil
}
