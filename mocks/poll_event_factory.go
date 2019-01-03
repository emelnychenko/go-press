// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/poll_event_factory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	contracts "github.com/emelnychenko/go-press/contracts"
	entities "github.com/emelnychenko/go-press/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPollEventFactory is a mock of PollEventFactory interface
type MockPollEventFactory struct {
	ctrl     *gomock.Controller
	recorder *MockPollEventFactoryMockRecorder
}

// MockPollEventFactoryMockRecorder is the mock recorder for MockPollEventFactory
type MockPollEventFactoryMockRecorder struct {
	mock *MockPollEventFactory
}

// NewMockPollEventFactory creates a new mock instance
func NewMockPollEventFactory(ctrl *gomock.Controller) *MockPollEventFactory {
	mock := &MockPollEventFactory{ctrl: ctrl}
	mock.recorder = &MockPollEventFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPollEventFactory) EXPECT() *MockPollEventFactoryMockRecorder {
	return m.recorder
}

// CreatePollCreatedEvent mocks base method
func (m *MockPollEventFactory) CreatePollCreatedEvent(pollEntity *entities.PollEntity) contracts.PollEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePollCreatedEvent", pollEntity)
	ret0, _ := ret[0].(contracts.PollEvent)
	return ret0
}

// CreatePollCreatedEvent indicates an expected call of CreatePollCreatedEvent
func (mr *MockPollEventFactoryMockRecorder) CreatePollCreatedEvent(pollEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePollCreatedEvent", reflect.TypeOf((*MockPollEventFactory)(nil).CreatePollCreatedEvent), pollEntity)
}

// CreatePollUpdatedEvent mocks base method
func (m *MockPollEventFactory) CreatePollUpdatedEvent(pollEntity *entities.PollEntity) contracts.PollEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePollUpdatedEvent", pollEntity)
	ret0, _ := ret[0].(contracts.PollEvent)
	return ret0
}

// CreatePollUpdatedEvent indicates an expected call of CreatePollUpdatedEvent
func (mr *MockPollEventFactoryMockRecorder) CreatePollUpdatedEvent(pollEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePollUpdatedEvent", reflect.TypeOf((*MockPollEventFactory)(nil).CreatePollUpdatedEvent), pollEntity)
}

// CreatePollDeletedEvent mocks base method
func (m *MockPollEventFactory) CreatePollDeletedEvent(pollEntity *entities.PollEntity) contracts.PollEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePollDeletedEvent", pollEntity)
	ret0, _ := ret[0].(contracts.PollEvent)
	return ret0
}

// CreatePollDeletedEvent indicates an expected call of CreatePollDeletedEvent
func (mr *MockPollEventFactoryMockRecorder) CreatePollDeletedEvent(pollEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePollDeletedEvent", reflect.TypeOf((*MockPollEventFactory)(nil).CreatePollDeletedEvent), pollEntity)
}
