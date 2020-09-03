// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/database (interfaces: AsyncOperations,Billing,OpenShiftClusters,Subscriptions)

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/Azure/ARO-RP/pkg/api"
	cosmosdb "github.com/Azure/ARO-RP/pkg/database/cosmosdb"
)

// MockAsyncOperations is a mock of AsyncOperations interface
type MockAsyncOperations struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncOperationsMockRecorder
}

// MockAsyncOperationsMockRecorder is the mock recorder for MockAsyncOperations
type MockAsyncOperationsMockRecorder struct {
	mock *MockAsyncOperations
}

// NewMockAsyncOperations creates a new mock instance
func NewMockAsyncOperations(ctrl *gomock.Controller) *MockAsyncOperations {
	mock := &MockAsyncOperations{ctrl: ctrl}
	mock.recorder = &MockAsyncOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAsyncOperations) EXPECT() *MockAsyncOperationsMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAsyncOperations) Create(arg0 context.Context, arg1 *api.AsyncOperationDocument) (*api.AsyncOperationDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*api.AsyncOperationDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAsyncOperationsMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAsyncOperations)(nil).Create), arg0, arg1)
}

// Get mocks base method
func (m *MockAsyncOperations) Get(arg0 context.Context, arg1 string) (*api.AsyncOperationDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*api.AsyncOperationDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAsyncOperationsMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAsyncOperations)(nil).Get), arg0, arg1)
}

// Patch mocks base method
func (m *MockAsyncOperations) Patch(arg0 context.Context, arg1 string, arg2 func(*api.AsyncOperationDocument) error) (*api.AsyncOperationDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.AsyncOperationDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockAsyncOperationsMockRecorder) Patch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockAsyncOperations)(nil).Patch), arg0, arg1, arg2)
}

// MockBilling is a mock of Billing interface
type MockBilling struct {
	ctrl     *gomock.Controller
	recorder *MockBillingMockRecorder
}

// MockBillingMockRecorder is the mock recorder for MockBilling
type MockBillingMockRecorder struct {
	mock *MockBilling
}

// NewMockBilling creates a new mock instance
func NewMockBilling(ctrl *gomock.Controller) *MockBilling {
	mock := &MockBilling{ctrl: ctrl}
	mock.recorder = &MockBillingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBilling) EXPECT() *MockBillingMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockBilling) Create(arg0 context.Context, arg1 *api.BillingDocument) (*api.BillingDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*api.BillingDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockBillingMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBilling)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockBilling) Delete(arg0 context.Context, arg1 *api.BillingDocument) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockBillingMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBilling)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockBilling) Get(arg0 context.Context, arg1 string) (*api.BillingDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*api.BillingDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockBillingMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBilling)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockBilling) List(arg0 string) cosmosdb.BillingDocumentIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(cosmosdb.BillingDocumentIterator)
	return ret0
}

// List indicates an expected call of List
func (mr *MockBillingMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockBilling)(nil).List), arg0)
}

// ListAll mocks base method
func (m *MockBilling) ListAll(arg0 context.Context) (*api.BillingDocuments, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(*api.BillingDocuments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll
func (mr *MockBillingMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockBilling)(nil).ListAll), arg0)
}

// MarkForDeletion mocks base method
func (m *MockBilling) MarkForDeletion(arg0 context.Context, arg1 string) (*api.BillingDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkForDeletion", arg0, arg1)
	ret0, _ := ret[0].(*api.BillingDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkForDeletion indicates an expected call of MarkForDeletion
func (mr *MockBillingMockRecorder) MarkForDeletion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkForDeletion", reflect.TypeOf((*MockBilling)(nil).MarkForDeletion), arg0, arg1)
}

// UpdateLastBillingTimestamp mocks base method
func (m *MockBilling) UpdateLastBillingTimestamp(arg0 context.Context, arg1 string, arg2 int) (*api.BillingDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastBillingTimestamp", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.BillingDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLastBillingTimestamp indicates an expected call of UpdateLastBillingTimestamp
func (mr *MockBillingMockRecorder) UpdateLastBillingTimestamp(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastBillingTimestamp", reflect.TypeOf((*MockBilling)(nil).UpdateLastBillingTimestamp), arg0, arg1, arg2)
}

// MockOpenShiftClusters is a mock of OpenShiftClusters interface
type MockOpenShiftClusters struct {
	ctrl     *gomock.Controller
	recorder *MockOpenShiftClustersMockRecorder
}

// MockOpenShiftClustersMockRecorder is the mock recorder for MockOpenShiftClusters
type MockOpenShiftClustersMockRecorder struct {
	mock *MockOpenShiftClusters
}

// NewMockOpenShiftClusters creates a new mock instance
func NewMockOpenShiftClusters(ctrl *gomock.Controller) *MockOpenShiftClusters {
	mock := &MockOpenShiftClusters{ctrl: ctrl}
	mock.recorder = &MockOpenShiftClustersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOpenShiftClusters) EXPECT() *MockOpenShiftClustersMockRecorder {
	return m.recorder
}

// ChangeFeed mocks base method
func (m *MockOpenShiftClusters) ChangeFeed() cosmosdb.OpenShiftClusterDocumentIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeFeed")
	ret0, _ := ret[0].(cosmosdb.OpenShiftClusterDocumentIterator)
	return ret0
}

// ChangeFeed indicates an expected call of ChangeFeed
func (mr *MockOpenShiftClustersMockRecorder) ChangeFeed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeFeed", reflect.TypeOf((*MockOpenShiftClusters)(nil).ChangeFeed))
}

