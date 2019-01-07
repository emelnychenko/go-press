package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	BannerService interface {
		ListBanners(bannerPaginationQuery *models.BannerPaginationQuery) (*models.PaginationResult, errors.Error)
		GetBanner(bannerId *models.BannerId) (bannerEntity *entities.BannerEntity, err errors.Error)
		CreateBanner(data *models.BannerCreate) (bannerEntity *entities.BannerEntity, err errors.Error)
		UpdateBanner(bannerEntity *entities.BannerEntity, data *models.BannerUpdate) (err errors.Error)
		DeleteBanner(bannerEntity *entities.BannerEntity) (err errors.Error)
	}
)
