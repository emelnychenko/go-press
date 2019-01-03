// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/poll_http_helper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	contracts "github.com/emelnychenko/go-press/contracts"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPollHttpHelper is a mock of PollHttpHelper interface
type MockPollHttpHelper struct {
	ctrl     *gomock.Controller
	recorder *MockPollHttpHelperMockRecorder
}

// MockPollHttpHelperMockRecorder is the mock recorder for MockPollHttpHelper
type MockPollHttpHelperMockRecorder struct {
	mock *MockPollHttpHelper
}

// NewMockPollHttpHelper creates a new mock instance
func NewMockPollHttpHelper(ctrl *gomock.Controller) *MockPollHttpHelper {
	mock := &MockPollHttpHelper{ctrl: ctrl}
	mock.recorder = &MockPollHttpHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPollHttpHelper) EXPECT() *MockPollHttpHelperMockRecorder {
	return m.recorder
}

// ParsePollId mocks base method
func (m *MockPollHttpHelper) ParsePollId(httpContext contracts.HttpContext) (*models.PollId, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParsePollId", httpContext)
	ret0, _ := ret[0].(*models.PollId)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// ParsePollId indicates an expected call of ParsePollId
func (mr *MockPollHttpHelperMockRecorder) ParsePollId(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParsePollId", reflect.TypeOf((*MockPollHttpHelper)(nil).ParsePollId), httpContext)
}
