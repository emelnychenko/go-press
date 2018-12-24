package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostVideoEvents(t *testing.T) {
	t.Run("NewPostVideoChangedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postVideo := new(entities.FileEntity)

		postVideoEvent, isPostVideoEvent := NewPostVideoChangedEvent(postEntity, postVideo).(*PostVideoEvent)

		assert.True(t, isPostVideoEvent)
		assert.Equal(t, postEntity, postVideoEvent.postEntity)
		assert.Equal(t, postVideo, postVideoEvent.postVideo)
		assert.Equal(t, PostVideoChangedEventName, postVideoEvent.name)
	})

	t.Run("NewPostVideoRemovedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)

		postVideoEvent, isPostVideoEvent := NewPostVideoRemovedEvent(postEntity).(*PostVideoEvent)

		assert.True(t, isPostVideoEvent)
		assert.Equal(t, postEntity, postVideoEvent.postEntity)
		assert.Nil(t, postVideoEvent.postVideo)
		assert.Equal(t, PostVideoRemovedEventName, postVideoEvent.name)
	})

	t.Run("PostVideoEntity", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postVideo := new(entities.FileEntity)

		postVideoEvent := &PostVideoEvent{postEntity: postEntity, postVideo: postVideo}

		assert.Equal(t, postEntity, postVideoEvent.PostEntity())
		assert.Equal(t, postVideo, postVideoEvent.PostVideoEntity())
	})
}
