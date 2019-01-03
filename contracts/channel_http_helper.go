package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	ChannelHttpHelper interface {
		ParseChannelId(httpContext HttpContext) (*models.ChannelId, common.Error)
	}
)
