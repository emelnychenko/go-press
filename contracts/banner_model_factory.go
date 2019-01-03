package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	BannerModelFactory interface {
		CreateBannerPaginationQuery() *models.BannerPaginationQuery
		CreateBanner() *models.Banner
		CreateBannerCreate() *models.BannerCreate
		CreateBannerUpdate() *models.BannerUpdate
	}
)
