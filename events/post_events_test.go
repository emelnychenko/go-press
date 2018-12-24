package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostEvents(t *testing.T) {
	t.Run("NewPostCreatedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEvent, isPostEvent := NewPostCreatedEvent(postEntity).(*PostEvent)

		assert.True(t, isPostEvent)
		assert.Equal(t, postEntity, postEvent.postEntity)
		assert.Equal(t, PostCreatedEventName, postEvent.name)
	})

	t.Run("NewPostUpdatedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEvent, isPostEvent := NewPostUpdatedEvent(postEntity).(*PostEvent)

		assert.True(t, isPostEvent)
		assert.Equal(t, postEntity, postEvent.postEntity)
		assert.Equal(t, PostUpdatedEventName, postEvent.name)
	})

	t.Run("NewPostDeletedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEvent, isPostEvent := NewPostDeletedEvent(postEntity).(*PostEvent)

		assert.True(t, isPostEvent)
		assert.Equal(t, postEntity, postEvent.postEntity)
		assert.Equal(t, PostDeletedEventName, postEvent.name)
	})

	t.Run("NewPostAuthorChangedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEvent, isPostEvent := NewPostAuthorChangedEvent(postEntity).(*PostEvent)

		assert.True(t, isPostEvent)
		assert.Equal(t, postEntity, postEvent.postEntity)
		assert.Equal(t, PostAuthorChangedEventName, postEvent.name)
	})

	t.Run("PostEntity", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEvent := &PostEvent{postEntity: postEntity}

		assert.Equal(t, postEntity, postEvent.PostEntity())
	})
}
