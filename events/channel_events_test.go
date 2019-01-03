package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelEvents(t *testing.T) {
	t.Run("NewChannelCreatedEvent", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channelEvent, isChannelEvent := NewChannelCreatedEvent(channelEntity).(*ChannelEvent)

		assert.True(t, isChannelEvent)
		assert.Equal(t, channelEntity, channelEvent.channelEntity)
		assert.Equal(t, ChannelCreatedEventName, channelEvent.name)
	})

	t.Run("NewChannelUpdatedEvent", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channelEvent, isChannelEvent := NewChannelUpdatedEvent(channelEntity).(*ChannelEvent)

		assert.True(t, isChannelEvent)
		assert.Equal(t, channelEntity, channelEvent.channelEntity)
		assert.Equal(t, ChannelUpdatedEventName, channelEvent.name)
	})

	t.Run("NewChannelDeletedEvent", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channelEvent, isChannelEvent := NewChannelDeletedEvent(channelEntity).(*ChannelEvent)

		assert.True(t, isChannelEvent)
		assert.Equal(t, channelEntity, channelEvent.channelEntity)
		assert.Equal(t, ChannelDeletedEventName, channelEvent.name)
	})

	t.Run("ChannelEntity", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channelEvent := &ChannelEvent{channelEntity: channelEntity}

		assert.Equal(t, channelEntity, channelEvent.ChannelEntity())
	})
}
