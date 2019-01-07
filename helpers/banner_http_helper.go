package helpers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
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

func (*bannerHttpHelperImpl) ParseBannerId(httpContext contracts.HttpContext) (*models.BannerId, errors.Error) {
	return models.ParseModelId(httpContext.Parameter(BannerIdParameterName))
}
