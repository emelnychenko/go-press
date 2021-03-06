// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/poll_aggregator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPollAggregator is a mock of PollAggregator interface
type MockPollAggregator struct {
	ctrl     *gomock.Controller
	recorder *MockPollAggregatorMockRecorder
}

// MockPollAggregatorMockRecorder is the mock recorder for MockPollAggregator
type MockPollAggregatorMockRecorder struct {
	mock *MockPollAggregator
}

// NewMockPollAggregator creates a new mock instance
func NewMockPollAggregator(ctrl *gomock.Controller) *MockPollAggregator {
	mock := &MockPollAggregator{ctrl: ctrl}
	mock.recorder = &MockPollAggregatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPollAggregator) EXPECT() *MockPollAggregatorMockRecorder {
	return m.recorder
}

// AggregatePoll mocks base method
func (m *MockPollAggregator) AggregatePoll(pollEntity *entities.PollEntity) *models.Poll {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregatePoll", pollEntity)
	ret0, _ := ret[0].(*models.Poll)
	return ret0
}

// AggregatePoll indicates an expected call of AggregatePoll
func (mr *MockPollAggregatorMockRecorder) AggregatePoll(pollEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregatePoll", reflect.TypeOf((*MockPollAggregator)(nil).AggregatePoll), pollEntity)
}

// AggregatePolls mocks base method
func (m *MockPollAggregator) AggregatePolls(pollEntities []*entities.PollEntity) []*models.Poll {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregatePolls", pollEntities)
	ret0, _ := ret[0].([]*models.Poll)
	return ret0
}

// AggregatePolls indicates an expected call of AggregatePolls
func (mr *MockPollAggregatorMockRecorder) AggregatePolls(pollEntities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregatePolls", reflect.TypeOf((*MockPollAggregator)(nil).AggregatePolls), pollEntities)
}

// AggregatePaginationResult mocks base method
func (m *MockPollAggregator) AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregatePaginationResult", entityPaginationResult)
	ret0, _ := ret[0].(*models.PaginationResult)
	return ret0
}

// AggregatePaginationResult indicates an expected call of AggregatePaginationResult
func (mr *MockPollAggregatorMockRecorder) AggregatePaginationResult(entityPaginationResult interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregatePaginationResult", reflect.TypeOf((*MockPollAggregator)(nil).AggregatePaginationResult), entityPaginationResult)
}
