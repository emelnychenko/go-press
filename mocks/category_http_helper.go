// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/category_http_helper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	contracts "github.com/emelnychenko/go-press/contracts"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCategoryHttpHelper is a mock of CategoryHttpHelper interface
type MockCategoryHttpHelper struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryHttpHelperMockRecorder
}

// MockCategoryHttpHelperMockRecorder is the mock recorder for MockCategoryHttpHelper
type MockCategoryHttpHelperMockRecorder struct {
	mock *MockCategoryHttpHelper
}

// NewMockCategoryHttpHelper creates a new mock instance
func NewMockCategoryHttpHelper(ctrl *gomock.Controller) *MockCategoryHttpHelper {
	mock := &MockCategoryHttpHelper{ctrl: ctrl}
	mock.recorder = &MockCategoryHttpHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCategoryHttpHelper) EXPECT() *MockCategoryHttpHelperMockRecorder {
	return m.recorder
}

// ParseCategoryId mocks base method
func (m *MockCategoryHttpHelper) ParseCategoryId(arg0 contracts.HttpContext) (*models.CategoryId, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseCategoryId", arg0)
	ret0, _ := ret[0].(*models.CategoryId)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ParseCategoryId indicates an expected call of ParseCategoryId
func (mr *MockCategoryHttpHelperMockRecorder) ParseCategoryId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseCategoryId", reflect.TypeOf((*MockCategoryHttpHelper)(nil).ParseCategoryId), arg0)
}

// ParseParentCategoryId mocks base method
func (m *MockCategoryHttpHelper) ParseParentCategoryId(arg0 contracts.HttpContext) (*models.CategoryId, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseParentCategoryId", arg0)
	ret0, _ := ret[0].(*models.CategoryId)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ParseParentCategoryId indicates an expected call of ParseParentCategoryId
func (mr *MockCategoryHttpHelperMockRecorder) ParseParentCategoryId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseParentCategoryId", reflect.TypeOf((*MockCategoryHttpHelper)(nil).ParseParentCategoryId), arg0)
}
