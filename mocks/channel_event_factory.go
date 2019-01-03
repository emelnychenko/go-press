// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/channel_event_factory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	contracts "github.com/emelnychenko/go-press/contracts"
	entities "github.com/emelnychenko/go-press/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockChannelEventFactory is a mock of ChannelEventFactory interface
type MockChannelEventFactory struct {
	ctrl     *gomock.Controller
	recorder *MockChannelEventFactoryMockRecorder
}

// MockChannelEventFactoryMockRecorder is the mock recorder for MockChannelEventFactory
type MockChannelEventFactoryMockRecorder struct {
	mock *MockChannelEventFactory
}

// NewMockChannelEventFactory creates a new mock instance
func NewMockChannelEventFactory(ctrl *gomock.Controller) *MockChannelEventFactory {
	mock := &MockChannelEventFactory{ctrl: ctrl}
	mock.recorder = &MockChannelEventFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChannelEventFactory) EXPECT() *MockChannelEventFactoryMockRecorder {
	return m.recorder
}

// CreateChannelCreatedEvent mocks base method
func (m *MockChannelEventFactory) CreateChannelCreatedEvent(channelEntity *entities.ChannelEntity) contracts.ChannelEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChannelCreatedEvent", channelEntity)
	ret0, _ := ret[0].(contracts.ChannelEvent)
	return ret0
}

// CreateChannelCreatedEvent indicates an expected call of CreateChannelCreatedEvent
func (mr *MockChannelEventFactoryMockRecorder) CreateChannelCreatedEvent(channelEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChannelCreatedEvent", reflect.TypeOf((*MockChannelEventFactory)(nil).CreateChannelCreatedEvent), channelEntity)
}

// CreateChannelUpdatedEvent mocks base method
func (m *MockChannelEventFactory) CreateChannelUpdatedEvent(channelEntity *entities.ChannelEntity) contracts.ChannelEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChannelUpdatedEvent", channelEntity)
	ret0, _ := ret[0].(contracts.ChannelEvent)
	return ret0
}

// CreateChannelUpdatedEvent indicates an expected call of CreateChannelUpdatedEvent
func (mr *MockChannelEventFactoryMockRecorder) CreateChannelUpdatedEvent(channelEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChannelUpdatedEvent", reflect.TypeOf((*MockChannelEventFactory)(nil).CreateChannelUpdatedEvent), channelEntity)
}

// CreateChannelDeletedEvent mocks base method
func (m *MockChannelEventFactory) CreateChannelDeletedEvent(channelEntity *entities.ChannelEntity) contracts.ChannelEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChannelDeletedEvent", channelEntity)
	ret0, _ := ret[0].(contracts.ChannelEvent)
	return ret0
}

// CreateChannelDeletedEvent indicates an expected call of CreateChannelDeletedEvent
func (mr *MockChannelEventFactoryMockRecorder) CreateChannelDeletedEvent(channelEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChannelDeletedEvent", reflect.TypeOf((*MockChannelEventFactory)(nil).CreateChannelDeletedEvent), channelEntity)
}
