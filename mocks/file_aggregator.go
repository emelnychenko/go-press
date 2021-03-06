// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/file_aggregator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFileAggregator is a mock of FileAggregator interface
type MockFileAggregator struct {
	ctrl     *gomock.Controller
	recorder *MockFileAggregatorMockRecorder
}

// MockFileAggregatorMockRecorder is the mock recorder for MockFileAggregator
type MockFileAggregatorMockRecorder struct {
	mock *MockFileAggregator
}

// NewMockFileAggregator creates a new mock instance
func NewMockFileAggregator(ctrl *gomock.Controller) *MockFileAggregator {
	mock := &MockFileAggregator{ctrl: ctrl}
	mock.recorder = &MockFileAggregatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileAggregator) EXPECT() *MockFileAggregatorMockRecorder {
	return m.recorder
}

// AggregateFile mocks base method
func (m *MockFileAggregator) AggregateFile(fileEntity *entities.FileEntity) *models.File {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateFile", fileEntity)
	ret0, _ := ret[0].(*models.File)
	return ret0
}

// AggregateFile indicates an expected call of AggregateFile
func (mr *MockFileAggregatorMockRecorder) AggregateFile(fileEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateFile", reflect.TypeOf((*MockFileAggregator)(nil).AggregateFile), fileEntity)
}

// AggregateFiles mocks base method
func (m *MockFileAggregator) AggregateFiles(fileEntities []*entities.FileEntity) []*models.File {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateFiles", fileEntities)
	ret0, _ := ret[0].([]*models.File)
	return ret0
}

// AggregateFiles indicates an expected call of AggregateFiles
func (mr *MockFileAggregatorMockRecorder) AggregateFiles(fileEntities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateFiles", reflect.TypeOf((*MockFileAggregator)(nil).AggregateFiles), fileEntities)
}

// AggregatePaginationResult mocks base method
func (m *MockFileAggregator) AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregatePaginationResult", entityPaginationResult)
	ret0, _ := ret[0].(*models.PaginationResult)
	return ret0
}

// AggregatePaginationResult indicates an expected call of AggregatePaginationResult
func (mr *MockFileAggregatorMockRecorder) AggregatePaginationResult(entityPaginationResult interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregatePaginationResult", reflect.TypeOf((*MockFileAggregator)(nil).AggregatePaginationResult), entityPaginationResult)
}
