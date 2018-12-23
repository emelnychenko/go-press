// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/aws_s3_factory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	session "github.com/aws/aws-sdk-go/aws/session"
	s3iface "github.com/aws/aws-sdk-go/service/s3/s3iface"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAwsS3Factory is a mock of AwsS3Factory interface
type MockAwsS3Factory struct {
	ctrl     *gomock.Controller
	recorder *MockAwsS3FactoryMockRecorder
}

// MockAwsS3FactoryMockRecorder is the mock recorder for MockAwsS3Factory
type MockAwsS3FactoryMockRecorder struct {
	mock *MockAwsS3Factory
}

// NewMockAwsS3Factory creates a new mock instance
func NewMockAwsS3Factory(ctrl *gomock.Controller) *MockAwsS3Factory {
	mock := &MockAwsS3Factory{ctrl: ctrl}
	mock.recorder = &MockAwsS3FactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAwsS3Factory) EXPECT() *MockAwsS3FactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAwsS3Factory) Create(sess *session.Session) s3iface.S3API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", sess)
	ret0, _ := ret[0].(s3iface.S3API)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockAwsS3FactoryMockRecorder) Create(sess interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAwsS3Factory)(nil).Create), sess)
}