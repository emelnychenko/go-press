// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/tag_aggregator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTagAggregator is a mock of TagAggregator interface
type MockTagAggregator struct {
	ctrl     *gomock.Controller
	recorder *MockTagAggregatorMockRecorder
}

// MockTagAggregatorMockRecorder is the mock recorder for MockTagAggregator
type MockTagAggregatorMockRecorder struct {
	mock *MockTagAggregator
}

// NewMockTagAggregator creates a new mock instance
func NewMockTagAggregator(ctrl *gomock.Controller) *MockTagAggregator {
	mock := &MockTagAggregator{ctrl: ctrl}
	mock.recorder = &MockTagAggregatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagAggregator) EXPECT() *MockTagAggregatorMockRecorder {
	return m.recorder
}

// AggregateTag mocks base method
func (m *MockTagAggregator) AggregateTag(tagEntity *entities.TagEntity) *models.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateTag", tagEntity)
	ret0, _ := ret[0].(*models.Tag)
	return ret0
}

// AggregateTag indicates an expected call of AggregateTag
func (mr *MockTagAggregatorMockRecorder) AggregateTag(tagEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateTag", reflect.TypeOf((*MockTagAggregator)(nil).AggregateTag), tagEntity)
}

// AggregateTags mocks base method
func (m *MockTagAggregator) AggregateTags(tagEntities []*entities.TagEntity) []*models.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateTags", tagEntities)
	ret0, _ := ret[0].([]*models.Tag)
	return ret0
}

// AggregateTags indicates an expected call of AggregateTags
func (mr *MockTagAggregatorMockRecorder) AggregateTags(tagEntities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateTags", reflect.TypeOf((*MockTagAggregator)(nil).AggregateTags), tagEntities)
}

// AggregatePaginationResult mocks base method
func (m *MockTagAggregator) AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregatePaginationResult", entityPaginationResult)
	ret0, _ := ret[0].(*models.PaginationResult)
	return ret0
}

// AggregatePaginationResult indicates an expected call of AggregatePaginationResult
func (mr *MockTagAggregatorMockRecorder) AggregatePaginationResult(entityPaginationResult interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregatePaginationResult", reflect.TypeOf((*MockTagAggregator)(nil).AggregatePaginationResult), entityPaginationResult)
}
