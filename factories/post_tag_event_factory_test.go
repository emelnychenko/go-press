package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostTagEventFactory(t *testing.T) {
	t.Run("NewPostTagEventFactory", func(t *testing.T) {
		_, isPostTagEventFactory := NewPostTagEventFactory().(*postTagEventFactoryImpl)

		assert.True(t, isPostTagEventFactory)
	})

	t.Run("CreatePostTagAddedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		tagEntity := new(entities.TagEntity)

		postTagEventFactory := new(postTagEventFactoryImpl)
		postTagEvent := postTagEventFactory.CreatePostTagAddedEvent(postEntity, tagEntity)

		assert.Equal(t, events.PostTagAddedEventName, postTagEvent.Name())
		assert.Equal(t, postEntity, postTagEvent.PostEntity())
		assert.Equal(t, tagEntity, postTagEvent.TagEntity())
	})

	t.Run("CreatePostTagRemovedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		tagEntity := new(entities.TagEntity)

		postTagEventFactory := new(postTagEventFactoryImpl)
		postTagEvent := postTagEventFactory.CreatePostTagRemovedEvent(postEntity, tagEntity)

		assert.Equal(t, events.PostTagRemovedEventName, postTagEvent.Name())
		assert.Equal(t, postEntity, postTagEvent.PostEntity())
		assert.Equal(t, tagEntity, postTagEvent.TagEntity())
	})
}
