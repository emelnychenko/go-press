// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_author_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostAuthorApi is a mock of PostAuthorApi interface
type MockPostAuthorApi struct {
	ctrl     *gomock.Controller
	recorder *MockPostAuthorApiMockRecorder
}

// MockPostAuthorApiMockRecorder is the mock recorder for MockPostAuthorApi
type MockPostAuthorApiMockRecorder struct {
	mock *MockPostAuthorApi
}

// NewMockPostAuthorApi creates a new mock instance
func NewMockPostAuthorApi(ctrl *gomock.Controller) *MockPostAuthorApi {
	mock := &MockPostAuthorApi{ctrl: ctrl}
	mock.recorder = &MockPostAuthorApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostAuthorApi) EXPECT() *MockPostAuthorApiMockRecorder {
	return m.recorder
}

// ChangePostAuthor mocks base method
func (m *MockPostAuthorApi) ChangePostAuthor(postId *models.PostId, postAuthorId *models.UserId) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePostAuthor", postId, postAuthorId)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// ChangePostAuthor indicates an expected call of ChangePostAuthor
func (mr *MockPostAuthorApiMockRecorder) ChangePostAuthor(postId, postAuthorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePostAuthor", reflect.TypeOf((*MockPostAuthorApi)(nil).ChangePostAuthor), postId, postAuthorId)
}
