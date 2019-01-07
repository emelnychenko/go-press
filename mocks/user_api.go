// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/user_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserApi is a mock of UserApi interface
type MockUserApi struct {
	ctrl     *gomock.Controller
	recorder *MockUserApiMockRecorder
}

// MockUserApiMockRecorder is the mock recorder for MockUserApi
type MockUserApiMockRecorder struct {
	mock *MockUserApi
}

// NewMockUserApi creates a new mock instance
func NewMockUserApi(ctrl *gomock.Controller) *MockUserApi {
	mock := &MockUserApi{ctrl: ctrl}
	mock.recorder = &MockUserApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserApi) EXPECT() *MockUserApiMockRecorder {
	return m.recorder
}

// ListUsers mocks base method
func (m *MockUserApi) ListUsers(userPaginationQuery *models.UserPaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", userPaginationQuery)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers
func (mr *MockUserApiMockRecorder) ListUsers(userPaginationQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockUserApi)(nil).ListUsers), userPaginationQuery)
}

// GetUser mocks base method
func (m *MockUserApi) GetUser(userId *models.UserId) (*models.User, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", userId)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserApiMockRecorder) GetUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserApi)(nil).GetUser), userId)
}

// CreateUser mocks base method
func (m *MockUserApi) CreateUser(data *models.UserCreate) (*models.User, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", data)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserApiMockRecorder) CreateUser(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserApi)(nil).CreateUser), data)
}

// UpdateUser mocks base method
func (m *MockUserApi) UpdateUser(userId *models.UserId, data *models.UserUpdate) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userId, data)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUserApiMockRecorder) UpdateUser(userId, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserApi)(nil).UpdateUser), userId, data)
}

// VerifyUser mocks base method
func (m *MockUserApi) VerifyUser(userId *models.UserId) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUser", userId)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// VerifyUser indicates an expected call of VerifyUser
func (mr *MockUserApiMockRecorder) VerifyUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUser", reflect.TypeOf((*MockUserApi)(nil).VerifyUser), userId)
}

// ChangeUserIdentity mocks base method
func (m *MockUserApi) ChangeUserIdentity(userId *models.UserId, data *models.UserChangeIdentity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserIdentity", userId, data)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// ChangeUserIdentity indicates an expected call of ChangeUserIdentity
func (mr *MockUserApiMockRecorder) ChangeUserIdentity(userId, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserIdentity", reflect.TypeOf((*MockUserApi)(nil).ChangeUserIdentity), userId, data)
}

// ChangeUserPassword mocks base method
func (m *MockUserApi) ChangeUserPassword(userId *models.UserId, data *models.UserChangePassword) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserPassword", userId, data)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// ChangeUserPassword indicates an expected call of ChangeUserPassword
func (mr *MockUserApiMockRecorder) ChangeUserPassword(userId, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserPassword", reflect.TypeOf((*MockUserApi)(nil).ChangeUserPassword), userId, data)
}

// DeleteUser mocks base method
func (m *MockUserApi) DeleteUser(userId *models.UserId) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userId)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockUserApiMockRecorder) DeleteUser(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserApi)(nil).DeleteUser), userId)
}
