package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewBannerService", func(t *testing.T) {
		bannerEntityFactory := mocks.NewMockBannerEntityFactory(ctrl)
		bannerRepository := mocks.NewMockBannerRepository(ctrl)

		bannerService, isBannerService := NewBannerService(bannerEntityFactory, bannerRepository).(*bannerServiceImpl)

		assert.True(t, isBannerService)
		assert.Equal(t, bannerEntityFactory, bannerService.bannerEntityFactory)
		assert.Equal(t, bannerRepository, bannerService.bannerRepository)
	})

	t.Run("ListBanners", func(t *testing.T) {
		bannerPaginationQuery := new(models.BannerPaginationQuery)

		var bannerEntities *models.PaginationResult
		bannerRepository := mocks.NewMockBannerRepository(ctrl)
		bannerRepository.EXPECT().ListBanners(bannerPaginationQuery).Return(bannerEntities, nil)

		bannerService := &bannerServiceImpl{bannerRepository: bannerRepository}
		response, err := bannerService.ListBanners(bannerPaginationQuery)

		assert.Equal(t, bannerEntities, response)
		assert.Nil(t, err)
	})

	t.Run("CreateBanner", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		bannerEntityFactory := mocks.NewMockBannerEntityFactory(ctrl)
		bannerEntityFactory.EXPECT().CreateBannerEntity().Return(bannerEntity)

		bannerRepository := mocks.NewMockBannerRepository(ctrl)
		bannerRepository.EXPECT().SaveBanner(bannerEntity).Return(nil)

		data := &models.BannerCreate{
			Title: "0",
			Key:   "1",
		}
		bannerService := &bannerServiceImpl{
			bannerEntityFactory: bannerEntityFactory,
			bannerRepository:    bannerRepository,
		}
		response, err := bannerService.CreateBanner(data)

		assert.IsType(t, bannerEntity, response)
		assert.Nil(t, err)
		assert.Equal(t, data.Title, bannerEntity.Title)
		assert.Equal(t, data.Key, bannerEntity.Key)
	})

	t.Run("GetBanner", func(t *testing.T) {
		bannerId := new(models.BannerId)
		bannerEntity := new(entities.BannerEntity)
		bannerRepository := mocks.NewMockBannerRepository(ctrl)
		bannerRepository.EXPECT().GetBanner(bannerId).Return(bannerEntity, nil)

		bannerService := &bannerServiceImpl{bannerRepository: bannerRepository}
		response, err := bannerService.GetBanner(bannerId)

		assert.Equal(t, bannerEntity, response)
		assert.Nil(t, err)
	})

	t.Run("UpdateBanner", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		bannerRepository := mocks.NewMockBannerRepository(ctrl)
		bannerRepository.EXPECT().SaveBanner(bannerEntity).Return(nil)

		data := &models.BannerUpdate{
			Title: "0",
			Key:   "1",
		}
		bannerService := &bannerServiceImpl{bannerRepository: bannerRepository}
		assert.Nil(t, bannerService.UpdateBanner(bannerEntity, data))

		assert.Equal(t, data.Title, bannerEntity.Title)
		assert.Equal(t, data.Key, bannerEntity.Key)
		assert.NotNil(t, bannerEntity.Updated)
	})

	t.Run("DeleteBanner", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		bannerRepository := mocks.NewMockBannerRepository(ctrl)
		bannerRepository.EXPECT().RemoveBanner(bannerEntity).Return(nil)

		bannerService := &bannerServiceImpl{bannerRepository: bannerRepository}
		assert.Nil(t, bannerService.DeleteBanner(bannerEntity))
	})
}
