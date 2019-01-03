package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelEventFactory(t *testing.T) {
	t.Run("NewChannelEventFactory", func(t *testing.T) {
		_, isChannelEventFactory := NewChannelEventFactory().(*channelEventFactoryImpl)

		assert.True(t, isChannelEventFactory)
	})

	t.Run("CreateChannelCreatedEvent", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channelEventFactory := new(channelEventFactoryImpl)
		channelEvent := channelEventFactory.CreateChannelCreatedEvent(channelEntity)

		assert.Equal(t, events.ChannelCreatedEventName, channelEvent.Name())
		assert.Equal(t, channelEntity, channelEvent.ChannelEntity())
	})

	t.Run("CreateChannelUpdatedEvent", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channelEventFactory := new(channelEventFactoryImpl)
		channelEvent := channelEventFactory.CreateChannelUpdatedEvent(channelEntity)

		assert.Equal(t, events.ChannelUpdatedEventName, channelEvent.Name())
		assert.Equal(t, channelEntity, channelEvent.ChannelEntity())
	})

	t.Run("CreateChannelDeletedEvent", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channelEventFactory := new(channelEventFactoryImpl)
		channelEvent := channelEventFactory.CreateChannelDeletedEvent(channelEntity)

		assert.Equal(t, events.ChannelDeletedEventName, channelEvent.Name())
		assert.Equal(t, channelEntity, channelEvent.ChannelEntity())
	})
}
