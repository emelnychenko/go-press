package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	ChannelRepository interface {
		ListChannels(channelPaginationQuery *models.ChannelPaginationQuery) (*models.PaginationResult, errors.Error)
		GetChannel(channelId *models.ChannelId) (channelEntity *entities.ChannelEntity, err errors.Error)
		SaveChannel(channelEntity *entities.ChannelEntity) (err errors.Error)
		RemoveChannel(channelEntity *entities.ChannelEntity) (err errors.Error)
	}
)
