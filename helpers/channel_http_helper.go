package helpers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
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

func (*channelHttpHelperImpl) ParseChannelId(httpContext contracts.HttpContext) (*models.ChannelId, common.Error) {
	return common.ParseModelId(httpContext.Parameter(ChannelIdParameterName))
}
