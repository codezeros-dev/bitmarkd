// SPDX-License-Identifier: ISC
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bitmark-inc/bitmarkd/zmqutil (interfaces: ClientIntf)

// Package mocks is a generated GoMock package.
package mocks

import (
	util "github.com/bitmark-inc/bitmarkd/util"
	zmqutil "github.com/bitmark-inc/bitmarkd/zmqutil"
	gomock "github.com/golang/mock/gomock"
	zmq4 "github.com/pebbe/zmq4"
	reflect "reflect"
)

// MockClientIntf is a mock of ClientIntf interface
type MockClientIntf struct {
	ctrl     *gomock.Controller
	recorder *MockClientIntfMockRecorder
}

// MockClientIntfMockRecorder is the mock recorder for MockClientIntf
type MockClientIntfMockRecorder struct {
	mock *MockClientIntf
}

// NewMockClientIntf creates a new mock instance
func NewMockClientIntf(ctrl *gomock.Controller) *MockClientIntf {
	mock := &MockClientIntf{ctrl: ctrl}
	mock.recorder = &MockClientIntfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientIntf) EXPECT() *MockClientIntfMockRecorder {
	return m.recorder
}

// BeginPolling mocks base method
func (m *MockClientIntf) BeginPolling(arg0 *zmqutil.Poller, arg1 zmq4.State) *zmq4.Socket {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginPolling", arg0, arg1)
	ret0, _ := ret[0].(*zmq4.Socket)
	return ret0
}

// BeginPolling indicates an expected call of BeginPolling
func (mr *MockClientIntfMockRecorder) BeginPolling(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginPolling", reflect.TypeOf((*MockClientIntf)(nil).BeginPolling), arg0, arg1)
}

// Close mocks base method
func (m *MockClientIntf) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClientIntfMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientIntf)(nil).Close))
}

// Connect mocks base method
func (m *MockClientIntf) Connect(arg0 *util.Connection, arg1 []byte, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockClientIntfMockRecorder) Connect(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClientIntf)(nil).Connect), arg0, arg1, arg2)
}

// ConnectedTo mocks base method
func (m *MockClientIntf) ConnectedTo() *zmqutil.Connected {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedTo")
	ret0, _ := ret[0].(*zmqutil.Connected)
	return ret0
}

// ConnectedTo indicates an expected call of ConnectedTo
func (mr *MockClientIntfMockRecorder) ConnectedTo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedTo", reflect.TypeOf((*MockClientIntf)(nil).ConnectedTo))
}

// GetMonitorSocket mocks base method
func (m *MockClientIntf) GetMonitorSocket() *zmq4.Socket {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMonitorSocket")
	ret0, _ := ret[0].(*zmq4.Socket)
	return ret0
}

// GetMonitorSocket indicates an expected call of GetMonitorSocket
func (mr *MockClientIntfMockRecorder) GetMonitorSocket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMonitorSocket", reflect.TypeOf((*MockClientIntf)(nil).GetMonitorSocket))
}

// GetSocket mocks base method
func (m *MockClientIntf) GetSocket() *zmq4.Socket {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSocket")
	ret0, _ := ret[0].(*zmq4.Socket)
	return ret0
}

// GetSocket indicates an expected call of GetSocket
func (mr *MockClientIntfMockRecorder) GetSocket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSocket", reflect.TypeOf((*MockClientIntf)(nil).GetSocket))
}

// GoString mocks base method
func (m *MockClientIntf) GoString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GoString")
	ret0, _ := ret[0].(string)
	return ret0
}

// GoString indicates an expected call of GoString
func (mr *MockClientIntfMockRecorder) GoString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GoString", reflect.TypeOf((*MockClientIntf)(nil).GoString))
}

// IsConnected mocks base method
func (m *MockClientIntf) IsConnected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConnected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConnected indicates an expected call of IsConnected
func (mr *MockClientIntfMockRecorder) IsConnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnected", reflect.TypeOf((*MockClientIntf)(nil).IsConnected))
}

// IsConnectedTo mocks base method
func (m *MockClientIntf) IsConnectedTo(arg0 []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConnectedTo", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConnectedTo indicates an expected call of IsConnectedTo
func (mr *MockClientIntfMockRecorder) IsConnectedTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnectedTo", reflect.TypeOf((*MockClientIntf)(nil).IsConnectedTo), arg0)
}

// Receive mocks base method
func (m *MockClientIntf) Receive(arg0 zmq4.Flag) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", arg0)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive
func (mr *MockClientIntfMockRecorder) Receive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockClientIntf)(nil).Receive), arg0)
}

// Reconnect mocks base method
func (m *MockClientIntf) Reconnect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconnect indicates an expected call of Reconnect
func (mr *MockClientIntfMockRecorder) Reconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconnect", reflect.TypeOf((*MockClientIntf)(nil).Reconnect))
}

// ReconnectOpenedSocket mocks base method
func (m *MockClientIntf) ReconnectOpenedSocket() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconnectOpenedSocket")
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconnectOpenedSocket indicates an expected call of ReconnectOpenedSocket
func (mr *MockClientIntfMockRecorder) ReconnectOpenedSocket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconnectOpenedSocket", reflect.TypeOf((*MockClientIntf)(nil).ReconnectOpenedSocket))
}

// ReconnectReturningSocket mocks base method
func (m *MockClientIntf) ReconnectReturningSocket() (*zmq4.Socket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconnectReturningSocket")
	ret0, _ := ret[0].(*zmq4.Socket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconnectReturningSocket indicates an expected call of ReconnectReturningSocket
func (mr *MockClientIntfMockRecorder) ReconnectReturningSocket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconnectReturningSocket", reflect.TypeOf((*MockClientIntf)(nil).ReconnectReturningSocket))
}

// ResetServer mocks base method
func (m *MockClientIntf) ResetServer() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetServer")
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetServer indicates an expected call of ResetServer
func (mr *MockClientIntfMockRecorder) ResetServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetServer", reflect.TypeOf((*MockClientIntf)(nil).ResetServer))
}

// Send mocks base method
func (m *MockClientIntf) Send(arg0 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockClientIntfMockRecorder) Send(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockClientIntf)(nil).Send), arg0...)
}

// ServerPublicKey mocks base method
func (m *MockClientIntf) ServerPublicKey() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServerPublicKey")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// ServerPublicKey indicates an expected call of ServerPublicKey
func (mr *MockClientIntfMockRecorder) ServerPublicKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServerPublicKey", reflect.TypeOf((*MockClientIntf)(nil).ServerPublicKey))
}

// String mocks base method
func (m *MockClientIntf) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockClientIntfMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockClientIntf)(nil).String))
}
