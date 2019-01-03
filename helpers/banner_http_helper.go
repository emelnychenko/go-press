package helpers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

const (
	BannerIdParameterName = "bannerId"
)

type (
	bannerHttpHelperImpl struct {
	}
)

func NewBannerHttpHelper() contracts.BannerHttpHelper {
	return new(bannerHttpHelperImpl)
}

func (*bannerHttpHelperImpl) ParseBannerId(httpContext contracts.HttpContext) (*models.BannerId, common.Error) {
	return common.ParseModelId(httpContext.Parameter(BannerIdParameterName))
}
