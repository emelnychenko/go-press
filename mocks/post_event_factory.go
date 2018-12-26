// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_event_factory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	contracts "github.com/emelnychenko/go-press/contracts"
	entities "github.com/emelnychenko/go-press/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostEventFactory is a mock of PostEventFactory interface
type MockPostEventFactory struct {
	ctrl     *gomock.Controller
	recorder *MockPostEventFactoryMockRecorder
}

// MockPostEventFactoryMockRecorder is the mock recorder for MockPostEventFactory
type MockPostEventFactoryMockRecorder struct {
	mock *MockPostEventFactory
}

// NewMockPostEventFactory creates a new mock instance
func NewMockPostEventFactory(ctrl *gomock.Controller) *MockPostEventFactory {
	mock := &MockPostEventFactory{ctrl: ctrl}
	mock.recorder = &MockPostEventFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostEventFactory) EXPECT() *MockPostEventFactoryMockRecorder {
	return m.recorder
}

// CreatePostCreatedEvent mocks base method
func (m *MockPostEventFactory) CreatePostCreatedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostCreatedEvent", postEntity)
	ret0, _ := ret[0].(contracts.PostEvent)
	return ret0
}

// CreatePostCreatedEvent indicates an expected call of CreatePostCreatedEvent
func (mr *MockPostEventFactoryMockRecorder) CreatePostCreatedEvent(postEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostCreatedEvent", reflect.TypeOf((*MockPostEventFactory)(nil).CreatePostCreatedEvent), postEntity)
}

// CreatePostUpdatedEvent mocks base method
func (m *MockPostEventFactory) CreatePostUpdatedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostUpdatedEvent", postEntity)
	ret0, _ := ret[0].(contracts.PostEvent)
	return ret0
}

// CreatePostUpdatedEvent indicates an expected call of CreatePostUpdatedEvent
func (mr *MockPostEventFactoryMockRecorder) CreatePostUpdatedEvent(postEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostUpdatedEvent", reflect.TypeOf((*MockPostEventFactory)(nil).CreatePostUpdatedEvent), postEntity)
}

// CreatePostDeletedEvent mocks base method
func (m *MockPostEventFactory) CreatePostDeletedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostDeletedEvent", postEntity)
	ret0, _ := ret[0].(contracts.PostEvent)
	return ret0
}

// CreatePostDeletedEvent indicates an expected call of CreatePostDeletedEvent
func (mr *MockPostEventFactoryMockRecorder) CreatePostDeletedEvent(postEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostDeletedEvent", reflect.TypeOf((*MockPostEventFactory)(nil).CreatePostDeletedEvent), postEntity)
}

// CreatePostPublishedEvent mocks base method
func (m *MockPostEventFactory) CreatePostPublishedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostPublishedEvent", postEntity)
	ret0, _ := ret[0].(contracts.PostEvent)
	return ret0
}

// CreatePostPublishedEvent indicates an expected call of CreatePostPublishedEvent
func (mr *MockPostEventFactoryMockRecorder) CreatePostPublishedEvent(postEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostPublishedEvent", reflect.TypeOf((*MockPostEventFactory)(nil).CreatePostPublishedEvent), postEntity)
}

// CreatePostConcealedEvent mocks base method
func (m *MockPostEventFactory) CreatePostConcealedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostConcealedEvent", postEntity)
	ret0, _ := ret[0].(contracts.PostEvent)
	return ret0
}

// CreatePostConcealedEvent indicates an expected call of CreatePostConcealedEvent
func (mr *MockPostEventFactoryMockRecorder) CreatePostConcealedEvent(postEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostConcealedEvent", reflect.TypeOf((*MockPostEventFactory)(nil).CreatePostConcealedEvent), postEntity)
}
