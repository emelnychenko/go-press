package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	BannerService interface {
		ListBanners(bannerPaginationQuery *models.BannerPaginationQuery) (*models.PaginationResult, common.Error)
		GetBanner(bannerId *models.BannerId) (bannerEntity *entities.BannerEntity, err common.Error)
		CreateBanner(data *models.BannerCreate) (bannerEntity *entities.BannerEntity, err common.Error)
		UpdateBanner(bannerEntity *entities.BannerEntity, data *models.BannerUpdate) (err common.Error)
		DeleteBanner(bannerEntity *entities.BannerEntity) (err common.Error)
	}
)
