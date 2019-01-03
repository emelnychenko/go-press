package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	channelApiImpl struct {
		eventDispatcher     contracts.EventDispatcher
		channelEventFactory contracts.ChannelEventFactory
		channelService      contracts.ChannelService
		channelAggregator   contracts.ChannelAggregator
	}
)

func NewChannelApi(
	eventDispatcher contracts.EventDispatcher,
	channelEventFactory contracts.ChannelEventFactory,
	channelService contracts.ChannelService,
	channelAggregator contracts.ChannelAggregator,
) (channelApi contracts.ChannelApi) {
	return &channelApiImpl{
		eventDispatcher,
		channelEventFactory,
		channelService,
		channelAggregator,
	}
}

func (a *channelApiImpl) ListChannels(
	channelPaginationQuery *models.ChannelPaginationQuery,
) (paginationResult *models.PaginationResult, err common.Error) {
	entityPaginationResult, err := a.channelService.ListChannels(channelPaginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.channelAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}

func (a *channelApiImpl) GetChannel(channelId *models.ChannelId) (channel *models.Channel, err common.Error) {
	channelEntity, err := a.channelService.GetChannel(channelId)

	if nil != err {
		return
	}

	channel = a.channelAggregator.AggregateChannel(channelEntity)
	return
}

func (a *channelApiImpl) CreateChannel(data *models.ChannelCreate) (channel *models.Channel, err common.Error) {
	channelEntity, err := a.channelService.CreateChannel(data)

	if nil != err {
		return
	}

	channelCreatedEvent := a.channelEventFactory.CreateChannelCreatedEvent(channelEntity)
	a.eventDispatcher.Dispatch(channelCreatedEvent)

	channel = a.channelAggregator.AggregateChannel(channelEntity)
	return
}

func (a *channelApiImpl) UpdateChannel(channelId *models.ChannelId, data *models.ChannelUpdate) (err common.Error) {
	channelService := a.channelService
	channelEntity, err := channelService.GetChannel(channelId)

	if nil != err {
		return
	}

	err = channelService.UpdateChannel(channelEntity, data)

	if nil != err {
		return
	}

	channelUpdatedEvent := a.channelEventFactory.CreateChannelUpdatedEvent(channelEntity)
	a.eventDispatcher.Dispatch(channelUpdatedEvent)
	return
}

func (a *channelApiImpl) DeleteChannel(channelId *models.ChannelId) (err common.Error) {
	channelService := a.channelService
	channelEntity, err := channelService.GetChannel(channelId)

	if nil != err {
		return
	}

	err = channelService.DeleteChannel(channelEntity)

	if nil != err {
		return
	}

	channelDeletedEvent := a.channelEventFactory.CreateChannelDeletedEvent(channelEntity)
	a.eventDispatcher.Dispatch(channelDeletedEvent)

	return
}
