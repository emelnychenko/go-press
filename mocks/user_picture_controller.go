// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/user_picture_controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	contracts "github.com/emelnychenko/go-press/contracts"
	errors "github.com/emelnychenko/go-press/errors"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserPictureController is a mock of UserPictureController interface
type MockUserPictureController struct {
	ctrl     *gomock.Controller
	recorder *MockUserPictureControllerMockRecorder
}

// MockUserPictureControllerMockRecorder is the mock recorder for MockUserPictureController
type MockUserPictureControllerMockRecorder struct {
	mock *MockUserPictureController
}

// NewMockUserPictureController creates a new mock instance
func NewMockUserPictureController(ctrl *gomock.Controller) *MockUserPictureController {
	mock := &MockUserPictureController{ctrl: ctrl}
	mock.recorder = &MockUserPictureControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserPictureController) EXPECT() *MockUserPictureControllerMockRecorder {
	return m.recorder
}

// ChangeUserPicture mocks base method
func (m *MockUserPictureController) ChangeUserPicture(httpContext contracts.HttpContext) (interface{}, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserPicture", httpContext)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ChangeUserPicture indicates an expected call of ChangeUserPicture
func (mr *MockUserPictureControllerMockRecorder) ChangeUserPicture(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserPicture", reflect.TypeOf((*MockUserPictureController)(nil).ChangeUserPicture), httpContext)
}

// RemoveUserPicture mocks base method
func (m *MockUserPictureController) RemoveUserPicture(httpContext contracts.HttpContext) (interface{}, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUserPicture", httpContext)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// RemoveUserPicture indicates an expected call of RemoveUserPicture
func (mr *MockUserPictureControllerMockRecorder) RemoveUserPicture(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUserPicture", reflect.TypeOf((*MockUserPictureController)(nil).RemoveUserPicture), httpContext)
}
