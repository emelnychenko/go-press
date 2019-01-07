package helpers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

const (
	ChannelIdParameterName = "channelId"
)

type (
	channelHttpHelperImpl struct {
	}
)

func NewChannelHttpHelper() contracts.ChannelHttpHelper {
	return new(channelHttpHelperImpl)
}

func (*channelHttpHelperImpl) ParseChannelId(httpContext contracts.HttpContext) (*models.ChannelId, errors.Error) {
	return models.ParseModelId(httpContext.Parameter(ChannelIdParameterName))
}
