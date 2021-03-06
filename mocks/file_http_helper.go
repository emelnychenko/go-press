// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/file_http_helper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	contracts "github.com/emelnychenko/go-press/contracts"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	multipart "mime/multipart"
	reflect "reflect"
)

// MockFileHttpHelper is a mock of FileHttpHelper interface
type MockFileHttpHelper struct {
	ctrl     *gomock.Controller
	recorder *MockFileHttpHelperMockRecorder
}

// MockFileHttpHelperMockRecorder is the mock recorder for MockFileHttpHelper
type MockFileHttpHelperMockRecorder struct {
	mock *MockFileHttpHelper
}

// NewMockFileHttpHelper creates a new mock instance
func NewMockFileHttpHelper(ctrl *gomock.Controller) *MockFileHttpHelper {
	mock := &MockFileHttpHelper{ctrl: ctrl}
	mock.recorder = &MockFileHttpHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileHttpHelper) EXPECT() *MockFileHttpHelperMockRecorder {
	return m.recorder
}

// ParseFileId mocks base method
func (m *MockFileHttpHelper) ParseFileId(httpContext contracts.HttpContext) (*models.FileId, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseFileId", httpContext)
	ret0, _ := ret[0].(*models.FileId)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ParseFileId indicates an expected call of ParseFileId
func (mr *MockFileHttpHelperMockRecorder) ParseFileId(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseFileId", reflect.TypeOf((*MockFileHttpHelper)(nil).ParseFileId), httpContext)
}

// GetFileHeader mocks base method
func (m *MockFileHttpHelper) GetFileHeader(httpContext contracts.HttpContext) (*multipart.FileHeader, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileHeader", httpContext)
	ret0, _ := ret[0].(*multipart.FileHeader)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetFileHeader indicates an expected call of GetFileHeader
func (mr *MockFileHttpHelperMockRecorder) GetFileHeader(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileHeader", reflect.TypeOf((*MockFileHttpHelper)(nil).GetFileHeader), httpContext)
}

// OpenFormFile mocks base method
func (m *MockFileHttpHelper) OpenFormFile(httpContext *multipart.FileHeader) (multipart.File, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFormFile", httpContext)
	ret0, _ := ret[0].(multipart.File)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// OpenFormFile indicates an expected call of OpenFormFile
func (mr *MockFileHttpHelperMockRecorder) OpenFormFile(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFormFile", reflect.TypeOf((*MockFileHttpHelper)(nil).OpenFormFile), httpContext)
}

// PrepareFileDestination mocks base method
func (m *MockFileHttpHelper) PrepareFileDestination(httpContext contracts.HttpContext) contracts.PrepareFileDestination {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareFileDestination", httpContext)
	ret0, _ := ret[0].(contracts.PrepareFileDestination)
	return ret0
}

// PrepareFileDestination indicates an expected call of PrepareFileDestination
func (mr *MockFileHttpHelperMockRecorder) PrepareFileDestination(httpContext interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareFileDestination", reflect.TypeOf((*MockFileHttpHelper)(nil).PrepareFileDestination), httpContext)
}
