// Code generated by MockGen. DO NOT EDIT.
// Source: account.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccount is a mock of Account interface.
type MockAccount struct {
	ctrl     *gomock.Controller
	recorder *MockAccountMockRecorder
}

// MockAccountMockRecorder is the mock recorder for MockAccount.
type MockAccountMockRecorder struct {
	mock *MockAccount
}

// NewMockAccount creates a new mock instance.
func NewMockAccount(ctrl *gomock.Controller) *MockAccount {
	mock := &MockAccount{ctrl: ctrl}
	mock.recorder = &MockAccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccount) EXPECT() *MockAccountMockRecorder {
	return m.recorder
}

// HashedPass mocks base method.
func (m *MockAccount) HashedPass() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashedPass")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashedPass indicates an expected call of HashedPass.
func (mr *MockAccountMockRecorder) HashedPass() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashedPass", reflect.TypeOf((*MockAccount)(nil).HashedPass))
}

// Username mocks base method.
func (m *MockAccount) Username() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Username")
	ret0, _ := ret[0].(string)
	return ret0
}

// Username indicates an expected call of Username.
func (mr *MockAccountMockRecorder) Username() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Username", reflect.TypeOf((*MockAccount)(nil).Username))
}

// Verify mocks base method.
func (m *MockAccount) Verify(password string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", password)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockAccountMockRecorder) Verify(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockAccount)(nil).Verify), password)
}
