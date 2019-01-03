package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	BannerRepository interface {
		ListBanners(bannerPaginationQuery *models.BannerPaginationQuery) (*models.PaginationResult, common.Error)
		GetBanner(bannerId *models.BannerId) (bannerEntity *entities.BannerEntity, err common.Error)
		SaveBanner(bannerEntity *entities.BannerEntity) (err common.Error)
		RemoveBanner(bannerEntity *entities.BannerEntity) (err common.Error)
	}
)
