// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/poll_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPollService is a mock of PollService interface
type MockPollService struct {
	ctrl     *gomock.Controller
	recorder *MockPollServiceMockRecorder
}

// MockPollServiceMockRecorder is the mock recorder for MockPollService
type MockPollServiceMockRecorder struct {
	mock *MockPollService
}

// NewMockPollService creates a new mock instance
func NewMockPollService(ctrl *gomock.Controller) *MockPollService {
	mock := &MockPollService{ctrl: ctrl}
	mock.recorder = &MockPollServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPollService) EXPECT() *MockPollServiceMockRecorder {
	return m.recorder
}

// ListPolls mocks base method
func (m *MockPollService) ListPolls(pollPaginationQuery *models.PollPaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPolls", pollPaginationQuery)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListPolls indicates an expected call of ListPolls
func (mr *MockPollServiceMockRecorder) ListPolls(pollPaginationQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolls", reflect.TypeOf((*MockPollService)(nil).ListPolls), pollPaginationQuery)
}

// GetPoll mocks base method
func (m *MockPollService) GetPoll(pollId *models.PollId) (*entities.PollEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoll", pollId)
	ret0, _ := ret[0].(*entities.PollEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetPoll indicates an expected call of GetPoll
func (mr *MockPollServiceMockRecorder) GetPoll(pollId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoll", reflect.TypeOf((*MockPollService)(nil).GetPoll), pollId)
}

// CreatePoll mocks base method
func (m *MockPollService) CreatePoll(data *models.PollCreate) (*entities.PollEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePoll", data)
	ret0, _ := ret[0].(*entities.PollEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// CreatePoll indicates an expected call of CreatePoll
func (mr *MockPollServiceMockRecorder) CreatePoll(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePoll", reflect.TypeOf((*MockPollService)(nil).CreatePoll), data)
}

// UpdatePoll mocks base method
func (m *MockPollService) UpdatePoll(pollEntity *entities.PollEntity, data *models.PollUpdate) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePoll", pollEntity, data)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// UpdatePoll indicates an expected call of UpdatePoll
func (mr *MockPollServiceMockRecorder) UpdatePoll(pollEntity, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePoll", reflect.TypeOf((*MockPollService)(nil).UpdatePoll), pollEntity, data)
}

// DeletePoll mocks base method
func (m *MockPollService) DeletePoll(pollEntity *entities.PollEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePoll", pollEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// DeletePoll indicates an expected call of DeletePoll
func (mr *MockPollServiceMockRecorder) DeletePoll(pollEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePoll", reflect.TypeOf((*MockPollService)(nil).DeletePoll), pollEntity)
}
