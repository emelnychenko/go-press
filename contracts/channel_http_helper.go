package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	ChannelHttpHelper interface {
		ParseChannelId(httpContext HttpContext) (*models.ChannelId, errors.Error)
	}
)
