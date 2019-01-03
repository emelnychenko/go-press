package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	BannerApi interface {
		ListBanners(bannerPaginationQuery *models.BannerPaginationQuery) (*models.PaginationResult, common.Error)
		GetBanner(bannerId *models.BannerId) (banner *models.Banner, err common.Error)
		CreateBanner(data *models.BannerCreate) (banner *models.Banner, err common.Error)
		UpdateBanner(bannerId *models.BannerId, data *models.BannerUpdate) (err common.Error)
		DeleteBanner(bannerId *models.BannerId) (err common.Error)
	}
)
