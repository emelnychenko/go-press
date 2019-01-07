package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	ChannelService interface {
		ListChannels(channelPaginationQuery *models.ChannelPaginationQuery) (*models.PaginationResult, errors.Error)
		GetChannel(channelId *models.ChannelId) (channelEntity *entities.ChannelEntity, err errors.Error)
		CreateChannel(data *models.ChannelCreate) (channelEntity *entities.ChannelEntity, err errors.Error)
		UpdateChannel(channelEntity *entities.ChannelEntity, data *models.ChannelUpdate) (err errors.Error)
		DeleteChannel(channelEntity *entities.ChannelEntity) (err errors.Error)
	}
)
