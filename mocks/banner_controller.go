// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/banner_controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	contracts "github.com/emelnychenko/go-press/contracts"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBannerController is a mock of BannerController interface
type MockBannerController struct {
	ctrl     *gomock.Controller
	recorder *MockBannerControllerMockRecorder
}

// MockBannerControllerMockRecorder is the mock recorder for MockBannerController
type MockBannerControllerMockRecorder struct {
	mock *MockBannerController
}

// NewMockBannerController creates a new mock instance
func NewMockBannerController(ctrl *gomock.Controller) *MockBannerController {
	mock := &MockBannerController{ctrl: ctrl}
	mock.recorder = &MockBannerControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBannerController) EXPECT() *MockBannerControllerMockRecorder {
	return m.recorder
}

// ListBanners mocks base method
func (m *MockBannerController) ListBanners(httpContext contracts.HttpContext) (interface{}, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBanners", httpContext)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// ListBanners indicates an expected call of ListBanners
func (mr *MockBannerControllerMockRecorder) ListBanners(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBanners", reflect.TypeOf((*MockBannerController)(nil).ListBanners), httpContext)
}

// GetBanner mocks base method
func (m *MockBannerController) GetBanner(httpContext contracts.HttpContext) (interface{}, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBanner", httpContext)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// GetBanner indicates an expected call of GetBanner
func (mr *MockBannerControllerMockRecorder) GetBanner(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBanner", reflect.TypeOf((*MockBannerController)(nil).GetBanner), httpContext)
}

// CreateBanner mocks base method
func (m *MockBannerController) CreateBanner(httpContext contracts.HttpContext) (interface{}, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBanner", httpContext)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// CreateBanner indicates an expected call of CreateBanner
func (mr *MockBannerControllerMockRecorder) CreateBanner(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBanner", reflect.TypeOf((*MockBannerController)(nil).CreateBanner), httpContext)
}

// UpdateBanner mocks base method
func (m *MockBannerController) UpdateBanner(httpContext contracts.HttpContext) (interface{}, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBanner", httpContext)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// UpdateBanner indicates an expected call of UpdateBanner
func (mr *MockBannerControllerMockRecorder) UpdateBanner(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBanner", reflect.TypeOf((*MockBannerController)(nil).UpdateBanner), httpContext)
}

// DeleteBanner mocks base method
func (m *MockBannerController) DeleteBanner(httpContext contracts.HttpContext) (interface{}, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBanner", httpContext)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// DeleteBanner indicates an expected call of DeleteBanner
func (mr *MockBannerControllerMockRecorder) DeleteBanner(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBanner", reflect.TypeOf((*MockBannerController)(nil).DeleteBanner), httpContext)
}
