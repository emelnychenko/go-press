// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/user_entity_factory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserEntityFactory is a mock of UserEntityFactory interface
type MockUserEntityFactory struct {
	ctrl     *gomock.Controller
	recorder *MockUserEntityFactoryMockRecorder
}

// MockUserEntityFactoryMockRecorder is the mock recorder for MockUserEntityFactory
type MockUserEntityFactoryMockRecorder struct {
	mock *MockUserEntityFactory
}

// NewMockUserEntityFactory creates a new mock instance
func NewMockUserEntityFactory(ctrl *gomock.Controller) *MockUserEntityFactory {
	mock := &MockUserEntityFactory{ctrl: ctrl}
	mock.recorder = &MockUserEntityFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserEntityFactory) EXPECT() *MockUserEntityFactoryMockRecorder {
	return m.recorder
}

// CreateUserEntity mocks base method
func (m *MockUserEntityFactory) CreateUserEntity() *entities.UserEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserEntity")
	ret0, _ := ret[0].(*entities.UserEntity)
	return ret0
}

// CreateUserEntity indicates an expected call of CreateUserEntity
func (mr *MockUserEntityFactoryMockRecorder) CreateUserEntity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserEntity", reflect.TypeOf((*MockUserEntityFactory)(nil).CreateUserEntity))
}
