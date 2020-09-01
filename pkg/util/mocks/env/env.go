// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/env (interfaces: Interface)

// Package mock_env is a generated GoMock package.
package mock_env

import (
	context "context"
	rsa "crypto/rsa"
	x509 "crypto/x509"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	env "github.com/Azure/ARO-RP/pkg/env"
	refreshable "github.com/Azure/ARO-RP/pkg/util/refreshable"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// ACRName mocks base method
func (m *MockInterface) ACRName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ACRName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ACRName indicates an expected call of ACRName
func (mr *MockInterfaceMockRecorder) ACRName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ACRName", reflect.TypeOf((*MockInterface)(nil).ACRName))
}

// ACRResourceID mocks base method
func (m *MockInterface) ACRResourceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ACRResourceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ACRResourceID indicates an expected call of ACRResourceID
func (mr *MockInterfaceMockRecorder) ACRResourceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ACRResourceID", reflect.TypeOf((*MockInterface)(nil).ACRResourceID))
}

// ClustersGenevaLoggingSecret mocks base method
func (m *MockInterface) ClustersGenevaLoggingSecret() (*rsa.PrivateKey, *x509.Certificate) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClustersGenevaLoggingSecret")
	ret0, _ := ret[0].(*rsa.PrivateKey)
	ret1, _ := ret[1].(*x509.Certificate)
	return ret0, ret1
}

// ClustersGenevaLoggingSecret indicates an expected call of ClustersGenevaLoggingSecret
func (mr *MockInterfaceMockRecorder) ClustersGenevaLoggingSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClustersGenevaLoggingSecret", reflect.TypeOf((*MockInterface)(nil).ClustersGenevaLoggingSecret))
}

// ClustersKeyvaultURI mocks base method
func (m *MockInterface) ClustersKeyvaultURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClustersKeyvaultURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClustersKeyvaultURI indicates an expected call of ClustersKeyvaultURI
func (mr *MockInterfaceMockRecorder) ClustersKeyvaultURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClustersKeyvaultURI", reflect.TypeOf((*MockInterface)(nil).ClustersKeyvaultURI))
}

// Domain mocks base method
func (m *MockInterface) Domain() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Domain")
	ret0, _ := ret[0].(string)
	return ret0
}

// Domain indicates an expected call of Domain
func (mr *MockInterfaceMockRecorder) Domain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Domain", reflect.TypeOf((*MockInterface)(nil).Domain))
}

// FPAuthorizer mocks base method
func (m *MockInterface) FPAuthorizer(arg0, arg1 string) (refreshable.Authorizer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FPAuthorizer", arg0, arg1)
	ret0, _ := ret[0].(refreshable.Authorizer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FPAuthorizer indicates an expected call of FPAuthorizer
func (mr *MockInterfaceMockRecorder) FPAuthorizer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FPAuthorizer", reflect.TypeOf((*MockInterface)(nil).FPAuthorizer), arg0, arg1)
}

// GetCertificateSecret mocks base method
func (m *MockInterface) GetCertificateSecret(arg0 context.Context, arg1 string) (*rsa.PrivateKey, []*x509.Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateSecret", arg0, arg1)
	ret0, _ := ret[0].(*rsa.PrivateKey)
	ret1, _ := ret[1].([]*x509.Certificate)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCertificateSecret indicates an expected call of GetCertificateSecret
func (mr *MockInterfaceMockRecorder) GetCertificateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateSecret", reflect.TypeOf((*MockInterface)(nil).GetCertificateSecret), arg0, arg1)
}

// GetSecret mocks base method
func (m *MockInterface) GetSecret(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockInterfaceMockRecorder) GetSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockInterface)(nil).GetSecret), arg0, arg1)
}

// Location mocks base method
func (m *MockInterface) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location
func (mr *MockInterfaceMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockInterface)(nil).Location))
}

// ManagedDomain mocks base method
func (m *MockInterface) ManagedDomain(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManagedDomain", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ManagedDomain indicates an expected call of ManagedDomain
func (mr *MockInterfaceMockRecorder) ManagedDomain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManagedDomain", reflect.TypeOf((*MockInterface)(nil).ManagedDomain), arg0)
}

// ResourceGroup mocks base method
func (m *MockInterface) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup
func (mr *MockInterfaceMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockInterface)(nil).ResourceGroup))
}

// SubscriptionID mocks base method
func (m *MockInterface) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID
func (mr *MockInterfaceMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockInterface)(nil).SubscriptionID))
}

// TenantID mocks base method
func (m *MockInterface) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID
func (mr *MockInterfaceMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockInterface)(nil).TenantID))
}

// Type mocks base method
func (m *MockInterface) Type() env.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(env.Type)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockInterfaceMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockInterface)(nil).Type))
}

// Zones mocks base method
func (m *MockInterface) Zones(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Zones", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Zones indicates an expected call of Zones
func (mr *MockInterfaceMockRecorder) Zones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Zones", reflect.TypeOf((*MockInterface)(nil).Zones), arg0)
}
