// Code generated by MockGen. DO NOT EDIT.
// Source: devices.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	service "github.com/jakob-moeller-cloud/octi-sync-server/service"
)

// MockDevices is a mock of Devices interface.
type MockDevices struct {
	ctrl     *gomock.Controller
	recorder *MockDevicesMockRecorder
}

// MockDevicesMockRecorder is the mock recorder for MockDevices.
type MockDevicesMockRecorder struct {
	mock *MockDevices
}

// NewMockDevices creates a new mock instance.
func NewMockDevices(ctrl *gomock.Controller) *MockDevices {
	mock := &MockDevices{ctrl: ctrl}
	mock.recorder = &MockDevicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDevices) EXPECT() *MockDevicesMockRecorder {
	return m.recorder
}

// FindByAccount mocks base method.
func (m *MockDevices) FindByAccount(ctx context.Context, acc service.Account) ([]service.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAccount", ctx, acc)
	ret0, _ := ret[0].([]service.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAccount indicates an expected call of FindByAccount.
func (mr *MockDevicesMockRecorder) FindByAccount(ctx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAccount", reflect.TypeOf((*MockDevices)(nil).FindByAccount), ctx, acc)
}

// FindByDeviceID mocks base method.
func (m *MockDevices) FindByDeviceID(ctx context.Context, acc service.Account, deviceId service.DeviceID) (service.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByDeviceID", ctx, acc, deviceId)
	ret0, _ := ret[0].(service.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByDeviceID indicates an expected call of FindByDeviceID.
func (mr *MockDevicesMockRecorder) FindByDeviceID(ctx, acc, deviceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByDeviceID", reflect.TypeOf((*MockDevices)(nil).FindByDeviceID), ctx, acc, deviceId)
}

// HealthCheck mocks base method.
func (m *MockDevices) HealthCheck() service.HealthCheck {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(service.HealthCheck)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockDevicesMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockDevices)(nil).HealthCheck))
}

// Register mocks base method.
func (m *MockDevices) Register(ctx context.Context, acc service.Account, deviceId service.DeviceID) (service.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, acc, deviceId)
	ret0, _ := ret[0].(service.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockDevicesMockRecorder) Register(ctx, acc, deviceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockDevices)(nil).Register), ctx, acc, deviceId)
}
