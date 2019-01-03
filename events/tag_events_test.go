package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagEvents(t *testing.T) {
	t.Run("NewTagCreatedEvent", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tagEvent, isTagEvent := NewTagCreatedEvent(tagEntity).(*TagEvent)

		assert.True(t, isTagEvent)
		assert.Equal(t, tagEntity, tagEvent.tagEntity)
		assert.Equal(t, TagCreatedEventName, tagEvent.name)
	})

	t.Run("NewTagUpdatedEvent", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tagEvent, isTagEvent := NewTagUpdatedEvent(tagEntity).(*TagEvent)

		assert.True(t, isTagEvent)
		assert.Equal(t, tagEntity, tagEvent.tagEntity)
		assert.Equal(t, TagUpdatedEventName, tagEvent.name)
	})

	t.Run("NewTagDeletedEvent", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tagEvent, isTagEvent := NewTagDeletedEvent(tagEntity).(*TagEvent)

		assert.True(t, isTagEvent)
		assert.Equal(t, tagEntity, tagEvent.tagEntity)
		assert.Equal(t, TagDeletedEventName, tagEvent.name)
	})

	t.Run("TagEntity", func(t *testing.T) {
		tagEntity := new(entities.TagEntity)
		tagEvent := &TagEvent{tagEntity: tagEntity}

		assert.Equal(t, tagEntity, tagEvent.TagEntity())
	})
}
