package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	BannerAggregator interface {
		AggregateBanner(bannerEntity *entities.BannerEntity) *models.Banner
		AggregateBanners(bannerEntities []*entities.BannerEntity) []*models.Banner
		AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult
	}
)
