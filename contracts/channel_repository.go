package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	ChannelRepository interface {
		ListChannels(channelPaginationQuery *models.ChannelPaginationQuery) (*models.PaginationResult, common.Error)
		GetChannel(channelId *models.ChannelId) (channelEntity *entities.ChannelEntity, err common.Error)
		SaveChannel(channelEntity *entities.ChannelEntity) (err common.Error)
		RemoveChannel(channelEntity *entities.ChannelEntity) (err common.Error)
	}
)
