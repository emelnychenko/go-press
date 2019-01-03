// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/tag_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	common "github.com/emelnychenko/go-press/common"
	entities "github.com/emelnychenko/go-press/entities"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTagRepository is a mock of TagRepository interface
type MockTagRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTagRepositoryMockRecorder
}

// MockTagRepositoryMockRecorder is the mock recorder for MockTagRepository
type MockTagRepositoryMockRecorder struct {
	mock *MockTagRepository
}

// NewMockTagRepository creates a new mock instance
func NewMockTagRepository(ctrl *gomock.Controller) *MockTagRepository {
	mock := &MockTagRepository{ctrl: ctrl}
	mock.recorder = &MockTagRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTagRepository) EXPECT() *MockTagRepositoryMockRecorder {
	return m.recorder
}

// ListTags mocks base method
func (m *MockTagRepository) ListTags(tagPaginationQuery *models.TagPaginationQuery) (*models.PaginationResult, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTags", tagPaginationQuery)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags
func (mr *MockTagRepositoryMockRecorder) ListTags(tagPaginationQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockTagRepository)(nil).ListTags), tagPaginationQuery)
}

// GetTag mocks base method
func (m *MockTagRepository) GetTag(tagId *models.TagId) (*entities.TagEntity, common.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTag", tagId)
	ret0, _ := ret[0].(*entities.TagEntity)
	ret1, _ := ret[1].(common.Error)
	return ret0, ret1
}

// GetTag indicates an expected call of GetTag
func (mr *MockTagRepositoryMockRecorder) GetTag(tagId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTag", reflect.TypeOf((*MockTagRepository)(nil).GetTag), tagId)
}

// SaveTag mocks base method
func (m *MockTagRepository) SaveTag(tagEntity *entities.TagEntity) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTag", tagEntity)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// SaveTag indicates an expected call of SaveTag
func (mr *MockTagRepositoryMockRecorder) SaveTag(tagEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTag", reflect.TypeOf((*MockTagRepository)(nil).SaveTag), tagEntity)
}

// RemoveTag mocks base method
func (m *MockTagRepository) RemoveTag(tagEntity *entities.TagEntity) common.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTag", tagEntity)
	ret0, _ := ret[0].(common.Error)
	return ret0
}

// RemoveTag indicates an expected call of RemoveTag
func (mr *MockTagRepositoryMockRecorder) RemoveTag(tagEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTag", reflect.TypeOf((*MockTagRepository)(nil).RemoveTag), tagEntity)
}
