// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/category_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCategoryRepository is a mock of CategoryRepository interface
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// ListCategories mocks base method
func (m *MockCategoryRepository) ListCategories(categoryPaginationQuery *models.CategoryPaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCategories", categoryPaginationQuery)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListCategories indicates an expected call of ListCategories
func (mr *MockCategoryRepositoryMockRecorder) ListCategories(categoryPaginationQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCategories", reflect.TypeOf((*MockCategoryRepository)(nil).ListCategories), categoryPaginationQuery)
}

// GetCategories mocks base method
func (m *MockCategoryRepository) GetCategories() ([]*entities.CategoryEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories")
	ret0, _ := ret[0].([]*entities.CategoryEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories
func (mr *MockCategoryRepositoryMockRecorder) GetCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategories))
}

// GetCategoriesExcept mocks base method
func (m *MockCategoryRepository) GetCategoriesExcept(categoryEntity *entities.CategoryEntity) ([]*entities.CategoryEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoriesExcept", categoryEntity)
	ret0, _ := ret[0].([]*entities.CategoryEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetCategoriesExcept indicates an expected call of GetCategoriesExcept
func (mr *MockCategoryRepositoryMockRecorder) GetCategoriesExcept(categoryEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoriesExcept", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategoriesExcept), categoryEntity)
}

// GetCategoriesTree mocks base method
func (m *MockCategoryRepository) GetCategoriesTree() (*entities.CategoryEntityTree, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoriesTree")
	ret0, _ := ret[0].(*entities.CategoryEntityTree)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetCategoriesTree indicates an expected call of GetCategoriesTree
func (mr *MockCategoryRepositoryMockRecorder) GetCategoriesTree() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoriesTree", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategoriesTree))
}

// GetCategory mocks base method
func (m *MockCategoryRepository) GetCategory(categoryId *models.CategoryId) (*entities.CategoryEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", categoryId)
	ret0, _ := ret[0].(*entities.CategoryEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory
func (mr *MockCategoryRepositoryMockRecorder) GetCategory(categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategory), categoryId)
}

// GetCategoryTree mocks base method
func (m *MockCategoryRepository) GetCategoryTree(categoryId *models.CategoryId) (*entities.CategoryEntityTree, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryTree", categoryId)
	ret0, _ := ret[0].(*entities.CategoryEntityTree)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetCategoryTree indicates an expected call of GetCategoryTree
func (mr *MockCategoryRepositoryMockRecorder) GetCategoryTree(categoryId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryTree", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategoryTree), categoryId)
}

// SaveCategory mocks base method
func (m *MockCategoryRepository) SaveCategory(categoryEntity *entities.CategoryEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCategory", categoryEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// SaveCategory indicates an expected call of SaveCategory
func (mr *MockCategoryRepositoryMockRecorder) SaveCategory(categoryEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCategory", reflect.TypeOf((*MockCategoryRepository)(nil).SaveCategory), categoryEntity)
}

// RemoveCategory mocks base method
func (m *MockCategoryRepository) RemoveCategory(categoryEntity *entities.CategoryEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCategory", categoryEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// RemoveCategory indicates an expected call of RemoveCategory
func (mr *MockCategoryRepositoryMockRecorder) RemoveCategory(categoryEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCategory", reflect.TypeOf((*MockCategoryRepository)(nil).RemoveCategory), categoryEntity)
}

// GetCategoryXrefs mocks base method
func (m *MockCategoryRepository) GetCategoryXrefs(arg0 *entities.CategoryEntity) ([]*entities.CategoryXrefEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryXrefs", arg0)
	ret0, _ := ret[0].([]*entities.CategoryXrefEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetCategoryXrefs indicates an expected call of GetCategoryXrefs
func (mr *MockCategoryRepositoryMockRecorder) GetCategoryXrefs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryXrefs", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategoryXrefs), arg0)
}

// GetCategoryObjectXrefs mocks base method
func (m *MockCategoryRepository) GetCategoryObjectXrefs(arg0 models.Object) ([]*entities.CategoryXrefEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryObjectXrefs", arg0)
	ret0, _ := ret[0].([]*entities.CategoryXrefEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetCategoryObjectXrefs indicates an expected call of GetCategoryObjectXrefs
func (mr *MockCategoryRepositoryMockRecorder) GetCategoryObjectXrefs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryObjectXrefs", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategoryObjectXrefs), arg0)
}

// GetCategoryXref mocks base method
func (m *MockCategoryRepository) GetCategoryXref(arg0 *entities.CategoryEntity, arg1 models.Object) (*entities.CategoryXrefEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryXref", arg0, arg1)
	ret0, _ := ret[0].(*entities.CategoryXrefEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetCategoryXref indicates an expected call of GetCategoryXref
func (mr *MockCategoryRepositoryMockRecorder) GetCategoryXref(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryXref", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategoryXref), arg0, arg1)
}

// SaveCategoryXref mocks base method
func (m *MockCategoryRepository) SaveCategoryXref(arg0 *entities.CategoryXrefEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCategoryXref", arg0)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// SaveCategoryXref indicates an expected call of SaveCategoryXref
func (mr *MockCategoryRepositoryMockRecorder) SaveCategoryXref(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCategoryXref", reflect.TypeOf((*MockCategoryRepository)(nil).SaveCategoryXref), arg0)
}

// RemoveCategoryXref mocks base method
func (m *MockCategoryRepository) RemoveCategoryXref(arg0 *entities.CategoryXrefEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCategoryXref", arg0)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// RemoveCategoryXref indicates an expected call of RemoveCategoryXref
func (mr *MockCategoryRepositoryMockRecorder) RemoveCategoryXref(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCategoryXref", reflect.TypeOf((*MockCategoryRepository)(nil).RemoveCategoryXref), arg0)
}

// ListObjectCategories mocks base method
func (m *MockCategoryRepository) ListObjectCategories(arg0 models.Object, arg1 *models.CategoryPaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectCategories", arg0, arg1)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListObjectCategories indicates an expected call of ListObjectCategories
func (mr *MockCategoryRepositoryMockRecorder) ListObjectCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectCategories", reflect.TypeOf((*MockCategoryRepository)(nil).ListObjectCategories), arg0, arg1)
}
