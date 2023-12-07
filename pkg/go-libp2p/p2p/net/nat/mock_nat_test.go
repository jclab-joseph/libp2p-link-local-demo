// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libp2p/go-nat (interfaces: NAT)
//
// Generated by this command:
//
//	mockgen -package nat -destination mock_nat_test.go github.com/libp2p/go-nat NAT
//
// Package nat is a generated GoMock package.
package nat

import (
	context "context"
	net "net"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockNAT is a mock of NAT interface.
type MockNAT struct {
	ctrl     *gomock.Controller
	recorder *MockNATMockRecorder
}

// MockNATMockRecorder is the mock recorder for MockNAT.
type MockNATMockRecorder struct {
	mock *MockNAT
}

// NewMockNAT creates a new mock instance.
func NewMockNAT(ctrl *gomock.Controller) *MockNAT {
	mock := &MockNAT{ctrl: ctrl}
	mock.recorder = &MockNATMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNAT) EXPECT() *MockNATMockRecorder {
	return m.recorder
}

// AddPortMapping mocks base method.
func (m *MockNAT) AddPortMapping(arg0 context.Context, arg1 string, arg2 int, arg3 string, arg4 time.Duration) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPortMapping", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPortMapping indicates an expected call of AddPortMapping.
func (mr *MockNATMockRecorder) AddPortMapping(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPortMapping", reflect.TypeOf((*MockNAT)(nil).AddPortMapping), arg0, arg1, arg2, arg3, arg4)
}

// DeletePortMapping mocks base method.
func (m *MockNAT) DeletePortMapping(arg0 context.Context, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePortMapping", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePortMapping indicates an expected call of DeletePortMapping.
func (mr *MockNATMockRecorder) DeletePortMapping(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePortMapping", reflect.TypeOf((*MockNAT)(nil).DeletePortMapping), arg0, arg1, arg2)
}

// GetDeviceAddress mocks base method.
func (m *MockNAT) GetDeviceAddress() (net.IP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceAddress")
	ret0, _ := ret[0].(net.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceAddress indicates an expected call of GetDeviceAddress.
func (mr *MockNATMockRecorder) GetDeviceAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceAddress", reflect.TypeOf((*MockNAT)(nil).GetDeviceAddress))
}

// GetExternalAddress mocks base method.
func (m *MockNAT) GetExternalAddress() (net.IP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalAddress")
	ret0, _ := ret[0].(net.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalAddress indicates an expected call of GetExternalAddress.
func (mr *MockNATMockRecorder) GetExternalAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalAddress", reflect.TypeOf((*MockNAT)(nil).GetExternalAddress))
}

// GetInternalAddress mocks base method.
func (m *MockNAT) GetInternalAddress() (net.IP, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInternalAddress")
	ret0, _ := ret[0].(net.IP)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInternalAddress indicates an expected call of GetInternalAddress.
func (mr *MockNATMockRecorder) GetInternalAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInternalAddress", reflect.TypeOf((*MockNAT)(nil).GetInternalAddress))
}

// Type mocks base method.
func (m *MockNAT) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockNATMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockNAT)(nil).Type))
}
