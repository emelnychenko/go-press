package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostEventFactory(t *testing.T) {
	t.Run("NewPostEventFactory", func(t *testing.T) {
		_, isPostEventFactory := NewPostEventFactory().(*postEventFactoryImpl)

		assert.True(t, isPostEventFactory)
	})

	t.Run("CreatePostCreatedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEventFactory := new(postEventFactoryImpl)
		postEvent := postEventFactory.CreatePostCreatedEvent(postEntity)

		assert.Equal(t, events.PostCreatedEventName, postEvent.Name())
		assert.Equal(t, postEntity, postEvent.PostEntity())
	})

	t.Run("CreatePostUpdatedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEventFactory := new(postEventFactoryImpl)
		postEvent := postEventFactory.CreatePostUpdatedEvent(postEntity)

		assert.Equal(t, events.PostUpdatedEventName, postEvent.Name())
		assert.Equal(t, postEntity, postEvent.PostEntity())
	})

	t.Run("CreatePostDeletedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEventFactory := new(postEventFactoryImpl)
		postEvent := postEventFactory.CreatePostDeletedEvent(postEntity)

		assert.Equal(t, events.PostDeletedEventName, postEvent.Name())
		assert.Equal(t, postEntity, postEvent.PostEntity())
	})

	t.Run("CreatePostPublishedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEventFactory := new(postEventFactoryImpl)
		postEvent := postEventFactory.CreatePostPublishedEvent(postEntity)

		assert.Equal(t, events.PostPublishedEventName, postEvent.Name())
		assert.Equal(t, postEntity, postEvent.PostEntity())
	})

	t.Run("CreatePostConcealedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postEventFactory := new(postEventFactoryImpl)
		postEvent := postEventFactory.CreatePostConcealedEvent(postEntity)

		assert.Equal(t, events.PostConcealedEventName, postEvent.Name())
		assert.Equal(t, postEntity, postEvent.PostEntity())
	})
}
