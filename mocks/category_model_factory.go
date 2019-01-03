// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/category_model_factory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCategoryModelFactory is a mock of CategoryModelFactory interface
type MockCategoryModelFactory struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryModelFactoryMockRecorder
}

// MockCategoryModelFactoryMockRecorder is the mock recorder for MockCategoryModelFactory
type MockCategoryModelFactoryMockRecorder struct {
	mock *MockCategoryModelFactory
}

// NewMockCategoryModelFactory creates a new mock instance
func NewMockCategoryModelFactory(ctrl *gomock.Controller) *MockCategoryModelFactory {
	mock := &MockCategoryModelFactory{ctrl: ctrl}
	mock.recorder = &MockCategoryModelFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCategoryModelFactory) EXPECT() *MockCategoryModelFactoryMockRecorder {
	return m.recorder
}

// CreateCategoryPaginationQuery mocks base method
func (m *MockCategoryModelFactory) CreateCategoryPaginationQuery() *models.CategoryPaginationQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategoryPaginationQuery")
	ret0, _ := ret[0].(*models.CategoryPaginationQuery)
	return ret0
}

// CreateCategoryPaginationQuery indicates an expected call of CreateCategoryPaginationQuery
func (mr *MockCategoryModelFactoryMockRecorder) CreateCategoryPaginationQuery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategoryPaginationQuery", reflect.TypeOf((*MockCategoryModelFactory)(nil).CreateCategoryPaginationQuery))
}

// CreateCategory mocks base method
func (m *MockCategoryModelFactory) CreateCategory() *models.Category {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory")
	ret0, _ := ret[0].(*models.Category)
	return ret0
}

// CreateCategory indicates an expected call of CreateCategory
func (mr *MockCategoryModelFactoryMockRecorder) CreateCategory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockCategoryModelFactory)(nil).CreateCategory))
}

// CreateCategoryCreate mocks base method
func (m *MockCategoryModelFactory) CreateCategoryCreate() *models.CategoryCreate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategoryCreate")
	ret0, _ := ret[0].(*models.CategoryCreate)
	return ret0
}

// CreateCategoryCreate indicates an expected call of CreateCategoryCreate
func (mr *MockCategoryModelFactoryMockRecorder) CreateCategoryCreate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategoryCreate", reflect.TypeOf((*MockCategoryModelFactory)(nil).CreateCategoryCreate))
}

// CreateCategoryUpdate mocks base method
func (m *MockCategoryModelFactory) CreateCategoryUpdate() *models.CategoryUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategoryUpdate")
	ret0, _ := ret[0].(*models.CategoryUpdate)
	return ret0
}

// CreateCategoryUpdate indicates an expected call of CreateCategoryUpdate
func (mr *MockCategoryModelFactoryMockRecorder) CreateCategoryUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategoryUpdate", reflect.TypeOf((*MockCategoryModelFactory)(nil).CreateCategoryUpdate))
}
