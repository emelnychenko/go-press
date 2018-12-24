package events

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvent(t *testing.T) {
	t.Run("NewEvent", func(t *testing.T) {
		eventName := "Test"
		event, isEvent := NewEvent(eventName).(*Event)

		assert.True(t, isEvent)
		assert.Equal(t, eventName, event.name)
	})

	t.Run("Name", func(t *testing.T) {
		eventName := "Test"
		fileEvent := &Event{name: eventName}

		assert.Equal(t, eventName, fileEvent.Name())
	})
}
