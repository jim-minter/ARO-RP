// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/azureclient/storage (interfaces: BaseClient)

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	io "io"
	reflect "reflect"

	storage "github.com/Azure/azure-sdk-for-go/storage"
	gomock "github.com/golang/mock/gomock"
)

// MockBaseClient is a mock of BaseClient interface
type MockBaseClient struct {
	ctrl     *gomock.Controller
	recorder *MockBaseClientMockRecorder
}

// MockBaseClientMockRecorder is the mock recorder for MockBaseClient
type MockBaseClientMockRecorder struct {
	mock *MockBaseClient
}

// NewMockBaseClient creates a new mock instance
func NewMockBaseClient(ctrl *gomock.Controller) *MockBaseClient {
	mock := &MockBaseClient{ctrl: ctrl}
	mock.recorder = &MockBaseClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBaseClient) EXPECT() *MockBaseClientMockRecorder {
	return m.recorder
}

// CreateBlockBlobFromReader mocks base method
func (m *MockBaseClient) CreateBlockBlobFromReader(arg0 io.Reader, arg1 *storage.PutBlobOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBlockBlobFromReader", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBlockBlobFromReader indicates an expected call of CreateBlockBlobFromReader
func (mr *MockBaseClientMockRecorder) CreateBlockBlobFromReader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlockBlobFromReader", reflect.TypeOf((*MockBaseClient)(nil).CreateBlockBlobFromReader), arg0, arg1)
}

// CreateIfNotExists mocks base method
func (m *MockBaseClient) CreateIfNotExists(arg0 *storage.CreateContainerOptions) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIfNotExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIfNotExists indicates an expected call of CreateIfNotExists
func (mr *MockBaseClientMockRecorder) CreateIfNotExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIfNotExists", reflect.TypeOf((*MockBaseClient)(nil).CreateIfNotExists), arg0)
}

// GetBlobReference mocks base method
func (m *MockBaseClient) GetBlobReference(arg0 string) *storage.Blob {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlobReference", arg0)
	ret0, _ := ret[0].(*storage.Blob)
	return ret0
}

// GetBlobReference indicates an expected call of GetBlobReference
func (mr *MockBaseClientMockRecorder) GetBlobReference(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlobReference", reflect.TypeOf((*MockBaseClient)(nil).GetBlobReference), arg0)
}

// GetContainerReference mocks base method
func (m *MockBaseClient) GetContainerReference(arg0 string) *storage.Container {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerReference", arg0)
	ret0, _ := ret[0].(*storage.Container)
	return ret0
}

// GetContainerReference indicates an expected call of GetContainerReference
func (mr *MockBaseClientMockRecorder) GetContainerReference(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerReference", reflect.TypeOf((*MockBaseClient)(nil).GetContainerReference), arg0)
}
