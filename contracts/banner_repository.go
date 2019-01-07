package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	BannerRepository interface {
		ListBanners(bannerPaginationQuery *models.BannerPaginationQuery) (*models.PaginationResult, errors.Error)
		GetBanner(bannerId *models.BannerId) (bannerEntity *entities.BannerEntity, err errors.Error)
		SaveBanner(bannerEntity *entities.BannerEntity) (err errors.Error)
		RemoveBanner(bannerEntity *entities.BannerEntity) (err errors.Error)
	}
)
