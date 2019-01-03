package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type channelAggregatorImpl struct {
	channelModelFactory contracts.ChannelModelFactory
}

func NewChannelAggregator(channelModelFactory contracts.ChannelModelFactory) contracts.ChannelAggregator {
	return &channelAggregatorImpl{channelModelFactory}
}

func (a *channelAggregatorImpl) AggregateChannel(channelEntity *entities.ChannelEntity) (channel *models.Channel) {
	channel = a.channelModelFactory.CreateChannel()
	channel.Id = channelEntity.Id
	channel.Name = channelEntity.Name
	channel.Created = channelEntity.Created

	return
}

func (a *channelAggregatorImpl) AggregateChannels(channelEntities []*entities.ChannelEntity) (channels []*models.Channel) {
	channels = make([]*models.Channel, len(channelEntities))

	for k, postEntity := range channelEntities {
		channels[k] = a.AggregateChannel(postEntity)
	}

	return
}

func (a *channelAggregatorImpl) AggregatePaginationResult(
	entityPaginationResult *models.PaginationResult,
) (
	paginationResult *models.PaginationResult,
) {
	channelEntities := entityPaginationResult.Data.([]*entities.ChannelEntity)
	channels := a.AggregateChannels(channelEntities)
	return &models.PaginationResult{Total: entityPaginationResult.Total, Data: channels}
}
