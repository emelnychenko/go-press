package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	ChannelService interface {
		ListChannels(channelPaginationQuery *models.ChannelPaginationQuery) (*models.PaginationResult, common.Error)
		GetChannel(channelId *models.ChannelId) (channelEntity *entities.ChannelEntity, err common.Error)
		CreateChannel(data *models.ChannelCreate) (channelEntity *entities.ChannelEntity, err common.Error)
		UpdateChannel(channelEntity *entities.ChannelEntity, data *models.ChannelUpdate) (err common.Error)
		DeleteChannel(channelEntity *entities.ChannelEntity) (err common.Error)
	}
)
