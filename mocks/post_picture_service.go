// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_picture_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostPictureService is a mock of PostPictureService interface
type MockPostPictureService struct {
	ctrl     *gomock.Controller
	recorder *MockPostPictureServiceMockRecorder
}

// MockPostPictureServiceMockRecorder is the mock recorder for MockPostPictureService
type MockPostPictureServiceMockRecorder struct {
	mock *MockPostPictureService
}

// NewMockPostPictureService creates a new mock instance
func NewMockPostPictureService(ctrl *gomock.Controller) *MockPostPictureService {
	mock := &MockPostPictureService{ctrl: ctrl}
	mock.recorder = &MockPostPictureServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostPictureService) EXPECT() *MockPostPictureServiceMockRecorder {
	return m.recorder
}

// ChangePostPicture mocks base method
func (m *MockPostPictureService) ChangePostPicture(postEntity *entities.PostEntity, postPictureEntity *entities.FileEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePostPicture", postEntity, postPictureEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// ChangePostPicture indicates an expected call of ChangePostPicture
func (mr *MockPostPictureServiceMockRecorder) ChangePostPicture(postEntity, postPictureEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePostPicture", reflect.TypeOf((*MockPostPictureService)(nil).ChangePostPicture), postEntity, postPictureEntity)
}

// RemovePostPicture mocks base method
func (m *MockPostPictureService) RemovePostPicture(postEntity *entities.PostEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePostPicture", postEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// RemovePostPicture indicates an expected call of RemovePostPicture
func (mr *MockPostPictureServiceMockRecorder) RemovePostPicture(postEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePostPicture", reflect.TypeOf((*MockPostPictureService)(nil).RemovePostPicture), postEntity)
}
