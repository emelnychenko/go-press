// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/file_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	contracts "github.com/emelnychenko/go-press/contracts"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockFileApi is a mock of FileApi interface
type MockFileApi struct {
	ctrl     *gomock.Controller
	recorder *MockFileApiMockRecorder
}

// MockFileApiMockRecorder is the mock recorder for MockFileApi
type MockFileApiMockRecorder struct {
	mock *MockFileApi
}

// NewMockFileApi creates a new mock instance
func NewMockFileApi(ctrl *gomock.Controller) *MockFileApi {
	mock := &MockFileApi{ctrl: ctrl}
	mock.recorder = &MockFileApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileApi) EXPECT() *MockFileApiMockRecorder {
	return m.recorder
}

// ListFiles mocks base method
func (m *MockFileApi) ListFiles(filePaginationQuery *models.FilePaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", filePaginationQuery)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles
func (mr *MockFileApiMockRecorder) ListFiles(filePaginationQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockFileApi)(nil).ListFiles), filePaginationQuery)
}

// GetFile mocks base method
func (m *MockFileApi) GetFile(fileId *models.FileId) (*models.File, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", fileId)
	ret0, _ := ret[0].(*models.File)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile
func (mr *MockFileApiMockRecorder) GetFile(fileId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockFileApi)(nil).GetFile), fileId)
}

// UploadFile mocks base method
func (m *MockFileApi) UploadFile(fileSource io.Reader, data *models.FileUpload) (*models.File, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", fileSource, data)
	ret0, _ := ret[0].(*models.File)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile
func (mr *MockFileApiMockRecorder) UploadFile(fileSource, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockFileApi)(nil).UploadFile), fileSource, data)
}

// DownloadFile mocks base method
func (m *MockFileApi) DownloadFile(fileId *models.FileId, prepareFileDestination contracts.PrepareFileDestination) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", fileId, prepareFileDestination)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// DownloadFile indicates an expected call of DownloadFile
func (mr *MockFileApiMockRecorder) DownloadFile(fileId, prepareFileDestination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockFileApi)(nil).DownloadFile), fileId, prepareFileDestination)
}

// UpdateFile mocks base method
func (m *MockFileApi) UpdateFile(fileId *models.FileId, data *models.FileUpdate) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFile", fileId, data)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// UpdateFile indicates an expected call of UpdateFile
func (mr *MockFileApiMockRecorder) UpdateFile(fileId, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFile", reflect.TypeOf((*MockFileApi)(nil).UpdateFile), fileId, data)
}

// DeleteFile mocks base method
func (m *MockFileApi) DeleteFile(fileId *models.FileId) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", fileId)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile
func (mr *MockFileApiMockRecorder) DeleteFile(fileId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockFileApi)(nil).DeleteFile), fileId)
}
