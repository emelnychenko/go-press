package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollEvents(t *testing.T) {
	t.Run("NewPollCreatedEvent", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		pollEvent, isPollEvent := NewPollCreatedEvent(pollEntity).(*PollEvent)

		assert.True(t, isPollEvent)
		assert.Equal(t, pollEntity, pollEvent.pollEntity)
		assert.Equal(t, PollCreatedEventName, pollEvent.name)
	})

	t.Run("NewPollUpdatedEvent", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		pollEvent, isPollEvent := NewPollUpdatedEvent(pollEntity).(*PollEvent)

		assert.True(t, isPollEvent)
		assert.Equal(t, pollEntity, pollEvent.pollEntity)
		assert.Equal(t, PollUpdatedEventName, pollEvent.name)
	})

	t.Run("NewPollDeletedEvent", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		pollEvent, isPollEvent := NewPollDeletedEvent(pollEntity).(*PollEvent)

		assert.True(t, isPollEvent)
		assert.Equal(t, pollEntity, pollEvent.pollEntity)
		assert.Equal(t, PollDeletedEventName, pollEvent.name)
	})

	t.Run("PollEntity", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		pollEvent := &PollEvent{pollEntity: pollEntity}

		assert.Equal(t, pollEntity, pollEvent.PollEntity())
	})
}
