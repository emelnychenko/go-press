package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	BannerHttpHelper interface {
		ParseBannerId(httpContext HttpContext) (*models.BannerId, errors.Error)
	}
)
