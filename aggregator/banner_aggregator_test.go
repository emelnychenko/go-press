package aggregator

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerAggregator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewBannerAggregator", func(t *testing.T) {
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerAggregator, isBannerAggregator := NewBannerAggregator(bannerModelFactory).(*bannerAggregatorImpl)

		assert.True(t, isBannerAggregator)
		assert.Equal(t, bannerModelFactory, bannerAggregator.bannerModelFactory)
	})

	t.Run("AggregateBanner", func(t *testing.T) {
		banner := new(models.Banner)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBanner().Return(banner)

		bannerAggregator := &bannerAggregatorImpl{bannerModelFactory: bannerModelFactory}
		response := bannerAggregator.AggregateBanner(new(entities.BannerEntity))

		assert.Equal(t, banner, response)
	})

	t.Run("AggregateBanners", func(t *testing.T) {
		banners := new(models.Banner)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBanner().Return(banners)

		bannerAggregator := &bannerAggregatorImpl{bannerModelFactory: bannerModelFactory}
		bannerEntities := []*entities.BannerEntity{new(entities.BannerEntity)}
		response := bannerAggregator.AggregateBanners(bannerEntities)

		assert.IsType(t, []*models.Banner{}, response)
		assert.Equal(t, len(bannerEntities), len(response))
	})

	t.Run("AggregatePaginationResult", func(t *testing.T) {
		banner := new(models.Banner)
		bannerModelFactory := mocks.NewMockBannerModelFactory(ctrl)
		bannerModelFactory.EXPECT().CreateBanner().Return(banner)

		bannerEntities := []*entities.BannerEntity{entities.NewBannerEntity()}
		bannerAggregator := &bannerAggregatorImpl{bannerModelFactory: bannerModelFactory}

		entityPaginationResult := &models.PaginationResult{Data: bannerEntities}
		paginationResult := bannerAggregator.AggregatePaginationResult(entityPaginationResult)

		assert.IsType(t, []*models.Banner{}, paginationResult.Data)
		assert.Equal(t, len(bannerEntities), len(paginationResult.Data.([]*models.Banner)))
	})
}
