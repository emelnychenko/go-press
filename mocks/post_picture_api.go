// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_picture_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostPictureApi is a mock of PostPictureApi interface
type MockPostPictureApi struct {
	ctrl     *gomock.Controller
	recorder *MockPostPictureApiMockRecorder
}

// MockPostPictureApiMockRecorder is the mock recorder for MockPostPictureApi
type MockPostPictureApiMockRecorder struct {
	mock *MockPostPictureApi
}

// NewMockPostPictureApi creates a new mock instance
func NewMockPostPictureApi(ctrl *gomock.Controller) *MockPostPictureApi {
	mock := &MockPostPictureApi{ctrl: ctrl}
	mock.recorder = &MockPostPictureApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostPictureApi) EXPECT() *MockPostPictureApiMockRecorder {
	return m.recorder
}

// ChangePostPicture mocks base method
func (m *MockPostPictureApi) ChangePostPicture(postId *models.PostId, postPictureId *models.FileId) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePostPicture", postId, postPictureId)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// ChangePostPicture indicates an expected call of ChangePostPicture
func (mr *MockPostPictureApiMockRecorder) ChangePostPicture(postId, postPictureId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePostPicture", reflect.TypeOf((*MockPostPictureApi)(nil).ChangePostPicture), postId, postPictureId)
}

// RemovePostPicture mocks base method
func (m *MockPostPictureApi) RemovePostPicture(postId *models.PostId) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePostPicture", postId)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// RemovePostPicture indicates an expected call of RemovePostPicture
func (mr *MockPostPictureApiMockRecorder) RemovePostPicture(postId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePostPicture", reflect.TypeOf((*MockPostPictureApi)(nil).RemovePostPicture), postId)
}
