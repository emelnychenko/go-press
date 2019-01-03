package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagEventFactory(t *testing.T) {
	t.Run("NewTagEventFactory", func(t *testing.T) {
		_, isTagEventFactory := NewTagEventFactory().(*tagEventFactoryImpl)

		assert.True(t, isTagEventFactory)
	})

	t.Run("CreateTagCreatedEvent", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tagEventFactory := new(tagEventFactoryImpl)
		tagEvent := tagEventFactory.CreateTagCreatedEvent(tagEntity)

		assert.Equal(t, events.TagCreatedEventName, tagEvent.Name())
		assert.Equal(t, tagEntity, tagEvent.TagEntity())
	})

	t.Run("CreateTagUpdatedEvent", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tagEventFactory := new(tagEventFactoryImpl)
		tagEvent := tagEventFactory.CreateTagUpdatedEvent(tagEntity)

		assert.Equal(t, events.TagUpdatedEventName, tagEvent.Name())
		assert.Equal(t, tagEntity, tagEvent.TagEntity())
	})

	t.Run("CreateTagDeletedEvent", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tagEventFactory := new(tagEventFactoryImpl)
		tagEvent := tagEventFactory.CreateTagDeletedEvent(tagEntity)

		assert.Equal(t, events.TagDeletedEventName, tagEvent.Name())
		assert.Equal(t, tagEntity, tagEvent.TagEntity())
	})
}
