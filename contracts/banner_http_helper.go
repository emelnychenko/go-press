package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	BannerHttpHelper interface {
		ParseBannerId(httpContext HttpContext) (*models.BannerId, common.Error)
	}
)
