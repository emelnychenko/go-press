// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_publisher_job.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostPublisherJob is a mock of PostPublisherJob interface
type MockPostPublisherJob struct {
	ctrl     *gomock.Controller
	recorder *MockPostPublisherJobMockRecorder
}

// MockPostPublisherJobMockRecorder is the mock recorder for MockPostPublisherJob
type MockPostPublisherJobMockRecorder struct {
	mock *MockPostPublisherJob
}

// NewMockPostPublisherJob creates a new mock instance
func NewMockPostPublisherJob(ctrl *gomock.Controller) *MockPostPublisherJob {
	mock := &MockPostPublisherJob{ctrl: ctrl}
	mock.recorder = &MockPostPublisherJobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostPublisherJob) EXPECT() *MockPostPublisherJobMockRecorder {
	return m.recorder
}

// PublishPost mocks base method
func (m *MockPostPublisherJob) PublishPost(postEntity *entities.PostEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishPost", postEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// PublishPost indicates an expected call of PublishPost
func (mr *MockPostPublisherJobMockRecorder) PublishPost(postEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishPost", reflect.TypeOf((*MockPostPublisherJob)(nil).PublishPost), postEntity)
}
