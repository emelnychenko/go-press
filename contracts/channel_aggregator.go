package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	ChannelAggregator interface {
		AggregateChannel(channelEntity *entities.ChannelEntity) *models.Channel
		AggregateChannels(channelEntities []*entities.ChannelEntity) []*models.Channel
		AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult
	}
)
