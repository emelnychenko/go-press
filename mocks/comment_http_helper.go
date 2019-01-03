// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/comment_http_helper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	contracts "github.com/emelnychenko/go-press/contracts"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCommentHttpHelper is a mock of CommentHttpHelper interface
type MockCommentHttpHelper struct {
	ctrl     *gomock.Controller
	recorder *MockCommentHttpHelperMockRecorder
}

// MockCommentHttpHelperMockRecorder is the mock recorder for MockCommentHttpHelper
type MockCommentHttpHelperMockRecorder struct {
	mock *MockCommentHttpHelper
}

// NewMockCommentHttpHelper creates a new mock instance
func NewMockCommentHttpHelper(ctrl *gomock.Controller) *MockCommentHttpHelper {
	mock := &MockCommentHttpHelper{ctrl: ctrl}
	mock.recorder = &MockCommentHttpHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommentHttpHelper) EXPECT() *MockCommentHttpHelperMockRecorder {
	return m.recorder
}

// ParseCommentId mocks base method
func (m *MockCommentHttpHelper) ParseCommentId(httpContext contracts.HttpContext) (*models.CommentId, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseCommentId", httpContext)
	ret0, _ := ret[0].(*models.CommentId)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// ParseCommentId indicates an expected call of ParseCommentId
func (mr *MockCommentHttpHelperMockRecorder) ParseCommentId(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseCommentId", reflect.TypeOf((*MockCommentHttpHelper)(nil).ParseCommentId), httpContext)
}
