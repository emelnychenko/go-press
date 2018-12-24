package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostPictureEventFactory(t *testing.T) {
	t.Run("NewPostPictureEventFactory", func(t *testing.T) {
		_, isPostPictureEventFactory := NewPostPictureEventFactory().(*postPictureEventFactoryImpl)

		assert.True(t, isPostPictureEventFactory)
	})

	t.Run("CreatePostPictureChangedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postPictureEntity := new(entities.FileEntity)

		postPictureEventFactory := new(postPictureEventFactoryImpl)
		postPictureEvent := postPictureEventFactory.CreatePostPictureChangedEvent(postEntity, postPictureEntity)

		assert.Equal(t, events.PostPictureChangedEventName, postPictureEvent.Name())
		assert.Equal(t, postEntity, postPictureEvent.PostEntity())
		assert.Equal(t, postPictureEntity, postPictureEvent.PostPictureEntity())
	})

	t.Run("CreatePostCreatedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postPictureEventFactory := new(postPictureEventFactoryImpl)
		postPictureEvent := postPictureEventFactory.CreatePostPictureRemovedEvent(postEntity)

		assert.Equal(t, events.PostPictureRemovedEventName, postPictureEvent.Name())
		assert.Equal(t, postEntity, postPictureEvent.PostEntity())
	})
}
