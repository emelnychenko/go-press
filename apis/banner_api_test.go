package apis

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewBannerApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		bannerEventFactory := mocks.NewMockBannerEventFactory(ctrl)
		bannerService := mocks.NewMockBannerService(ctrl)
		bannerAggregator := mocks.NewMockBannerAggregator(ctrl)

		bannerApi, isBannerApi := NewBannerApi(
			eventDispatcher, bannerEventFactory, bannerService, bannerAggregator,
		).(*bannerApiImpl)

		assert.True(t, isBannerApi)
		assert.Equal(t, eventDispatcher, bannerApi.eventDispatcher)
		assert.Equal(t, bannerEventFactory, bannerApi.bannerEventFactory)
		assert.Equal(t, bannerService, bannerApi.bannerService)
		assert.Equal(t, bannerAggregator, bannerApi.bannerAggregator)
	})

	t.Run("ListBanners", func(t *testing.T) {
		paginationQuery := new(models.BannerPaginationQuery)
		entityPaginationResult := new(models.PaginationResult)
		paginationResult := new(models.PaginationResult)

		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().ListBanners(paginationQuery).Return(entityPaginationResult, nil)

		bannerAggregator := mocks.NewMockBannerAggregator(ctrl)
		bannerAggregator.EXPECT().AggregatePaginationResult(entityPaginationResult).Return(paginationResult)

		bannerApi := &bannerApiImpl{bannerService: bannerService, bannerAggregator: bannerAggregator}
		response, err := bannerApi.ListBanners(paginationQuery)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListBanners:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		paginationQuery := new(models.BannerPaginationQuery)

		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().ListBanners(paginationQuery).Return(nil, systemErr)

		bannerApi := &bannerApiImpl{bannerService: bannerService}
		response, err := bannerApi.ListBanners(paginationQuery)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetBanner", func(t *testing.T) {
		bannerId := new(models.BannerId)
		bannerEntity := new(entities.BannerEntity)
		banner := new(models.Banner)

		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().GetBanner(bannerId).Return(bannerEntity, nil)

		bannerAggregator := mocks.NewMockBannerAggregator(ctrl)
		bannerAggregator.EXPECT().AggregateBanner(bannerEntity).Return(banner)

		bannerApi := &bannerApiImpl{bannerService: bannerService, bannerAggregator: bannerAggregator}
		response, err := bannerApi.GetBanner(bannerId)

		assert.Equal(t, banner, response)
		assert.Nil(t, err)
	})

	t.Run("GetBanner:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		bannerId := new(models.BannerId)
		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().GetBanner(bannerId).Return(nil, systemErr)

		bannerApi := &bannerApiImpl{bannerService: bannerService}
		response, err := bannerApi.GetBanner(bannerId)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateBanner", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		banner := new(models.Banner)
		data := new(models.BannerCreate)

		bannerEvent := new(events.BannerEvent)
		bannerEventFactory := mocks.NewMockBannerEventFactory(ctrl)
		bannerEventFactory.EXPECT().CreateBannerCreatedEvent(bannerEntity).Return(bannerEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(bannerEvent)

		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().CreateBanner(data).Return(bannerEntity, nil)

		bannerAggregator := mocks.NewMockBannerAggregator(ctrl)
		bannerAggregator.EXPECT().AggregateBanner(bannerEntity).Return(banner)

		bannerApi := &bannerApiImpl{
			eventDispatcher:    eventDispatcher,
			bannerEventFactory: bannerEventFactory,
			bannerService:      bannerService,
			bannerAggregator:   bannerAggregator,
		}
		response, err := bannerApi.CreateBanner(data)

		assert.Equal(t, banner, response)
		assert.Nil(t, err)
	})

	t.Run("CreateBanner:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		data := new(models.BannerCreate)
		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().CreateBanner(data).Return(nil, systemErr)

		bannerApi := &bannerApiImpl{bannerService: bannerService}
		response, err := bannerApi.CreateBanner(data)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateBanner", func(t *testing.T) {
		bannerId := new(models.BannerId)
		bannerEntity := new(entities.BannerEntity)
		data := new(models.BannerUpdate)

		bannerEvent := new(events.BannerEvent)
		bannerEventFactory := mocks.NewMockBannerEventFactory(ctrl)
		bannerEventFactory.EXPECT().CreateBannerUpdatedEvent(bannerEntity).Return(bannerEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(bannerEvent)

		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().GetBanner(bannerId).Return(bannerEntity, nil)
		bannerService.EXPECT().UpdateBanner(bannerEntity, data).Return(nil)

		bannerApi := &bannerApiImpl{
			eventDispatcher:    eventDispatcher,
			bannerEventFactory: bannerEventFactory,
			bannerService:      bannerService,
		}
		assert.Nil(t, bannerApi.UpdateBanner(bannerId, data))
	})

	t.Run("UpdateBanner:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		bannerId := new(models.BannerId)
		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().GetBanner(bannerId).Return(nil, systemErr)

		data := new(models.BannerUpdate)
		bannerApi := &bannerApiImpl{bannerService: bannerService}
		assert.Equal(t, systemErr, bannerApi.UpdateBanner(bannerId, data))
	})

	t.Run("UpdateBanner:UpdateBannerError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		bannerId := new(models.BannerId)
		bannerEntity := new(entities.BannerEntity)
		data := new(models.BannerUpdate)

		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().GetBanner(bannerId).Return(bannerEntity, nil)
		bannerService.EXPECT().UpdateBanner(bannerEntity, data).Return(systemErr)

		bannerApi := &bannerApiImpl{
			bannerService: bannerService,
		}

		err := bannerApi.UpdateBanner(bannerId, data)
		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteBanner", func(t *testing.T) {
		bannerId := new(models.BannerId)
		bannerEntity := new(entities.BannerEntity)

		bannerEvent := new(events.BannerEvent)
		bannerEventFactory := mocks.NewMockBannerEventFactory(ctrl)
		bannerEventFactory.EXPECT().CreateBannerDeletedEvent(bannerEntity).Return(bannerEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(bannerEvent)

		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().GetBanner(bannerId).Return(bannerEntity, nil)
		bannerService.EXPECT().DeleteBanner(bannerEntity).Return(nil)

		bannerApi := &bannerApiImpl{
			eventDispatcher:    eventDispatcher,
			bannerEventFactory: bannerEventFactory,
			bannerService:      bannerService,
		}
		assert.Nil(t, bannerApi.DeleteBanner(bannerId))
	})

	t.Run("DeleteBanner:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		bannerId := new(models.BannerId)
		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().GetBanner(bannerId).Return(nil, systemErr)

		bannerApi := &bannerApiImpl{bannerService: bannerService}
		assert.Equal(t, systemErr, bannerApi.DeleteBanner(bannerId))
	})

	t.Run("DeleteBanner:DeleteBannerError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		bannerId := new(models.BannerId)
		bannerEntity := new(entities.BannerEntity)

		bannerService := mocks.NewMockBannerService(ctrl)
		bannerService.EXPECT().GetBanner(bannerId).Return(bannerEntity, nil)
		bannerService.EXPECT().DeleteBanner(bannerEntity).Return(systemErr)

		bannerApi := &bannerApiImpl{
			bannerService: bannerService,
		}

		err := bannerApi.DeleteBanner(bannerId)
		assert.Equal(t, systemErr, err)
	})
}
