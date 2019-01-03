package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	bannerModelFactoryImpl struct {
		paginationModelFactory contracts.PaginationModelFactory
	}
)

func NewBannerModelFactory(paginationModelFactory contracts.PaginationModelFactory) contracts.BannerModelFactory {
	return &bannerModelFactoryImpl{paginationModelFactory}
}

func (f *bannerModelFactoryImpl) CreateBannerPaginationQuery() *models.BannerPaginationQuery {
	paginationQuery := f.paginationModelFactory.CreatePaginationQuery()
	return &models.BannerPaginationQuery{PaginationQuery: paginationQuery}
}

func (*bannerModelFactoryImpl) CreateBanner() *models.Banner {
	return new(models.Banner)
}

func (*bannerModelFactoryImpl) CreateBannerCreate() *models.BannerCreate {
	return new(models.BannerCreate)
}

func (*bannerModelFactoryImpl) CreateBannerUpdate() *models.BannerUpdate {
	return new(models.BannerUpdate)
}
