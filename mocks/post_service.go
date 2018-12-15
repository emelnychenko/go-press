// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	entities "github.com/emelnychenko/go-press/entities"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostService is a mock of PostService interface
type MockPostService struct {
	ctrl     *gomock.Controller
	recorder *MockPostServiceMockRecorder
}

// MockPostServiceMockRecorder is the mock recorder for MockPostService
type MockPostServiceMockRecorder struct {
	mock *MockPostService
}

// NewMockPostService creates a new mock instance
func NewMockPostService(ctrl *gomock.Controller) *MockPostService {
	mock := &MockPostService{ctrl: ctrl}
	mock.recorder = &MockPostServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostService) EXPECT() *MockPostServiceMockRecorder {
	return m.recorder
}

// ListPosts mocks base method
func (m *MockPostService) ListPosts() ([]*entities.PostEntity, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPosts")
	ret0, _ := ret[0].([]*entities.PostEntity)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// ListPosts indicates an expected call of ListPosts
func (mr *MockPostServiceMockRecorder) ListPosts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPosts", reflect.TypeOf((*MockPostService)(nil).ListPosts))
}

// GetPost mocks base method
func (m *MockPostService) GetPost(postId *models.PostId) (*entities.PostEntity, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", postId)
	ret0, _ := ret[0].(*entities.PostEntity)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost
func (mr *MockPostServiceMockRecorder) GetPost(postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockPostService)(nil).GetPost), postId)
}

// CreatePost mocks base method
func (m *MockPostService) CreatePost(data *models.PostCreate) (*entities.PostEntity, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", data)
	ret0, _ := ret[0].(*entities.PostEntity)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost
func (mr *MockPostServiceMockRecorder) CreatePost(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockPostService)(nil).CreatePost), data)
}

// UpdatePost mocks base method
func (m *MockPostService) UpdatePost(postEntity *entities.PostEntity, data *models.PostUpdate) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", postEntity, data)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost
func (mr *MockPostServiceMockRecorder) UpdatePost(postEntity, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostService)(nil).UpdatePost), postEntity, data)
}

// DeletePost mocks base method
func (m *MockPostService) DeletePost(postEntity *entities.PostEntity) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", postEntity)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost
func (mr *MockPostServiceMockRecorder) DeletePost(postEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostService)(nil).DeletePost), postEntity)
}
