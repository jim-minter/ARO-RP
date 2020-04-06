// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/billing (interfaces: E2EManager)

// Package mock_billing is a generated GoMock package.
package mock_billing

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/Azure/ARO-RP/pkg/api"
	database "github.com/Azure/ARO-RP/pkg/database"
	env "github.com/Azure/ARO-RP/pkg/env"
)

// MockE2EManager is a mock of E2EManager interface
type MockE2EManager struct {
	ctrl     *gomock.Controller
	recorder *MockE2EManagerMockRecorder
}

// MockE2EManagerMockRecorder is the mock recorder for MockE2EManager
type MockE2EManagerMockRecorder struct {
	mock *MockE2EManager
}

// NewMockE2EManager creates a new mock instance
func NewMockE2EManager(ctrl *gomock.Controller) *MockE2EManager {
	mock := &MockE2EManager{ctrl: ctrl}
	mock.recorder = &MockE2EManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockE2EManager) EXPECT() *MockE2EManagerMockRecorder {
	return m.recorder
}

// CreateOrUpdateE2EBlob mocks base method
func (m *MockE2EManager) CreateOrUpdateE2EBlob(arg0 context.Context, arg1 env.Interface, arg2 database.Subscriptions, arg3 *api.BillingDocument) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateE2EBlob", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateE2EBlob indicates an expected call of CreateOrUpdateE2EBlob
func (mr *MockE2EManagerMockRecorder) CreateOrUpdateE2EBlob(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateE2EBlob", reflect.TypeOf((*MockE2EManager)(nil).CreateOrUpdateE2EBlob), arg0, arg1, arg2, arg3)
}
