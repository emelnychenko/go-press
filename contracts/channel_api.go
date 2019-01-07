package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	ChannelApi interface {
		ListChannels(channelPaginationQuery *models.ChannelPaginationQuery) (*models.PaginationResult, errors.Error)
		GetChannel(channelId *models.ChannelId) (channel *models.Channel, err errors.Error)
		CreateChannel(data *models.ChannelCreate) (channel *models.Channel, err errors.Error)
		UpdateChannel(channelId *models.ChannelId, data *models.ChannelUpdate) (err errors.Error)
		DeleteChannel(channelId *models.ChannelId) (err errors.Error)
	}
)
