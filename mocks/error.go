// Code generated by MockGen. DO NOT EDIT.
// Source: errors/error.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockError is a mock of Error interface
type MockError struct {
	ctrl     *gomock.Controller
	recorder *MockErrorMockRecorder
}

// MockErrorMockRecorder is the mock recorder for MockError
type MockErrorMockRecorder struct {
	mock *MockError
}

// NewMockError creates a new mock instance
func NewMockError(ctrl *gomock.Controller) *MockError {
	mock := &MockError{ctrl: ctrl}
	mock.recorder = &MockErrorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockError) EXPECT() *MockErrorMockRecorder {
	return m.recorder
}

// Error mocks base method
func (m *MockError) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockErrorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockError)(nil).Error))
}

// Code mocks base method
func (m *MockError) Code() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Code")
	ret0, _ := ret[0].(int)
	return ret0
}

// Code indicates an expected call of Code
func (mr *MockErrorMockRecorder) Code() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Code", reflect.TypeOf((*MockError)(nil).Code))
}
