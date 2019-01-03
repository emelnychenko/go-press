package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollEventFactory(t *testing.T) {
	t.Run("NewPollEventFactory", func(t *testing.T) {
		_, isPollEventFactory := NewPollEventFactory().(*pollEventFactoryImpl)

		assert.True(t, isPollEventFactory)
	})

	t.Run("CreatePollCreatedEvent", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		pollEventFactory := new(pollEventFactoryImpl)
		pollEvent := pollEventFactory.CreatePollCreatedEvent(pollEntity)

		assert.Equal(t, events.PollCreatedEventName, pollEvent.Name())
		assert.Equal(t, pollEntity, pollEvent.PollEntity())
	})

	t.Run("CreatePollUpdatedEvent", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		pollEventFactory := new(pollEventFactoryImpl)
		pollEvent := pollEventFactory.CreatePollUpdatedEvent(pollEntity)

		assert.Equal(t, events.PollUpdatedEventName, pollEvent.Name())
		assert.Equal(t, pollEntity, pollEvent.PollEntity())
	})

	t.Run("CreatePollDeletedEvent", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		pollEventFactory := new(pollEventFactoryImpl)
		pollEvent := pollEventFactory.CreatePollDeletedEvent(pollEntity)

		assert.Equal(t, events.PollDeletedEventName, pollEvent.Name())
		assert.Equal(t, pollEntity, pollEvent.PollEntity())
	})
}
