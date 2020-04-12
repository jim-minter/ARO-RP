// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/frontend/kubeactions (interfaces: Interface)

// Package mock_kubeactions is a generated GoMock package.
package mock_kubeactions

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"

	api "github.com/Azure/ARO-RP/pkg/api"
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

// ClusterUpgrade mocks base method
func (m *MockInterface) ClusterUpgrade(arg0 context.Context, arg1 *api.OpenShiftCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterUpgrade", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClusterUpgrade indicates an expected call of ClusterUpgrade
func (mr *MockInterfaceMockRecorder) ClusterUpgrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterUpgrade", reflect.TypeOf((*MockInterface)(nil).ClusterUpgrade), arg0, arg1)
}

// Get mocks base method
func (m *MockInterface) Get(arg0 context.Context, arg1 *api.OpenShiftCluster, arg2, arg3, arg4 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockInterfaceMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// List mocks base method
func (m *MockInterface) List(arg0 context.Context, arg1 *api.OpenShiftCluster, arg2, arg3 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockInterfaceMockRecorder) List(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), arg0, arg1, arg2, arg3)
}

// MustGather mocks base method
func (m *MockInterface) MustGather(arg0 context.Context, arg1 *logrus.Entry, arg2 *api.OpenShiftCluster, arg3 io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustGather", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// MustGather indicates an expected call of MustGather
func (mr *MockInterfaceMockRecorder) MustGather(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustGather", reflect.TypeOf((*MockInterface)(nil).MustGather), arg0, arg1, arg2, arg3)
}
