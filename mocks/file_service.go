// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/file_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockFileService is a mock of FileService interface
type MockFileService struct {
	ctrl     *gomock.Controller
	recorder *MockFileServiceMockRecorder
}

// MockFileServiceMockRecorder is the mock recorder for MockFileService
type MockFileServiceMockRecorder struct {
	mock *MockFileService
}

// NewMockFileService creates a new mock instance
func NewMockFileService(ctrl *gomock.Controller) *MockFileService {
	mock := &MockFileService{ctrl: ctrl}
	mock.recorder = &MockFileServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileService) EXPECT() *MockFileServiceMockRecorder {
	return m.recorder
}

// ListFiles mocks base method
func (m *MockFileService) ListFiles(arg0 *models.FilePaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", arg0)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles
func (mr *MockFileServiceMockRecorder) ListFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockFileService)(nil).ListFiles), arg0)
}

// GetFile mocks base method
func (m *MockFileService) GetFile(fileId *models.FileId) (*entities.FileEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", fileId)
	ret0, _ := ret[0].(*entities.FileEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile
func (mr *MockFileServiceMockRecorder) GetFile(fileId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockFileService)(nil).GetFile), fileId)
}

// UploadFile mocks base method
func (m *MockFileService) UploadFile(fileSource io.Reader, data *models.FileUpload) (*entities.FileEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", fileSource, data)
	ret0, _ := ret[0].(*entities.FileEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile
func (mr *MockFileServiceMockRecorder) UploadFile(fileSource, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockFileService)(nil).UploadFile), fileSource, data)
}

// DownloadFile mocks base method
func (m *MockFileService) DownloadFile(fileEntity *entities.FileEntity, fileDestination io.Writer) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", fileEntity, fileDestination)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// DownloadFile indicates an expected call of DownloadFile
func (mr *MockFileServiceMockRecorder) DownloadFile(fileEntity, fileDestination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockFileService)(nil).DownloadFile), fileEntity, fileDestination)
}

// UpdateFile mocks base method
func (m *MockFileService) UpdateFile(fileEntity *entities.FileEntity, data *models.FileUpdate) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFile", fileEntity, data)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// UpdateFile indicates an expected call of UpdateFile
func (mr *MockFileServiceMockRecorder) UpdateFile(fileEntity, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFile", reflect.TypeOf((*MockFileService)(nil).UpdateFile), fileEntity, data)
}

// DeleteFile mocks base method
func (m *MockFileService) DeleteFile(fileEntity *entities.FileEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", fileEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile
func (mr *MockFileServiceMockRecorder) DeleteFile(fileEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockFileService)(nil).DeleteFile), fileEntity)
}
