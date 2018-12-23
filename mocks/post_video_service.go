// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_video_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	entities "github.com/emelnychenko/go-press/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostVideoService is a mock of PostVideoService interface
type MockPostVideoService struct {
	ctrl     *gomock.Controller
	recorder *MockPostVideoServiceMockRecorder
}

// MockPostVideoServiceMockRecorder is the mock recorder for MockPostVideoService
type MockPostVideoServiceMockRecorder struct {
	mock *MockPostVideoService
}

// NewMockPostVideoService creates a new mock instance
func NewMockPostVideoService(ctrl *gomock.Controller) *MockPostVideoService {
	mock := &MockPostVideoService{ctrl: ctrl}
	mock.recorder = &MockPostVideoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostVideoService) EXPECT() *MockPostVideoServiceMockRecorder {
	return m.recorder
}

// ChangePostVideo mocks base method
func (m *MockPostVideoService) ChangePostVideo(postEntity *entities.PostEntity, postVideo *entities.FileEntity) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePostVideo", postEntity, postVideo)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// ChangePostVideo indicates an expected call of ChangePostVideo
func (mr *MockPostVideoServiceMockRecorder) ChangePostVideo(postEntity, postVideo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePostVideo", reflect.TypeOf((*MockPostVideoService)(nil).ChangePostVideo), postEntity, postVideo)
}

// RemovePostVideo mocks base method
func (m *MockPostVideoService) RemovePostVideo(postEntity *entities.PostEntity) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePostVideo", postEntity)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// RemovePostVideo indicates an expected call of RemovePostVideo
func (mr *MockPostVideoServiceMockRecorder) RemovePostVideo(postEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePostVideo", reflect.TypeOf((*MockPostVideoService)(nil).RemovePostVideo), postEntity)
}