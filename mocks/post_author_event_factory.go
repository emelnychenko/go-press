// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/post_author_event_factory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	contracts "github.com/emelnychenko/go-press/contracts"
	entities "github.com/emelnychenko/go-press/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPostAuthorEventFactory is a mock of PostAuthorEventFactory interface
type MockPostAuthorEventFactory struct {
	ctrl     *gomock.Controller
	recorder *MockPostAuthorEventFactoryMockRecorder
}

// MockPostAuthorEventFactoryMockRecorder is the mock recorder for MockPostAuthorEventFactory
type MockPostAuthorEventFactoryMockRecorder struct {
	mock *MockPostAuthorEventFactory
}

// NewMockPostAuthorEventFactory creates a new mock instance
func NewMockPostAuthorEventFactory(ctrl *gomock.Controller) *MockPostAuthorEventFactory {
	mock := &MockPostAuthorEventFactory{ctrl: ctrl}
	mock.recorder = &MockPostAuthorEventFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostAuthorEventFactory) EXPECT() *MockPostAuthorEventFactoryMockRecorder {
	return m.recorder
}

// CreatePostAuthorChangedEvent mocks base method
func (m *MockPostAuthorEventFactory) CreatePostAuthorChangedEvent(postEntity *entities.PostEntity, postAuthorEntity *entities.UserEntity) contracts.PostAuthorEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePostAuthorChangedEvent", postEntity, postAuthorEntity)
	ret0, _ := ret[0].(contracts.PostAuthorEvent)
	return ret0
}

// CreatePostAuthorChangedEvent indicates an expected call of CreatePostAuthorChangedEvent
func (mr *MockPostAuthorEventFactoryMockRecorder) CreatePostAuthorChangedEvent(postEntity, postAuthorEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePostAuthorChangedEvent", reflect.TypeOf((*MockPostAuthorEventFactory)(nil).CreatePostAuthorChangedEvent), postEntity, postAuthorEntity)
}
