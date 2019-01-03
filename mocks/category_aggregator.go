// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/category_aggregator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCategoryAggregator is a mock of CategoryAggregator interface
type MockCategoryAggregator struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryAggregatorMockRecorder
}

// MockCategoryAggregatorMockRecorder is the mock recorder for MockCategoryAggregator
type MockCategoryAggregatorMockRecorder struct {
	mock *MockCategoryAggregator
}

// NewMockCategoryAggregator creates a new mock instance
func NewMockCategoryAggregator(ctrl *gomock.Controller) *MockCategoryAggregator {
	mock := &MockCategoryAggregator{ctrl: ctrl}
	mock.recorder = &MockCategoryAggregatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCategoryAggregator) EXPECT() *MockCategoryAggregatorMockRecorder {
	return m.recorder
}

// AggregateCategory mocks base method
func (m *MockCategoryAggregator) AggregateCategory(categoryEntity *entities.CategoryEntity) *models.Category {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateCategory", categoryEntity)
	ret0, _ := ret[0].(*models.Category)
	return ret0
}

// AggregateCategory indicates an expected call of AggregateCategory
func (mr *MockCategoryAggregatorMockRecorder) AggregateCategory(categoryEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateCategory", reflect.TypeOf((*MockCategoryAggregator)(nil).AggregateCategory), categoryEntity)
}

// AggregateCategories mocks base method
func (m *MockCategoryAggregator) AggregateCategories(categoryEntities []*entities.CategoryEntity) []*models.Category {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateCategories", categoryEntities)
	ret0, _ := ret[0].([]*models.Category)
	return ret0
}

// AggregateCategories indicates an expected call of AggregateCategories
func (mr *MockCategoryAggregatorMockRecorder) AggregateCategories(categoryEntities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateCategories", reflect.TypeOf((*MockCategoryAggregator)(nil).AggregateCategories), categoryEntities)
}

// AggregatePaginationResult mocks base method
func (m *MockCategoryAggregator) AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregatePaginationResult", entityPaginationResult)
	ret0, _ := ret[0].(*models.PaginationResult)
	return ret0
}

// AggregatePaginationResult indicates an expected call of AggregatePaginationResult
func (mr *MockCategoryAggregatorMockRecorder) AggregatePaginationResult(entityPaginationResult interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregatePaginationResult", reflect.TypeOf((*MockCategoryAggregator)(nil).AggregatePaginationResult), entityPaginationResult)
}
