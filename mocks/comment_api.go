// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/comment_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCommentApi is a mock of CommentApi interface
type MockCommentApi struct {
	ctrl     *gomock.Controller
	recorder *MockCommentApiMockRecorder
}

// MockCommentApiMockRecorder is the mock recorder for MockCommentApi
type MockCommentApiMockRecorder struct {
	mock *MockCommentApi
}

// NewMockCommentApi creates a new mock instance
func NewMockCommentApi(ctrl *gomock.Controller) *MockCommentApi {
	mock := &MockCommentApi{ctrl: ctrl}
	mock.recorder = &MockCommentApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommentApi) EXPECT() *MockCommentApiMockRecorder {
	return m.recorder
}

// ListComments mocks base method
func (m *MockCommentApi) ListComments(commentPaginationQuery *models.CommentPaginationQuery) (*models.PaginationResult, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComments", commentPaginationQuery)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// ListComments indicates an expected call of ListComments
func (mr *MockCommentApiMockRecorder) ListComments(commentPaginationQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComments", reflect.TypeOf((*MockCommentApi)(nil).ListComments), commentPaginationQuery)
}

// GetComment mocks base method
func (m *MockCommentApi) GetComment(commentId *models.CommentId) (*models.Comment, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", commentId)
	ret0, _ := ret[0].(*models.Comment)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment
func (mr *MockCommentApiMockRecorder) GetComment(commentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockCommentApi)(nil).GetComment), commentId)
}

// CreateComment mocks base method
func (m *MockCommentApi) CreateComment(data *models.CommentCreate) (*models.Comment, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", data)
	ret0, _ := ret[0].(*models.Comment)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment
func (mr *MockCommentApiMockRecorder) CreateComment(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockCommentApi)(nil).CreateComment), data)
}

// UpdateComment mocks base method
func (m *MockCommentApi) UpdateComment(commentId *models.CommentId, data *models.CommentUpdate) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", commentId, data)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// UpdateComment indicates an expected call of UpdateComment
func (mr *MockCommentApiMockRecorder) UpdateComment(commentId, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockCommentApi)(nil).UpdateComment), commentId, data)
}

// DeleteComment mocks base method
func (m *MockCommentApi) DeleteComment(commentId *models.CommentId) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", commentId)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment
func (mr *MockCommentApiMockRecorder) DeleteComment(commentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockCommentApi)(nil).DeleteComment), commentId)
}