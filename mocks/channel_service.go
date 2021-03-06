// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/channel_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockChannelService is a mock of ChannelService interface
type MockChannelService struct {
	ctrl     *gomock.Controller
	recorder *MockChannelServiceMockRecorder
}

// MockChannelServiceMockRecorder is the mock recorder for MockChannelService
type MockChannelServiceMockRecorder struct {
	mock *MockChannelService
}

// NewMockChannelService creates a new mock instance
func NewMockChannelService(ctrl *gomock.Controller) *MockChannelService {
	mock := &MockChannelService{ctrl: ctrl}
	mock.recorder = &MockChannelServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChannelService) EXPECT() *MockChannelServiceMockRecorder {
	return m.recorder
}

// ListChannels mocks base method
func (m *MockChannelService) ListChannels(channelPaginationQuery *models.ChannelPaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChannels", channelPaginationQuery)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListChannels indicates an expected call of ListChannels
func (mr *MockChannelServiceMockRecorder) ListChannels(channelPaginationQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChannels", reflect.TypeOf((*MockChannelService)(nil).ListChannels), channelPaginationQuery)
}

// GetChannel mocks base method
func (m *MockChannelService) GetChannel(channelId *models.ChannelId) (*entities.ChannelEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannel", channelId)
	ret0, _ := ret[0].(*entities.ChannelEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetChannel indicates an expected call of GetChannel
func (mr *MockChannelServiceMockRecorder) GetChannel(channelId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannel", reflect.TypeOf((*MockChannelService)(nil).GetChannel), channelId)
}

// CreateChannel mocks base method
func (m *MockChannelService) CreateChannel(data *models.ChannelCreate) (*entities.ChannelEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChannel", data)
	ret0, _ := ret[0].(*entities.ChannelEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// CreateChannel indicates an expected call of CreateChannel
func (mr *MockChannelServiceMockRecorder) CreateChannel(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChannel", reflect.TypeOf((*MockChannelService)(nil).CreateChannel), data)
}

// UpdateChannel mocks base method
func (m *MockChannelService) UpdateChannel(channelEntity *entities.ChannelEntity, data *models.ChannelUpdate) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateChannel", channelEntity, data)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// UpdateChannel indicates an expected call of UpdateChannel
func (mr *MockChannelServiceMockRecorder) UpdateChannel(channelEntity, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateChannel", reflect.TypeOf((*MockChannelService)(nil).UpdateChannel), channelEntity, data)
}

// DeleteChannel mocks base method
func (m *MockChannelService) DeleteChannel(channelEntity *entities.ChannelEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChannel", channelEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// DeleteChannel indicates an expected call of DeleteChannel
func (mr *MockChannelServiceMockRecorder) DeleteChannel(channelEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChannel", reflect.TypeOf((*MockChannelService)(nil).DeleteChannel), channelEntity)
}
