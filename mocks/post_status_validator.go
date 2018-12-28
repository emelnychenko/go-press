// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_status_validator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostStatusValidator is a mock of PostStatusValidator interface
type MockPostStatusValidator struct {
	ctrl     *gomock.Controller
	recorder *MockPostStatusValidatorMockRecorder
}

// MockPostStatusValidatorMockRecorder is the mock recorder for MockPostStatusValidator
type MockPostStatusValidatorMockRecorder struct {
	mock *MockPostStatusValidator
}

// NewMockPostStatusValidator creates a new mock instance
func NewMockPostStatusValidator(ctrl *gomock.Controller) *MockPostStatusValidator {
	mock := &MockPostStatusValidator{ctrl: ctrl}
	mock.recorder = &MockPostStatusValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostStatusValidator) EXPECT() *MockPostStatusValidatorMockRecorder {
	return m.recorder
}

// ValidatePostCreate mocks base method
func (m *MockPostStatusValidator) ValidatePostCreate(data *models.PostCreate) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePostCreate", data)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// ValidatePostCreate indicates an expected call of ValidatePostCreate
func (mr *MockPostStatusValidatorMockRecorder) ValidatePostCreate(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePostCreate", reflect.TypeOf((*MockPostStatusValidator)(nil).ValidatePostCreate), data)
}

// ValidatePostUpdate mocks base method
func (m *MockPostStatusValidator) ValidatePostUpdate(data *models.PostUpdate) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePostUpdate", data)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// ValidatePostUpdate indicates an expected call of ValidatePostUpdate
func (mr *MockPostStatusValidatorMockRecorder) ValidatePostUpdate(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePostUpdate", reflect.TypeOf((*MockPostStatusValidator)(nil).ValidatePostUpdate), data)
}
