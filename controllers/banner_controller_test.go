package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewBannerController", func(t *testing.T) {
		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerApi := mocks.NewMockBannerApi(ctrl)
		bannerController, isBannerController := NewBannerController(
			bannerHttpHelper,
			bannerModelFactory,
			bannerApi,
		).(*bannerControllerImpl)

		assert.True(t, isBannerController)
		assert.Equal(t, bannerHttpHelper, bannerController.bannerHttpHelper)
		assert.Equal(t, bannerModelFactory, bannerController.bannerModelFactory)
		assert.Equal(t, bannerApi, bannerController.bannerApi)
	})

	t.Run("ListBanners", func(t *testing.T) {
		bannerPaginationQuery := new(models.BannerPaginationQuery)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBannerPaginationQuery().Return(bannerPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(bannerPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(bannerPaginationQuery).Return(nil)

		var paginationResult *models.PaginationResult
		bannerApi := mocks.NewMockBannerApi(ctrl)
		bannerApi.EXPECT().ListBanners(bannerPaginationQuery).Return(paginationResult, nil)

		bannerController := &bannerControllerImpl{
			bannerModelFactory: bannerModelFactory,
			bannerApi:          bannerApi,
		}
		response, err := bannerController.ListBanners(httpContext)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListBanners:BindPaginationQueryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		bannerPaginationQuery := new(models.BannerPaginationQuery)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBannerPaginationQuery().Return(bannerPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(bannerPaginationQuery.PaginationQuery).Return(systemErr)

		bannerController := &bannerControllerImpl{
			bannerModelFactory: bannerModelFactory,
		}
		response, err := bannerController.ListBanners(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListBanners:BindBannerPaginationQueryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		bannerPaginationQuery := new(models.BannerPaginationQuery)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBannerPaginationQuery().Return(bannerPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(bannerPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(bannerPaginationQuery).Return(systemErr)

		bannerController := &bannerControllerImpl{
			bannerModelFactory: bannerModelFactory,
		}
		response, err := bannerController.ListBanners(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetBanner", func(t *testing.T) {
		bannerId := new(models.BannerId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		var banner *models.Banner
		bannerApi := mocks.NewMockBannerApi(ctrl)
		bannerApi.EXPECT().GetBanner(bannerId).Return(banner, nil)

		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerHttpHelper.EXPECT().ParseBannerId(httpContext).Return(bannerId, nil)

		bannerController := &bannerControllerImpl{bannerHttpHelper: bannerHttpHelper, bannerApi: bannerApi}
		response, err := bannerController.GetBanner(httpContext)

		assert.Equal(t, banner, response)
		assert.Nil(t, err)
	})

	t.Run("GetBanner:ParserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerHttpHelper.EXPECT().ParseBannerId(httpContext).Return(nil, systemErr)

		bannerController := &bannerControllerImpl{bannerHttpHelper: bannerHttpHelper}
		response, err := bannerController.GetBanner(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetBanner:ApiError", func(t *testing.T) {
		bannerId := new(models.BannerId)
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		bannerApi := mocks.NewMockBannerApi(ctrl)
		bannerApi.EXPECT().GetBanner(bannerId).Return(nil, systemErr)

		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerHttpHelper.EXPECT().ParseBannerId(httpContext).Return(bannerId, nil)

		bannerController := &bannerControllerImpl{bannerHttpHelper: bannerHttpHelper, bannerApi: bannerApi}
		response, err := bannerController.GetBanner(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateBanner", func(t *testing.T) {
		banner := new(models.Banner)
		data := new(models.BannerCreate)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBannerCreate().Return(data)

		bannerApi := mocks.NewMockBannerApi(ctrl)
		bannerApi.EXPECT().CreateBanner(data).Return(banner, nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		bannerController := &bannerControllerImpl{
			bannerModelFactory: bannerModelFactory,
			bannerApi:          bannerApi,
		}
		response, err := bannerController.CreateBanner(httpContext)

		assert.Equal(t, banner, response)
		assert.Nil(t, err)
	})

	t.Run("CreateBanner:BindBannerUpdateError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.BannerCreate)

		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBannerCreate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		bannerController := &bannerControllerImpl{
			bannerModelFactory: bannerModelFactory,
		}
		_, err := bannerController.CreateBanner(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateBanner:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.BannerCreate)

		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBannerCreate().Return(data)

		bannerApi := mocks.NewMockBannerApi(ctrl)
		bannerApi.EXPECT().CreateBanner(data).Return(nil, systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		bannerController := &bannerControllerImpl{
			bannerModelFactory: bannerModelFactory,
			bannerApi:          bannerApi,
		}
		_, err := bannerController.CreateBanner(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateBanner", func(t *testing.T) {
		bannerId := new(models.BannerId)
		data := new(models.BannerUpdate)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBannerUpdate().Return(data)

		bannerApi := mocks.NewMockBannerApi(ctrl)
		bannerApi.EXPECT().UpdateBanner(bannerId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerHttpHelper.EXPECT().ParseBannerId(httpContext).Return(bannerId, nil)

		bannerController := &bannerControllerImpl{
			bannerHttpHelper:   bannerHttpHelper,
			bannerModelFactory: bannerModelFactory,
			bannerApi:          bannerApi,
		}
		_, err := bannerController.UpdateBanner(httpContext)

		assert.Nil(t, err)
	})

	t.Run("UpdateBanner:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerHttpHelper.EXPECT().ParseBannerId(httpContext).Return(nil, systemErr)

		bannerController := &bannerControllerImpl{bannerHttpHelper: bannerHttpHelper}
		_, err := bannerController.UpdateBanner(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateBanner:BindBannerUpdateError", func(t *testing.T) {
		bannerId := new(models.BannerId)
		systemErr := common.NewUnknownError()
		data := new(models.BannerUpdate)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBannerUpdate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerHttpHelper.EXPECT().ParseBannerId(httpContext).Return(bannerId, nil)

		bannerController := &bannerControllerImpl{
			bannerHttpHelper:   bannerHttpHelper,
			bannerModelFactory: bannerModelFactory,
		}
		_, err := bannerController.UpdateBanner(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateBanner:ApiError", func(t *testing.T) {
		bannerId := new(models.BannerId)
		systemErr := common.NewUnknownError()

		data := new(models.BannerUpdate)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBannerUpdate().Return(data)

		bannerApi := mocks.NewMockBannerApi(ctrl)
		bannerApi.EXPECT().UpdateBanner(bannerId, data).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerHttpHelper.EXPECT().ParseBannerId(httpContext).Return(bannerId, nil)

		bannerController := &bannerControllerImpl{
			bannerHttpHelper:   bannerHttpHelper,
			bannerModelFactory: bannerModelFactory,
			bannerApi:          bannerApi,
		}
		_, err := bannerController.UpdateBanner(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteBanner", func(t *testing.T) {
		bannerId := new(models.BannerId)

		bannerApi := mocks.NewMockBannerApi(ctrl)
		bannerApi.EXPECT().DeleteBanner(bannerId).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)

		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerHttpHelper.EXPECT().ParseBannerId(httpContext).Return(bannerId, nil)

		bannerController := &bannerControllerImpl{bannerHttpHelper: bannerHttpHelper, bannerApi: bannerApi}
		_, err := bannerController.DeleteBanner(httpContext)

		assert.Nil(t, err)
	})

	t.Run("DeleteBanner:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		httpContext := mocks.NewMockHttpContext(ctrl)

		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerHttpHelper.EXPECT().ParseBannerId(httpContext).Return(nil, systemErr)

		bannerController := &bannerControllerImpl{bannerHttpHelper: bannerHttpHelper}
		_, err := bannerController.DeleteBanner(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteBanner:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		bannerId := new(models.BannerId)

		bannerApi := mocks.NewMockBannerApi(ctrl)
		bannerApi.EXPECT().DeleteBanner(bannerId).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)

		bannerHttpHelper := mocks.NewMockBannerHttpHelper(ctrl)
		bannerHttpHelper.EXPECT().ParseBannerId(httpContext).Return(bannerId, nil)

		bannerController := &bannerControllerImpl{bannerHttpHelper: bannerHttpHelper, bannerApi: bannerApi}
		_, err := bannerController.DeleteBanner(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
