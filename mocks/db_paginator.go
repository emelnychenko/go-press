// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/db_paginator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
	reflect "reflect"
)

// MockDbPaginator is a mock of DbPaginator interface
type MockDbPaginator struct {
	ctrl     *gomock.Controller
	recorder *MockDbPaginatorMockRecorder
}

// MockDbPaginatorMockRecorder is the mock recorder for MockDbPaginator
type MockDbPaginatorMockRecorder struct {
	mock *MockDbPaginator
}

// NewMockDbPaginator creates a new mock instance
func NewMockDbPaginator(ctrl *gomock.Controller) *MockDbPaginator {
	mock := &MockDbPaginator{ctrl: ctrl}
	mock.recorder = &MockDbPaginatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDbPaginator) EXPECT() *MockDbPaginatorMockRecorder {
	return m.recorder
}

// Paginate mocks base method
func (m *MockDbPaginator) Paginate(db *gorm.DB, paginationQuery *models.PaginationQuery, paginationData interface{}, paginationTotal *int) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Paginate", db, paginationQuery, paginationData, paginationTotal)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// Paginate indicates an expected call of Paginate
func (mr *MockDbPaginatorMockRecorder) Paginate(db, paginationQuery, paginationData, paginationTotal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paginate", reflect.TypeOf((*MockDbPaginator)(nil).Paginate), db, paginationQuery, paginationData, paginationTotal)
}
