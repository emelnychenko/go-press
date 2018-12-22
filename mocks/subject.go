// Code generated by MockGen. DO NOT EDIT.
// Source: common/subject.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	enums "github.com/emelnychenko/go-press/enums"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSubject is a mock of Subject interface
type MockSubject struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectMockRecorder
}

// MockSubjectMockRecorder is the mock recorder for MockSubject
type MockSubjectMockRecorder struct {
	mock *MockSubject
}

// NewMockSubject creates a new mock instance
func NewMockSubject(ctrl *gomock.Controller) *MockSubject {
	mock := &MockSubject{ctrl: ctrl}
	mock.recorder = &MockSubjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubject) EXPECT() *MockSubjectMockRecorder {
	return m.recorder
}

// SubjectId mocks base method
func (m *MockSubject) SubjectId() *common.ModelId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubjectId")
	ret0, _ := ret[0].(*common.ModelId)
	return ret0
}

// SubjectId indicates an expected call of SubjectId
func (mr *MockSubjectMockRecorder) SubjectId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubjectId", reflect.TypeOf((*MockSubject)(nil).SubjectId))
}

// SubjectType mocks base method
func (m *MockSubject) SubjectType() enums.SubjectType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubjectType")
	ret0, _ := ret[0].(enums.SubjectType)
	return ret0
}

// SubjectType indicates an expected call of SubjectType
func (mr *MockSubjectMockRecorder) SubjectType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubjectType", reflect.TypeOf((*MockSubject)(nil).SubjectType))
}