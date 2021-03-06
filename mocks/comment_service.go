// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/comment_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCommentService is a mock of CommentService interface
type MockCommentService struct {
	ctrl     *gomock.Controller
	recorder *MockCommentServiceMockRecorder
}

// MockCommentServiceMockRecorder is the mock recorder for MockCommentService
type MockCommentServiceMockRecorder struct {
	mock *MockCommentService
}

// NewMockCommentService creates a new mock instance
func NewMockCommentService(ctrl *gomock.Controller) *MockCommentService {
	mock := &MockCommentService{ctrl: ctrl}
	mock.recorder = &MockCommentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommentService) EXPECT() *MockCommentServiceMockRecorder {
	return m.recorder
}

// ListComments mocks base method
func (m *MockCommentService) ListComments(commentPaginationQuery *models.CommentPaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComments", commentPaginationQuery)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListComments indicates an expected call of ListComments
func (mr *MockCommentServiceMockRecorder) ListComments(commentPaginationQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComments", reflect.TypeOf((*MockCommentService)(nil).ListComments), commentPaginationQuery)
}

// GetComment mocks base method
func (m *MockCommentService) GetComment(commentId *models.CommentId) (*entities.CommentEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", commentId)
	ret0, _ := ret[0].(*entities.CommentEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment
func (mr *MockCommentServiceMockRecorder) GetComment(commentId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockCommentService)(nil).GetComment), commentId)
}

// CreateComment mocks base method
func (m *MockCommentService) CreateComment(data *models.CommentCreate) (*entities.CommentEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", data)
	ret0, _ := ret[0].(*entities.CommentEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment
func (mr *MockCommentServiceMockRecorder) CreateComment(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockCommentService)(nil).CreateComment), data)
}

// UpdateComment mocks base method
func (m *MockCommentService) UpdateComment(commentEntity *entities.CommentEntity, data *models.CommentUpdate) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", commentEntity, data)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// UpdateComment indicates an expected call of UpdateComment
func (mr *MockCommentServiceMockRecorder) UpdateComment(commentEntity, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockCommentService)(nil).UpdateComment), commentEntity, data)
}

// DeleteComment mocks base method
func (m *MockCommentService) DeleteComment(commentEntity *entities.CommentEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", commentEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment
func (mr *MockCommentServiceMockRecorder) DeleteComment(commentEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockCommentService)(nil).DeleteComment), commentEntity)
}
