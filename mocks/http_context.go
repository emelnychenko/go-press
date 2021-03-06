// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/http_context.go

// Package mocks is a generated GoMock package.
package mocks

import (
	errors "github.com/emelnychenko/go-press/errors"
	gomock "github.com/golang/mock/gomock"
	multipart "mime/multipart"
	http "net/http"
	reflect "reflect"
)

// MockHttpContext is a mock of HttpContext interface
type MockHttpContext struct {
	ctrl     *gomock.Controller
	recorder *MockHttpContextMockRecorder
}

// MockHttpContextMockRecorder is the mock recorder for MockHttpContext
type MockHttpContextMockRecorder struct {
	mock *MockHttpContext
}

// NewMockHttpContext creates a new mock instance
func NewMockHttpContext(ctrl *gomock.Controller) *MockHttpContext {
	mock := &MockHttpContext{ctrl: ctrl}
	mock.recorder = &MockHttpContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHttpContext) EXPECT() *MockHttpContextMockRecorder {
	return m.recorder
}

// Request mocks base method
func (m *MockHttpContext) Request() *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// Request indicates an expected call of Request
func (mr *MockHttpContextMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockHttpContext)(nil).Request))
}

// Response mocks base method
func (m *MockHttpContext) Response() http.ResponseWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Response")
	ret0, _ := ret[0].(http.ResponseWriter)
	return ret0
}

// Response indicates an expected call of Response
func (mr *MockHttpContextMockRecorder) Response() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Response", reflect.TypeOf((*MockHttpContext)(nil).Response))
}

// Parameter mocks base method
func (m *MockHttpContext) Parameter(parameterName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameter", parameterName)
	ret0, _ := ret[0].(string)
	return ret0
}

// Parameter indicates an expected call of Parameter
func (mr *MockHttpContextMockRecorder) Parameter(parameterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameter", reflect.TypeOf((*MockHttpContext)(nil).Parameter), parameterName)
}

// FormFile mocks base method
func (m *MockHttpContext) FormFile(formFileName string) (*multipart.FileHeader, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormFile", formFileName)
	ret0, _ := ret[0].(*multipart.FileHeader)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// FormFile indicates an expected call of FormFile
func (mr *MockHttpContextMockRecorder) FormFile(formFileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormFile", reflect.TypeOf((*MockHttpContext)(nil).FormFile), formFileName)
}

// BindModel mocks base method
func (m *MockHttpContext) BindModel(data interface{}) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindModel", data)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// BindModel indicates an expected call of BindModel
func (mr *MockHttpContextMockRecorder) BindModel(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindModel", reflect.TypeOf((*MockHttpContext)(nil).BindModel), data)
}
