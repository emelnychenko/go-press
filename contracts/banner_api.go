package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	BannerApi interface {
		ListBanners(bannerPaginationQuery *models.BannerPaginationQuery) (*models.PaginationResult, errors.Error)
		GetBanner(bannerId *models.BannerId) (banner *models.Banner, err errors.Error)
		CreateBanner(data *models.BannerCreate) (banner *models.Banner, err errors.Error)
		UpdateBanner(bannerId *models.BannerId, data *models.BannerUpdate) (err errors.Error)
		DeleteBanner(bannerId *models.BannerId) (err errors.Error)
	}
)
