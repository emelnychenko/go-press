// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/banner_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/emelnychenko/go-press/entities"
	errors "github.com/emelnychenko/go-press/errors"
	models "github.com/emelnychenko/go-press/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBannerService is a mock of BannerService interface
type MockBannerService struct {
	ctrl     *gomock.Controller
	recorder *MockBannerServiceMockRecorder
}

// MockBannerServiceMockRecorder is the mock recorder for MockBannerService
type MockBannerServiceMockRecorder struct {
	mock *MockBannerService
}

// NewMockBannerService creates a new mock instance
func NewMockBannerService(ctrl *gomock.Controller) *MockBannerService {
	mock := &MockBannerService{ctrl: ctrl}
	mock.recorder = &MockBannerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBannerService) EXPECT() *MockBannerServiceMockRecorder {
	return m.recorder
}

// ListBanners mocks base method
func (m *MockBannerService) ListBanners(bannerPaginationQuery *models.BannerPaginationQuery) (*models.PaginationResult, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBanners", bannerPaginationQuery)
	ret0, _ := ret[0].(*models.PaginationResult)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// ListBanners indicates an expected call of ListBanners
func (mr *MockBannerServiceMockRecorder) ListBanners(bannerPaginationQuery interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBanners", reflect.TypeOf((*MockBannerService)(nil).ListBanners), bannerPaginationQuery)
}

// GetBanner mocks base method
func (m *MockBannerService) GetBanner(bannerId *models.BannerId) (*entities.BannerEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBanner", bannerId)
	ret0, _ := ret[0].(*entities.BannerEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// GetBanner indicates an expected call of GetBanner
func (mr *MockBannerServiceMockRecorder) GetBanner(bannerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBanner", reflect.TypeOf((*MockBannerService)(nil).GetBanner), bannerId)
}

// CreateBanner mocks base method
func (m *MockBannerService) CreateBanner(data *models.BannerCreate) (*entities.BannerEntity, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBanner", data)
	ret0, _ := ret[0].(*entities.BannerEntity)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// CreateBanner indicates an expected call of CreateBanner
func (mr *MockBannerServiceMockRecorder) CreateBanner(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBanner", reflect.TypeOf((*MockBannerService)(nil).CreateBanner), data)
}

// UpdateBanner mocks base method
func (m *MockBannerService) UpdateBanner(bannerEntity *entities.BannerEntity, data *models.BannerUpdate) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBanner", bannerEntity, data)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// UpdateBanner indicates an expected call of UpdateBanner
func (mr *MockBannerServiceMockRecorder) UpdateBanner(bannerEntity, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBanner", reflect.TypeOf((*MockBannerService)(nil).UpdateBanner), bannerEntity, data)
}

// DeleteBanner mocks base method
func (m *MockBannerService) DeleteBanner(bannerEntity *entities.BannerEntity) errors.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBanner", bannerEntity)
	ret0, _ := ret[0].(errors.Error)
	return ret0
}

// DeleteBanner indicates an expected call of DeleteBanner
func (mr *MockBannerServiceMockRecorder) DeleteBanner(bannerEntity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBanner", reflect.TypeOf((*MockBannerService)(nil).DeleteBanner), bannerEntity)
}
