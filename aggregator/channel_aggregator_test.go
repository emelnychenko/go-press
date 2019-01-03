package aggregator

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelAggregator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewChannelAggregator", func(t *testing.T) {
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelAggregator, isChannelAggregator := NewChannelAggregator(channelModelFactory).(*channelAggregatorImpl)

		assert.True(t, isChannelAggregator)
		assert.Equal(t, channelModelFactory, channelAggregator.channelModelFactory)
	})

	t.Run("AggregateChannel", func(t *testing.T) {
		channel := new(models.Channel)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannel().Return(channel)

		channelAggregator := &channelAggregatorImpl{channelModelFactory: channelModelFactory}
		response := channelAggregator.AggregateChannel(new(entities.ChannelEntity))

		assert.Equal(t, channel, response)
	})

	t.Run("AggregateChannels", func(t *testing.T) {
		channels := new(models.Channel)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannel().Return(channels)

		channelAggregator := &channelAggregatorImpl{channelModelFactory: channelModelFactory}
		channelEntities := []*entities.ChannelEntity{new(entities.ChannelEntity)}
		response := channelAggregator.AggregateChannels(channelEntities)

		assert.IsType(t, []*models.Channel{}, response)
		assert.Equal(t, len(channelEntities), len(response))
	})

	t.Run("AggregatePaginationResult", func(t *testing.T) {
		channel := new(models.Channel)
		channelModelFactory := mocks.NewMockChannelModelFactory(ctrl)
		channelModelFactory.EXPECT().CreateChannel().Return(channel)

		channelEntities := []*entities.ChannelEntity{entities.NewChannelEntity()}
		channelAggregator := &channelAggregatorImpl{channelModelFactory: channelModelFactory}

		entityPaginationResult := &models.PaginationResult{Data: channelEntities}
		paginationResult := channelAggregator.AggregatePaginationResult(entityPaginationResult)

		assert.IsType(t, []*models.Channel{}, paginationResult.Data)
		assert.Equal(t, len(channelEntities), len(paginationResult.Data.([]*models.Channel)))
	})
}
