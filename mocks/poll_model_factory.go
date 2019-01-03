// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/poll_model_factory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPollModelFactory is a mock of PollModelFactory interface
type MockPollModelFactory struct {
	ctrl     *gomock.Controller
	recorder *MockPollModelFactoryMockRecorder
}

// MockPollModelFactoryMockRecorder is the mock recorder for MockPollModelFactory
type MockPollModelFactoryMockRecorder struct {
	mock *MockPollModelFactory
}

// NewMockPollModelFactory creates a new mock instance
func NewMockPollModelFactory(ctrl *gomock.Controller) *MockPollModelFactory {
	mock := &MockPollModelFactory{ctrl: ctrl}
	mock.recorder = &MockPollModelFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPollModelFactory) EXPECT() *MockPollModelFactoryMockRecorder {
	return m.recorder
}

// CreatePollPaginationQuery mocks base method
func (m *MockPollModelFactory) CreatePollPaginationQuery() *models.PollPaginationQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePollPaginationQuery")
	ret0, _ := ret[0].(*models.PollPaginationQuery)
	return ret0
}

// CreatePollPaginationQuery indicates an expected call of CreatePollPaginationQuery
func (mr *MockPollModelFactoryMockRecorder) CreatePollPaginationQuery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePollPaginationQuery", reflect.TypeOf((*MockPollModelFactory)(nil).CreatePollPaginationQuery))
}

// CreatePoll mocks base method
func (m *MockPollModelFactory) CreatePoll() *models.Poll {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePoll")
	ret0, _ := ret[0].(*models.Poll)
	return ret0
}

// CreatePoll indicates an expected call of CreatePoll
func (mr *MockPollModelFactoryMockRecorder) CreatePoll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePoll", reflect.TypeOf((*MockPollModelFactory)(nil).CreatePoll))
}

// CreatePollCreate mocks base method
func (m *MockPollModelFactory) CreatePollCreate() *models.PollCreate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePollCreate")
	ret0, _ := ret[0].(*models.PollCreate)
	return ret0
}

// CreatePollCreate indicates an expected call of CreatePollCreate
func (mr *MockPollModelFactoryMockRecorder) CreatePollCreate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePollCreate", reflect.TypeOf((*MockPollModelFactory)(nil).CreatePollCreate))
}

// CreatePollUpdate mocks base method
func (m *MockPollModelFactory) CreatePollUpdate() *models.PollUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePollUpdate")
	ret0, _ := ret[0].(*models.PollUpdate)
	return ret0
}

// CreatePollUpdate indicates an expected call of CreatePollUpdate
func (mr *MockPollModelFactoryMockRecorder) CreatePollUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePollUpdate", reflect.TypeOf((*MockPollModelFactory)(nil).CreatePollUpdate))
}
