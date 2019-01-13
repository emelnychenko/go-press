// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_tag_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostTagApi is a mock of PostTagApi interface
type MockPostTagApi struct {
	ctrl     *gomock.Controller
	recorder *MockPostTagApiMockRecorder
}

// MockPostTagApiMockRecorder is the mock recorder for MockPostTagApi
type MockPostTagApiMockRecorder struct {
	mock *MockPostTagApi
}

// NewMockPostTagApi creates a new mock instance
func NewMockPostTagApi(ctrl *gomock.Controller) *MockPostTagApi {
	mock := &MockPostTagApi{ctrl: ctrl}
	mock.recorder = &MockPostTagApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostTagApi) EXPECT() *MockPostTagApiMockRecorder {
	return m.recorder
}

// ListPostTags mocks base method
func (m *MockPostTagApi) ListPostTags(arg0 *models.PostId, arg1 *models.TagPaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPostTags", arg0, arg1)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListPostTags indicates an expected call of ListPostTags
func (mr *MockPostTagApiMockRecorder) ListPostTags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPostTags", reflect.TypeOf((*MockPostTagApi)(nil).ListPostTags), arg0, arg1)
}

// AddPostTag mocks base method
func (m *MockPostTagApi) AddPostTag(arg0 *models.PostId, arg1 *models.TagId) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPostTag", arg0, arg1)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// AddPostTag indicates an expected call of AddPostTag
func (mr *MockPostTagApiMockRecorder) AddPostTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPostTag", reflect.TypeOf((*MockPostTagApi)(nil).AddPostTag), arg0, arg1)
}

// RemovePostTag mocks base method
func (m *MockPostTagApi) RemovePostTag(arg0 *models.PostId, arg1 *models.TagId) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePostTag", arg0, arg1)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// RemovePostTag indicates an expected call of RemovePostTag
func (mr *MockPostTagApiMockRecorder) RemovePostTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePostTag", reflect.TypeOf((*MockPostTagApi)(nil).RemovePostTag), arg0, arg1)
}