// Create mocks base method
func (m *MockOpenShiftClusters) Create(arg0 context.Context, arg1 *api.OpenShiftClusterDocument) (*api.OpenShiftClusterDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*api.OpenShiftClusterDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOpenShiftClustersMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOpenShiftClusters)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockOpenShiftClusters) Delete(arg0 context.Context, arg1 *api.OpenShiftClusterDocument) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockOpenShiftClustersMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOpenShiftClusters)(nil).Delete), arg0, arg1)
}

// Dequeue mocks base method
func (m *MockOpenShiftClusters) Dequeue(arg0 context.Context) (*api.OpenShiftClusterDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dequeue", arg0)
	ret0, _ := ret[0].(*api.OpenShiftClusterDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dequeue indicates an expected call of Dequeue
func (mr *MockOpenShiftClustersMockRecorder) Dequeue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dequeue", reflect.TypeOf((*MockOpenShiftClusters)(nil).Dequeue), arg0)
}

// EndLease mocks base method
func (m *MockOpenShiftClusters) EndLease(arg0 context.Context, arg1 string, arg2, arg3 api.ProvisioningState, arg4 *string) (*api.OpenShiftClusterDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndLease", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*api.OpenShiftClusterDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndLease indicates an expected call of EndLease
func (mr *MockOpenShiftClustersMockRecorder) EndLease(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndLease", reflect.TypeOf((*MockOpenShiftClusters)(nil).EndLease), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method
func (m *MockOpenShiftClusters) Get(arg0 context.Context, arg1 string) (*api.OpenShiftClusterDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*api.OpenShiftClusterDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockOpenShiftClustersMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOpenShiftClusters)(nil).Get), arg0, arg1)
}

// GetByClientID mocks base method
func (m *MockOpenShiftClusters) GetByClientID(arg0 context.Context, arg1, arg2 string) (*api.OpenShiftClusterDocuments, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByClientID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.OpenShiftClusterDocuments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByClientID indicates an expected call of GetByClientID
func (mr *MockOpenShiftClustersMockRecorder) GetByClientID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByClientID", reflect.TypeOf((*MockOpenShiftClusters)(nil).GetByClientID), arg0, arg1, arg2)
}

// GetByClusterResourceGroupID mocks base method
func (m *MockOpenShiftClusters) GetByClusterResourceGroupID(arg0 context.Context, arg1, arg2 string) (*api.OpenShiftClusterDocuments, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByClusterResourceGroupID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.OpenShiftClusterDocuments)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByClusterResourceGroupID indicates an expected call of GetByClusterResourceGroupID
func (mr *MockOpenShiftClustersMockRecorder) GetByClusterResourceGroupID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByClusterResourceGroupID", reflect.TypeOf((*MockOpenShiftClusters)(nil).GetByClusterResourceGroupID), arg0, arg1, arg2)
}

// Lease mocks base method
func (m *MockOpenShiftClusters) Lease(arg0 context.Context, arg1 string) (*api.OpenShiftClusterDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lease", arg0, arg1)
	ret0, _ := ret[0].(*api.OpenShiftClusterDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lease indicates an expected call of Lease
func (mr *MockOpenShiftClustersMockRecorder) Lease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lease", reflect.TypeOf((*MockOpenShiftClusters)(nil).Lease), arg0, arg1)
}

// List mocks base method
func (m *MockOpenShiftClusters) List() cosmosdb.OpenShiftClusterDocumentIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(cosmosdb.OpenShiftClusterDocumentIterator)
	return ret0
}

// List indicates an expected call of List
func (mr *MockOpenShiftClustersMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOpenShiftClusters)(nil).List))
}

// ListByPrefix mocks base method
func (m *MockOpenShiftClusters) ListByPrefix(arg0, arg1, arg2 string) (cosmosdb.OpenShiftClusterDocumentIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByPrefix", arg0, arg1, arg2)
	ret0, _ := ret[0].(cosmosdb.OpenShiftClusterDocumentIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByPrefix indicates an expected call of ListByPrefix
func (mr *MockOpenShiftClustersMockRecorder) ListByPrefix(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByPrefix", reflect.TypeOf((*MockOpenShiftClusters)(nil).ListByPrefix), arg0, arg1, arg2)
}

// Patch mocks base method
func (m *MockOpenShiftClusters) Patch(arg0 context.Context, arg1 string, arg2 func(*api.OpenShiftClusterDocument) error) (*api.OpenShiftClusterDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.OpenShiftClusterDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockOpenShiftClustersMockRecorder) Patch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockOpenShiftClusters)(nil).Patch), arg0, arg1, arg2)
}

