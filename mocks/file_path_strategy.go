// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/file_path_strategy.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFilePathStrategy is a mock of FilePathStrategy interface
type MockFilePathStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockFilePathStrategyMockRecorder
}

// MockFilePathStrategyMockRecorder is the mock recorder for MockFilePathStrategy
type MockFilePathStrategyMockRecorder struct {
	mock *MockFilePathStrategy
}

// NewMockFilePathStrategy creates a new mock instance
func NewMockFilePathStrategy(ctrl *gomock.Controller) *MockFilePathStrategy {
	mock := &MockFilePathStrategy{ctrl: ctrl}
	mock.recorder = &MockFilePathStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFilePathStrategy) EXPECT() *MockFilePathStrategyMockRecorder {
	return m.recorder
}

// BuildPath mocks base method
func (m *MockFilePathStrategy) BuildPath(fileEntity *entities.FileEntity) (string, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildPath", fileEntity)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// BuildPath indicates an expected call of BuildPath
func (mr *MockFilePathStrategyMockRecorder) BuildPath(fileEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildPath", reflect.TypeOf((*MockFilePathStrategy)(nil).BuildPath), fileEntity)
}
