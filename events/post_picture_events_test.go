package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostPictureEvents(t *testing.T) {
	t.Run("NewPostPictureChangedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postPicture := new(entities.FileEntity)

		postPictureEvent, isPostPictureEvent := NewPostPictureChangedEvent(postEntity, postPicture).(*PostPictureEvent)

		assert.True(t, isPostPictureEvent)
		assert.Equal(t, postEntity, postPictureEvent.postEntity)
		assert.Equal(t, postPicture, postPictureEvent.postPicture)
		assert.Equal(t, PostPictureChangedEventName, postPictureEvent.name)
	})

	t.Run("NewPostPictureRemovedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)

		postPictureEvent, isPostPictureEvent := NewPostPictureRemovedEvent(postEntity).(*PostPictureEvent)

		assert.True(t, isPostPictureEvent)
		assert.Equal(t, postEntity, postPictureEvent.postEntity)
		assert.Nil(t, postPictureEvent.postPicture)
		assert.Equal(t, PostPictureRemovedEventName, postPictureEvent.name)
	})

	t.Run("PostPictureEntity", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postPicture := new(entities.FileEntity)

		postPictureEvent := &PostPictureEvent{postEntity: postEntity, postPicture: postPicture}

		assert.Equal(t, postEntity, postPictureEvent.PostEntity())
		assert.Equal(t, postPicture, postPictureEvent.PostPictureEntity())
	})
}
