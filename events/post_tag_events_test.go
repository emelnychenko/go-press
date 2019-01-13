package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostTagEvents(t *testing.T) {
	t.Run("NewPostTagAddedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		tagEntity := new(entities.TagEntity)

		postTagEvent, isPostTagEvent := NewPostTagAddedEvent(postEntity, tagEntity).(*PostTagEvent)

		assert.True(t, isPostTagEvent)
		assert.Equal(t, postEntity, postTagEvent.postEntity)
		assert.Equal(t, tagEntity, postTagEvent.tagEntity)
		assert.Equal(t, PostTagAddedEventName, postTagEvent.name)
	})

	t.Run("NewPostTagRemovedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		tagEntity := new(entities.TagEntity)

		postTagEvent, isPostTagEvent := NewPostTagRemovedEvent(postEntity, tagEntity).(*PostTagEvent)

		assert.True(t, isPostTagEvent)
		assert.Equal(t, postEntity, postTagEvent.postEntity)
		assert.Equal(t, tagEntity, postTagEvent.tagEntity)
		assert.Equal(t, PostTagRemovedEventName, postTagEvent.name)
	})

	t.Run("PostTagEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		tagEntity := new(entities.TagEntity)

		postTagEvent := &PostTagEvent{postEntity: postEntity, tagEntity: tagEntity}

		assert.Equal(t, postEntity, postTagEvent.PostEntity())
		assert.Equal(t, tagEntity, postTagEvent.TagEntity())
	})
}
