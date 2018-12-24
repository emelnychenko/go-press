package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostVideoEventFactory(t *testing.T) {
	t.Run("NewPostVideoEventFactory", func(t *testing.T) {
		_, isPostVideoEventFactory := NewPostVideoEventFactory().(*postVideoEventFactoryImpl)

		assert.True(t, isPostVideoEventFactory)
	})

	t.Run("CreatePostVideoChangedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postVideoEntity := new(entities.FileEntity)

		postVideoEventFactory := new(postVideoEventFactoryImpl)
		postVideoEvent := postVideoEventFactory.CreatePostVideoChangedEvent(postEntity, postVideoEntity)

		assert.Equal(t, events.PostVideoChangedEventName, postVideoEvent.Name())
		assert.Equal(t, postEntity, postVideoEvent.PostEntity())
		assert.Equal(t, postVideoEntity, postVideoEvent.PostVideoEntity())
	})

	t.Run("CreatePostCreatedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postVideoEventFactory := new(postVideoEventFactoryImpl)
		postVideoEvent := postVideoEventFactory.CreatePostVideoRemovedEvent(postEntity)

		assert.Equal(t, events.PostVideoRemovedEventName, postVideoEvent.Name())
		assert.Equal(t, postEntity, postVideoEvent.PostEntity())
	})
}
