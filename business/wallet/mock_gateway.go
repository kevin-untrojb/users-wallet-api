// Code generated by MockGen. DO NOT EDIT.
// Source: gateway.go

// Package wallet is a generated GoMock package.
package wallet

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGateway is a mock of Gateway interface.
type MockGateway struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayMockRecorder
}

// MockGatewayMockRecorder is the mock recorder for MockGateway.
type MockGatewayMockRecorder struct {
	mock *MockGateway
}

// NewMockGateway creates a new mock instance.
func NewMockGateway(ctrl *gomock.Controller) *MockGateway {
	mock := &MockGateway{ctrl: ctrl}
	mock.recorder = &MockGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGateway) EXPECT() *MockGatewayMockRecorder {
	return m.recorder
}

// GetWalletsFroUser mocks base method.
func (m *MockGateway) GetWalletsFroUser(arg0 context.Context, arg1 int64) ([]Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletsFroUser", arg0, arg1)
	ret0, _ := ret[0].([]Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletsFroUser indicates an expected call of GetWalletsFroUser.
func (mr *MockGatewayMockRecorder) GetWalletsFroUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletsFroUser", reflect.TypeOf((*MockGateway)(nil).GetWalletsFroUser), arg0, arg1)
}

// NewTransaction mocks base method.
func (m *MockGateway) NewTransaction(ctx context.Context, nt Transaction) (Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTransaction", ctx, nt)
	ret0, _ := ret[0].(Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTransaction indicates an expected call of NewTransaction.
func (mr *MockGatewayMockRecorder) NewTransaction(ctx, nt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTransaction", reflect.TypeOf((*MockGateway)(nil).NewTransaction), ctx, nt)
}

// SearchTransactionsForUser mocks base method.
func (m *MockGateway) SearchTransactionsForUser(ctx context.Context, userID int64, params *SearchRequestParams) (SearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTransactionsForUser", ctx, userID, params)
	ret0, _ := ret[0].(SearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTransactionsForUser indicates an expected call of SearchTransactionsForUser.
func (mr *MockGatewayMockRecorder) SearchTransactionsForUser(ctx, userID, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTransactionsForUser", reflect.TypeOf((*MockGateway)(nil).SearchTransactionsForUser), ctx, userID, params)
}
