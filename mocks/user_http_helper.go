// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/user_http_helper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	contracts "github.com/emelnychenko/go-press/contracts"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserHttpHelper is a mock of UserHttpHelper interface
type MockUserHttpHelper struct {
	ctrl     *gomock.Controller
	recorder *MockUserHttpHelperMockRecorder
}

// MockUserHttpHelperMockRecorder is the mock recorder for MockUserHttpHelper
type MockUserHttpHelperMockRecorder struct {
	mock *MockUserHttpHelper
}

// NewMockUserHttpHelper creates a new mock instance
func NewMockUserHttpHelper(ctrl *gomock.Controller) *MockUserHttpHelper {
	mock := &MockUserHttpHelper{ctrl: ctrl}
	mock.recorder = &MockUserHttpHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserHttpHelper) EXPECT() *MockUserHttpHelperMockRecorder {
	return m.recorder
}

// ParseUserId mocks base method
func (m *MockUserHttpHelper) ParseUserId(httpContext contracts.HttpContext) (*models.UserId, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseUserId", httpContext)
	ret0, _ := ret[0].(*models.UserId)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ParseUserId indicates an expected call of ParseUserId
func (mr *MockUserHttpHelperMockRecorder) ParseUserId(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseUserId", reflect.TypeOf((*MockUserHttpHelper)(nil).ParseUserId), httpContext)
}
