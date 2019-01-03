package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type bannerAggregatorImpl struct {
	bannerModelFactory contracts.BannerModelFactory
}

func NewBannerAggregator(bannerModelFactory contracts.BannerModelFactory) contracts.BannerAggregator {
	return &bannerAggregatorImpl{bannerModelFactory}
}

func (a *bannerAggregatorImpl) AggregateBanner(bannerEntity *entities.BannerEntity) (banner *models.Banner) {
	banner = a.bannerModelFactory.CreateBanner()
	banner.Id = bannerEntity.Id
	banner.Title = bannerEntity.Title
	banner.Key = bannerEntity.Key
	banner.Created = bannerEntity.Created

	return
}

func (a *bannerAggregatorImpl) AggregateBanners(bannerEntities []*entities.BannerEntity) (banners []*models.Banner) {
	banners = make([]*models.Banner, len(bannerEntities))

	for k, postEntity := range bannerEntities {
		banners[k] = a.AggregateBanner(postEntity)
	}

	return
}

func (a *bannerAggregatorImpl) AggregatePaginationResult(
	entityPaginationResult *models.PaginationResult,
) (
	paginationResult *models.PaginationResult,
) {
	bannerEntities := entityPaginationResult.Data.([]*entities.BannerEntity)
	banners := a.AggregateBanners(bannerEntities)
	return &models.PaginationResult{Total: entityPaginationResult.Total, Data: banners}
}