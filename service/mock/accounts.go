// Code generated by MockGen. DO NOT EDIT.
// Source: accounts.go
//
// Generated by this command:
//
//	mockgen -source accounts.go -package mock -destination mock/accounts.go Accounts
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	service "github.com/jakobmoellerdev/octi-sync-server/service"
	gomock "go.uber.org/mock/gomock"
)

// MockAccounts is a mock of Accounts interface.
type MockAccounts struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsMockRecorder
}

// MockAccountsMockRecorder is the mock recorder for MockAccounts.
type MockAccountsMockRecorder struct {
	mock *MockAccounts
}

// NewMockAccounts creates a new mock instance.
func NewMockAccounts(ctrl *gomock.Controller) *MockAccounts {
	mock := &MockAccounts{ctrl: ctrl}
	mock.recorder = &MockAccountsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccounts) EXPECT() *MockAccountsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccounts) Create(ctx context.Context, username string) (service.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, username)
	ret0, _ := ret[0].(service.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAccountsMockRecorder) Create(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccounts)(nil).Create), ctx, username)
}

// Find mocks base method.
func (m *MockAccounts) Find(ctx context.Context, username string) (service.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, username)
	ret0, _ := ret[0].(service.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockAccountsMockRecorder) Find(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockAccounts)(nil).Find), ctx, username)
}

// HealthCheck mocks base method.
func (m *MockAccounts) HealthCheck() service.HealthCheck {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(service.HealthCheck)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockAccountsMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockAccounts)(nil).HealthCheck))
}
