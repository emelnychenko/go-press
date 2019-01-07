// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/hasher.go

// Package mocks is a generated GoMock package.
package mocks

import (
	errors "github.com/emelnychenko/go-press/errors"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHasher is a mock of Hasher interface
type MockHasher struct {
	ctrl     *gomock.Controller
	recorder *MockHasherMockRecorder
}

// MockHasherMockRecorder is the mock recorder for MockHasher
type MockHasherMockRecorder struct {
	mock *MockHasher
}

// NewMockHasher creates a new mock instance
func NewMockHasher(ctrl *gomock.Controller) *MockHasher {
	mock := &MockHasher{ctrl: ctrl}
	mock.recorder = &MockHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHasher) EXPECT() *MockHasherMockRecorder {
	return m.recorder
}

// Make mocks base method
func (m *MockHasher) Make(password string) (string, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Make", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// Make indicates an expected call of Make
func (mr *MockHasherMockRecorder) Make(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Make", reflect.TypeOf((*MockHasher)(nil).Make), password)
}

// Check mocks base method
func (m *MockHasher) Check(hashedPassword, password string) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", hashedPassword, password)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockHasherMockRecorder) Check(hashedPassword, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockHasher)(nil).Check), hashedPassword, password)
}
