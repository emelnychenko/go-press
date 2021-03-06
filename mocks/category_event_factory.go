// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/category_event_factory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	contracts "github.com/emelnychenko/go-press/contracts"
	entities "github.com/emelnychenko/go-press/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCategoryEventFactory is a mock of CategoryEventFactory interface
type MockCategoryEventFactory struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryEventFactoryMockRecorder
}

// MockCategoryEventFactoryMockRecorder is the mock recorder for MockCategoryEventFactory
type MockCategoryEventFactoryMockRecorder struct {
	mock *MockCategoryEventFactory
}

// NewMockCategoryEventFactory creates a new mock instance
func NewMockCategoryEventFactory(ctrl *gomock.Controller) *MockCategoryEventFactory {
	mock := &MockCategoryEventFactory{ctrl: ctrl}
	mock.recorder = &MockCategoryEventFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCategoryEventFactory) EXPECT() *MockCategoryEventFactoryMockRecorder {
	return m.recorder
}

// CreateCategoryCreatedEvent mocks base method
func (m *MockCategoryEventFactory) CreateCategoryCreatedEvent(arg0 *entities.CategoryEntity) contracts.CategoryEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategoryCreatedEvent", arg0)
	ret0, _ := ret[0].(contracts.CategoryEvent)
	return ret0
}

// CreateCategoryCreatedEvent indicates an expected call of CreateCategoryCreatedEvent
func (mr *MockCategoryEventFactoryMockRecorder) CreateCategoryCreatedEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategoryCreatedEvent", reflect.TypeOf((*MockCategoryEventFactory)(nil).CreateCategoryCreatedEvent), arg0)
}

// CreateCategoryUpdatedEvent mocks base method
func (m *MockCategoryEventFactory) CreateCategoryUpdatedEvent(arg0 *entities.CategoryEntity) contracts.CategoryEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategoryUpdatedEvent", arg0)
	ret0, _ := ret[0].(contracts.CategoryEvent)
	return ret0
}

// CreateCategoryUpdatedEvent indicates an expected call of CreateCategoryUpdatedEvent
func (mr *MockCategoryEventFactoryMockRecorder) CreateCategoryUpdatedEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategoryUpdatedEvent", reflect.TypeOf((*MockCategoryEventFactory)(nil).CreateCategoryUpdatedEvent), arg0)
}

// CreateCategoryDeletedEvent mocks base method
func (m *MockCategoryEventFactory) CreateCategoryDeletedEvent(arg0 *entities.CategoryEntity) contracts.CategoryEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategoryDeletedEvent", arg0)
	ret0, _ := ret[0].(contracts.CategoryEvent)
	return ret0
}

// CreateCategoryDeletedEvent indicates an expected call of CreateCategoryDeletedEvent
func (mr *MockCategoryEventFactoryMockRecorder) CreateCategoryDeletedEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategoryDeletedEvent", reflect.TypeOf((*MockCategoryEventFactory)(nil).CreateCategoryDeletedEvent), arg0)
}

// CreateCategoryParentChangedEvent mocks base method
func (m *MockCategoryEventFactory) CreateCategoryParentChangedEvent(arg0 *entities.CategoryEntity) contracts.CategoryEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategoryParentChangedEvent", arg0)
	ret0, _ := ret[0].(contracts.CategoryEvent)
	return ret0
}

// CreateCategoryParentChangedEvent indicates an expected call of CreateCategoryParentChangedEvent
func (mr *MockCategoryEventFactoryMockRecorder) CreateCategoryParentChangedEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategoryParentChangedEvent", reflect.TypeOf((*MockCategoryEventFactory)(nil).CreateCategoryParentChangedEvent), arg0)
}

// CreateCategoryParentRemovedEvent mocks base method
func (m *MockCategoryEventFactory) CreateCategoryParentRemovedEvent(arg0 *entities.CategoryEntity) contracts.CategoryEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategoryParentRemovedEvent", arg0)
	ret0, _ := ret[0].(contracts.CategoryEvent)
	return ret0
}

// CreateCategoryParentRemovedEvent indicates an expected call of CreateCategoryParentRemovedEvent
func (mr *MockCategoryEventFactoryMockRecorder) CreateCategoryParentRemovedEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategoryParentRemovedEvent", reflect.TypeOf((*MockCategoryEventFactory)(nil).CreateCategoryParentRemovedEvent), arg0)
}
