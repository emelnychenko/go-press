package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewChannelApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		channelEventFactory := mocks.NewMockChannelEventFactory(ctrl)
		channelService := mocks.NewMockChannelService(ctrl)
		channelAggregator := mocks.NewMockChannelAggregator(ctrl)

		channelApi, isChannelApi := NewChannelApi(
			eventDispatcher, channelEventFactory, channelService, channelAggregator,
		).(*channelApiImpl)

		assert.True(t, isChannelApi)
		assert.Equal(t, eventDispatcher, channelApi.eventDispatcher)
		assert.Equal(t, channelEventFactory, channelApi.channelEventFactory)
		assert.Equal(t, channelService, channelApi.channelService)
		assert.Equal(t, channelAggregator, channelApi.channelAggregator)
	})

	t.Run("ListChannels", func(t *testing.T) {
		paginationQuery := new(models.ChannelPaginationQuery)
		entityPaginationResult := new(models.PaginationResult)
		paginationResult := new(models.PaginationResult)

		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().ListChannels(paginationQuery).Return(entityPaginationResult, nil)

		channelAggregator := mocks.NewMockChannelAggregator(ctrl)
		channelAggregator.EXPECT().AggregatePaginationResult(entityPaginationResult).Return(paginationResult)

		channelApi := &channelApiImpl{channelService: channelService, channelAggregator: channelAggregator}
		response, err := channelApi.ListChannels(paginationQuery)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListChannels:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		paginationQuery := new(models.ChannelPaginationQuery)

		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().ListChannels(paginationQuery).Return(nil, systemErr)

		channelApi := &channelApiImpl{channelService: channelService}
		response, err := channelApi.ListChannels(paginationQuery)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetChannel", func(t *testing.T) {
		channelId := new(models.ChannelId)
		channelEntity := new(entities.ChannelEntity)
		channel := new(models.Channel)

		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().GetChannel(channelId).Return(channelEntity, nil)

		channelAggregator := mocks.NewMockChannelAggregator(ctrl)
		channelAggregator.EXPECT().AggregateChannel(channelEntity).Return(channel)

		channelApi := &channelApiImpl{channelService: channelService, channelAggregator: channelAggregator}
		response, err := channelApi.GetChannel(channelId)

		assert.Equal(t, channel, response)
		assert.Nil(t, err)
	})

	t.Run("GetChannel:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		channelId := new(models.ChannelId)
		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().GetChannel(channelId).Return(nil, systemErr)

		channelApi := &channelApiImpl{channelService: channelService}
		response, err := channelApi.GetChannel(channelId)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateChannel", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channel := new(models.Channel)
		data := new(models.ChannelCreate)

		channelEvent := new(events.ChannelEvent)
		channelEventFactory := mocks.NewMockChannelEventFactory(ctrl)
		channelEventFactory.EXPECT().CreateChannelCreatedEvent(channelEntity).Return(channelEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(channelEvent)

		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().CreateChannel(data).Return(channelEntity, nil)

		channelAggregator := mocks.NewMockChannelAggregator(ctrl)
		channelAggregator.EXPECT().AggregateChannel(channelEntity).Return(channel)

		channelApi := &channelApiImpl{
			eventDispatcher:     eventDispatcher,
			channelEventFactory: channelEventFactory,
			channelService:      channelService,
			channelAggregator:   channelAggregator,
		}
		response, err := channelApi.CreateChannel(data)

		assert.Equal(t, channel, response)
		assert.Nil(t, err)
	})

	t.Run("CreateChannel:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		data := new(models.ChannelCreate)
		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().CreateChannel(data).Return(nil, systemErr)

		channelApi := &channelApiImpl{channelService: channelService}
		response, err := channelApi.CreateChannel(data)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateChannel", func(t *testing.T) {
		channelId := new(models.ChannelId)
		channelEntity := new(entities.ChannelEntity)
		data := new(models.ChannelUpdate)

		channelEvent := new(events.ChannelEvent)
		channelEventFactory := mocks.NewMockChannelEventFactory(ctrl)
		channelEventFactory.EXPECT().CreateChannelUpdatedEvent(channelEntity).Return(channelEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(channelEvent)

		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().GetChannel(channelId).Return(channelEntity, nil)
		channelService.EXPECT().UpdateChannel(channelEntity, data).Return(nil)

		channelApi := &channelApiImpl{
			eventDispatcher:     eventDispatcher,
			channelEventFactory: channelEventFactory,
			channelService:      channelService,
		}
		assert.Nil(t, channelApi.UpdateChannel(channelId, data))
	})

	t.Run("UpdateChannel:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		channelId := new(models.ChannelId)
		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().GetChannel(channelId).Return(nil, systemErr)

		data := new(models.ChannelUpdate)
		channelApi := &channelApiImpl{channelService: channelService}
		assert.Equal(t, systemErr, channelApi.UpdateChannel(channelId, data))
	})

	t.Run("UpdateChannel:UpdateChannelError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		channelId := new(models.ChannelId)
		channelEntity := new(entities.ChannelEntity)
		data := new(models.ChannelUpdate)

		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().GetChannel(channelId).Return(channelEntity, nil)
		channelService.EXPECT().UpdateChannel(channelEntity, data).Return(systemErr)

		channelApi := &channelApiImpl{
			channelService: channelService,
		}

		err := channelApi.UpdateChannel(channelId, data)
		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteChannel", func(t *testing.T) {
		channelId := new(models.ChannelId)
		channelEntity := new(entities.ChannelEntity)

		channelEvent := new(events.ChannelEvent)
		channelEventFactory := mocks.NewMockChannelEventFactory(ctrl)
		channelEventFactory.EXPECT().CreateChannelDeletedEvent(channelEntity).Return(channelEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(channelEvent)

		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().GetChannel(channelId).Return(channelEntity, nil)
		channelService.EXPECT().DeleteChannel(channelEntity).Return(nil)

		channelApi := &channelApiImpl{
			eventDispatcher:     eventDispatcher,
			channelEventFactory: channelEventFactory,
			channelService:      channelService,
		}
		assert.Nil(t, channelApi.DeleteChannel(channelId))
	})

	t.Run("DeleteChannel:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		channelId := new(models.ChannelId)
		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().GetChannel(channelId).Return(nil, systemErr)

		channelApi := &channelApiImpl{channelService: channelService}
		assert.Equal(t, systemErr, channelApi.DeleteChannel(channelId))
	})

	t.Run("DeleteChannel:DeleteChannelError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		channelId := new(models.ChannelId)
		channelEntity := new(entities.ChannelEntity)

		channelService := mocks.NewMockChannelService(ctrl)
		channelService.EXPECT().GetChannel(channelId).Return(channelEntity, nil)
		channelService.EXPECT().DeleteChannel(channelEntity).Return(systemErr)

		channelApi := &channelApiImpl{
			channelService: channelService,
		}

		err := channelApi.DeleteChannel(channelId)
		assert.Equal(t, systemErr, err)
	})
}
