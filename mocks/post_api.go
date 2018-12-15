// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostApi is a mock of PostApi interface
type MockPostApi struct {
	ctrl     *gomock.Controller
	recorder *MockPostApiMockRecorder
}

// MockPostApiMockRecorder is the mock recorder for MockPostApi
type MockPostApiMockRecorder struct {
	mock *MockPostApi
}

// NewMockPostApi creates a new mock instance
func NewMockPostApi(ctrl *gomock.Controller) *MockPostApi {
	mock := &MockPostApi{ctrl: ctrl}
	mock.recorder = &MockPostApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostApi) EXPECT() *MockPostApiMockRecorder {
	return m.recorder
}

// ListPosts mocks base method
func (m *MockPostApi) ListPosts() ([]*models.Post, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPosts")
	ret0, _ := ret[0].([]*models.Post)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// ListPosts indicates an expected call of ListPosts
func (mr *MockPostApiMockRecorder) ListPosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPosts", reflect.TypeOf((*MockPostApi)(nil).ListPosts))
}

// GetPost mocks base method
func (m *MockPostApi) GetPost(postId *models.PostId) (*models.Post, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", postId)
	ret0, _ := ret[0].(*models.Post)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost
func (mr *MockPostApiMockRecorder) GetPost(postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockPostApi)(nil).GetPost), postId)
}

// CreatePost mocks base method
func (m *MockPostApi) CreatePost(postAuthor models.Subject, data *models.PostCreate) (*models.Post, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", postAuthor, data)
	ret0, _ := ret[0].(*models.Post)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost
func (mr *MockPostApiMockRecorder) CreatePost(postAuthor, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostApi)(nil).CreatePost), postAuthor, data)
}

// UpdatePost mocks base method
func (m *MockPostApi) UpdatePost(postId *models.PostId, data *models.PostUpdate) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", postId, data)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost
func (mr *MockPostApiMockRecorder) UpdatePost(postId, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostApi)(nil).UpdatePost), postId, data)
}

// ChangePostAuthor mocks base method
func (m *MockPostApi) ChangePostAuthor(postId *models.PostId, postAuthor models.Subject) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePostAuthor", postId, postAuthor)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// ChangePostAuthor indicates an expected call of ChangePostAuthor
func (mr *MockPostApiMockRecorder) ChangePostAuthor(postId, postAuthor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePostAuthor", reflect.TypeOf((*MockPostApi)(nil).ChangePostAuthor), postId, postAuthor)
}

// DeletePost mocks base method
func (m *MockPostApi) DeletePost(postId *models.PostId) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", postId)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost
func (mr *MockPostApiMockRecorder) DeletePost(postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostApi)(nil).DeletePost), postId)
}
