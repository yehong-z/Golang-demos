package gomock_demo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMyInterface is a mock of MyInterface interface.
type MockMyInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMyInterfaceMockRecorder
}

// MockMyInterfaceMockRecorder is the mock recorder for MockMyInterface.
type MockMyInterfaceMockRecorder struct {
	mock *MockMyInterface
}

// NewMockMyInterface creates a new mock instance.
func NewMockMyInterface(ctrl *gomock.Controller) *MockMyInterface {
	mock := &MockMyInterface{ctrl: ctrl}
	mock.recorder = &MockMyInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMyInterface) EXPECT() *MockMyInterfaceMockRecorder {
	return m.recorder
}

// DoSomething mocks base method.
func (m *MockMyInterface) DoSomething(arg1 string, arg2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoSomething", arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoSomething indicates an expected call of DoSomething.
func (mr *MockMyInterfaceMockRecorder) DoSomething(arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoSomething", reflect.TypeOf((*MockMyInterface)(nil).DoSomething), arg1, arg2)
}
