// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_author_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostAuthorService is a mock of PostAuthorService interface
type MockPostAuthorService struct {
	ctrl     *gomock.Controller
	recorder *MockPostAuthorServiceMockRecorder
}

// MockPostAuthorServiceMockRecorder is the mock recorder for MockPostAuthorService
type MockPostAuthorServiceMockRecorder struct {
	mock *MockPostAuthorService
}

// NewMockPostAuthorService creates a new mock instance
func NewMockPostAuthorService(ctrl *gomock.Controller) *MockPostAuthorService {
	mock := &MockPostAuthorService{ctrl: ctrl}
	mock.recorder = &MockPostAuthorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostAuthorService) EXPECT() *MockPostAuthorServiceMockRecorder {
	return m.recorder
}

// ChangePostAuthor mocks base method
func (m *MockPostAuthorService) ChangePostAuthor(postEntity *entities.PostEntity, postAuthorEntity *entities.UserEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePostAuthor", postEntity, postAuthorEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// ChangePostAuthor indicates an expected call of ChangePostAuthor
func (mr *MockPostAuthorServiceMockRecorder) ChangePostAuthor(postEntity, postAuthorEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePostAuthor", reflect.TypeOf((*MockPostAuthorService)(nil).ChangePostAuthor), postEntity, postAuthorEntity)
}
