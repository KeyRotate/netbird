// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/keyrotate/netbird/iface (interfaces: PacketFilter)

// Package mocks is a generated GoMock package.
package mocks

import (
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPacketFilter is a mock of PacketFilter interface.
type MockPacketFilter struct {
	ctrl     *gomock.Controller
	recorder *MockPacketFilterMockRecorder
}

// MockPacketFilterMockRecorder is the mock recorder for MockPacketFilter.
type MockPacketFilterMockRecorder struct {
	mock *MockPacketFilter
}

// NewMockPacketFilter creates a new mock instance.
func NewMockPacketFilter(ctrl *gomock.Controller) *MockPacketFilter {
	mock := &MockPacketFilter{ctrl: ctrl}
	mock.recorder = &MockPacketFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacketFilter) EXPECT() *MockPacketFilterMockRecorder {
	return m.recorder
}

// AddUDPPacketHook mocks base method.
func (m *MockPacketFilter) AddUDPPacketHook(arg0 bool, arg1 net.IP, arg2 uint16, arg3 func(*net.UDPAddr, []byte) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddUDPPacketHook", arg0, arg1, arg2, arg3)
}

// AddUDPPacketHook indicates an expected call of AddUDPPacketHook.
func (mr *MockPacketFilterMockRecorder) AddUDPPacketHook(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUDPPacketHook", reflect.TypeOf((*MockPacketFilter)(nil).AddUDPPacketHook), arg0, arg1, arg2, arg3)
}

// DropIncoming mocks base method.
func (m *MockPacketFilter) DropIncoming(arg0 []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropIncoming", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DropIncoming indicates an expected call of DropIncoming.
func (mr *MockPacketFilterMockRecorder) DropIncoming(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIncoming", reflect.TypeOf((*MockPacketFilter)(nil).DropIncoming), arg0)
}

// DropOutgoing mocks base method.
func (m *MockPacketFilter) DropOutgoing(arg0 []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropOutgoing", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DropOutgoing indicates an expected call of DropOutgoing.
func (mr *MockPacketFilterMockRecorder) DropOutgoing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropOutgoing", reflect.TypeOf((*MockPacketFilter)(nil).DropOutgoing), arg0)
}

// SetNetwork mocks base method.
func (m *MockPacketFilter) SetNetwork(arg0 *net.IPNet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNetwork", arg0)
}

// SetNetwork indicates an expected call of SetNetwork.
func (mr *MockPacketFilterMockRecorder) SetNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNetwork", reflect.TypeOf((*MockPacketFilter)(nil).SetNetwork), arg0)
}