// PatchWithLease mocks base method
func (m *MockOpenShiftClusters) PatchWithLease(arg0 context.Context, arg1 string, arg2 func(*api.OpenShiftClusterDocument) error) (*api.OpenShiftClusterDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchWithLease", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.OpenShiftClusterDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchWithLease indicates an expected call of PatchWithLease
func (mr *MockOpenShiftClustersMockRecorder) PatchWithLease(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchWithLease", reflect.TypeOf((*MockOpenShiftClusters)(nil).PatchWithLease), arg0, arg1, arg2)
}

// QueueLength mocks base method
func (m *MockOpenShiftClusters) QueueLength(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueLength", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueLength indicates an expected call of QueueLength
func (mr *MockOpenShiftClustersMockRecorder) QueueLength(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueLength", reflect.TypeOf((*MockOpenShiftClusters)(nil).QueueLength), arg0)
}

// Update mocks base method
func (m *MockOpenShiftClusters) Update(arg0 context.Context, arg1 *api.OpenShiftClusterDocument) (*api.OpenShiftClusterDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*api.OpenShiftClusterDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockOpenShiftClustersMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOpenShiftClusters)(nil).Update), arg0, arg1)
}

// MockSubscriptions is a mock of Subscriptions interface
type MockSubscriptions struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionsMockRecorder
}

// MockSubscriptionsMockRecorder is the mock recorder for MockSubscriptions
type MockSubscriptionsMockRecorder struct {
	mock *MockSubscriptions
}

// NewMockSubscriptions creates a new mock instance
func NewMockSubscriptions(ctrl *gomock.Controller) *MockSubscriptions {
	mock := &MockSubscriptions{ctrl: ctrl}
	mock.recorder = &MockSubscriptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubscriptions) EXPECT() *MockSubscriptionsMockRecorder {
	return m.recorder
}

// ChangeFeed mocks base method
func (m *MockSubscriptions) ChangeFeed() cosmosdb.SubscriptionDocumentIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeFeed")
	ret0, _ := ret[0].(cosmosdb.SubscriptionDocumentIterator)
	return ret0
}

// ChangeFeed indicates an expected call of ChangeFeed
func (mr *MockSubscriptionsMockRecorder) ChangeFeed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeFeed", reflect.TypeOf((*MockSubscriptions)(nil).ChangeFeed))
}

// Create mocks base method
func (m *MockSubscriptions) Create(arg0 context.Context, arg1 *api.SubscriptionDocument) (*api.SubscriptionDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*api.SubscriptionDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockSubscriptionsMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSubscriptions)(nil).Create), arg0, arg1)
}

// Dequeue mocks base method
func (m *MockSubscriptions) Dequeue(arg0 context.Context) (*api.SubscriptionDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dequeue", arg0)
	ret0, _ := ret[0].(*api.SubscriptionDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dequeue indicates an expected call of Dequeue
func (mr *MockSubscriptionsMockRecorder) Dequeue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dequeue", reflect.TypeOf((*MockSubscriptions)(nil).Dequeue), arg0)
}

// EndLease mocks base method
func (m *MockSubscriptions) EndLease(arg0 context.Context, arg1 string, arg2, arg3 bool) (*api.SubscriptionDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndLease", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*api.SubscriptionDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndLease indicates an expected call of EndLease
func (mr *MockSubscriptionsMockRecorder) EndLease(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndLease", reflect.TypeOf((*MockSubscriptions)(nil).EndLease), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockSubscriptions) Get(arg0 context.Context, arg1 string) (*api.SubscriptionDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*api.SubscriptionDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSubscriptionsMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubscriptions)(nil).Get), arg0, arg1)
}

// Lease mocks base method
func (m *MockSubscriptions) Lease(arg0 context.Context, arg1 string) (*api.SubscriptionDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lease", arg0, arg1)
	ret0, _ := ret[0].(*api.SubscriptionDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lease indicates an expected call of Lease
func (mr *MockSubscriptionsMockRecorder) Lease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lease", reflect.TypeOf((*MockSubscriptions)(nil).Lease), arg0, arg1)
}

// Update mocks base method
func (m *MockSubscriptions) Update(arg0 context.Context, arg1 *api.SubscriptionDocument) (*api.SubscriptionDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*api.SubscriptionDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockSubscriptionsMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSubscriptions)(nil).Update), arg0, arg1)
}
