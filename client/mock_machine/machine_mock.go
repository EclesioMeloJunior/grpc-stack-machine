// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/EclesioMeloJunior/grpc-stack-machine/machine (interfaces: MachineClient)

// Package mock_machine is a generated GoMock package.
package mock_machine

import (
	context "context"
	machine "github.com/EclesioMeloJunior/grpc-stack-machine/machine"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockMachineClient is a mock of MachineClient interface
type MockMachineClient struct {
	ctrl     *gomock.Controller
	recorder *MockMachineClientMockRecorder
}

// MockMachineClientMockRecorder is the mock recorder for MockMachineClient
type MockMachineClientMockRecorder struct {
	mock *MockMachineClient
}

// NewMockMachineClient creates a new mock instance
func NewMockMachineClient(ctrl *gomock.Controller) *MockMachineClient {
	mock := &MockMachineClient{ctrl: ctrl}
	mock.recorder = &MockMachineClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachineClient) EXPECT() *MockMachineClientMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockMachineClient) Execute(arg0 context.Context, arg1 *machine.InstructionSet, arg2 ...grpc.CallOption) (*machine.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Execute", varargs...)
	ret0, _ := ret[0].(*machine.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockMachineClientMockRecorder) Execute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockMachineClient)(nil).Execute), varargs...)
}
