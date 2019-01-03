package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	ChannelApi interface {
		ListChannels(channelPaginationQuery *models.ChannelPaginationQuery) (*models.PaginationResult, common.Error)
		GetChannel(channelId *models.ChannelId) (channel *models.Channel, err common.Error)
		CreateChannel(data *models.ChannelCreate) (channel *models.Channel, err common.Error)
		UpdateChannel(channelId *models.ChannelId, data *models.ChannelUpdate) (err common.Error)
		DeleteChannel(channelId *models.ChannelId) (err common.Error)
	}
)
