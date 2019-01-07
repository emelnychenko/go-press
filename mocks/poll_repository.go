// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/poll_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPollRepository is a mock of PollRepository interface
type MockPollRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPollRepositoryMockRecorder
}

// MockPollRepositoryMockRecorder is the mock recorder for MockPollRepository
type MockPollRepositoryMockRecorder struct {
	mock *MockPollRepository
}

// NewMockPollRepository creates a new mock instance
func NewMockPollRepository(ctrl *gomock.Controller) *MockPollRepository {
	mock := &MockPollRepository{ctrl: ctrl}
	mock.recorder = &MockPollRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPollRepository) EXPECT() *MockPollRepositoryMockRecorder {
	return m.recorder
}

// ListPolls mocks base method
func (m *MockPollRepository) ListPolls(pollPaginationQuery *models.PollPaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPolls", pollPaginationQuery)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListPolls indicates an expected call of ListPolls
func (mr *MockPollRepositoryMockRecorder) ListPolls(pollPaginationQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolls", reflect.TypeOf((*MockPollRepository)(nil).ListPolls), pollPaginationQuery)
}

// GetPoll mocks base method
func (m *MockPollRepository) GetPoll(pollId *models.PollId) (*entities.PollEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPoll", pollId)
	ret0, _ := ret[0].(*entities.PollEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetPoll indicates an expected call of GetPoll
func (mr *MockPollRepositoryMockRecorder) GetPoll(pollId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPoll", reflect.TypeOf((*MockPollRepository)(nil).GetPoll), pollId)
}

// SavePoll mocks base method
func (m *MockPollRepository) SavePoll(pollEntity *entities.PollEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePoll", pollEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// SavePoll indicates an expected call of SavePoll
func (mr *MockPollRepositoryMockRecorder) SavePoll(pollEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePoll", reflect.TypeOf((*MockPollRepository)(nil).SavePoll), pollEntity)
}

// RemovePoll mocks base method
func (m *MockPollRepository) RemovePoll(pollEntity *entities.PollEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePoll", pollEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// RemovePoll indicates an expected call of RemovePoll
func (mr *MockPollRepositoryMockRecorder) RemovePoll(pollEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePoll", reflect.TypeOf((*MockPollRepository)(nil).RemovePoll), pollEntity)
}
