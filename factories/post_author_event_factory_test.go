package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostAuthorEventFactory(t *testing.T) {
	t.Run("NewPostAuthorEventFactory", func(t *testing.T) {
		_, isPostAuthorEventFactory := NewPostAuthorEventFactory().(*postAuthorEventFactoryImpl)

		assert.True(t, isPostAuthorEventFactory)
	})

	t.Run("CreatePostAuthorChangedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postAuthorEntity := new(entities.UserEntity)

		postAuthorEventFactory := new(postAuthorEventFactoryImpl)
		postAuthorEvent := postAuthorEventFactory.CreatePostAuthorChangedEvent(postEntity, postAuthorEntity)

		assert.Equal(t, events.PostAuthorChangedEventName, postAuthorEvent.Name())
		assert.Equal(t, postEntity, postAuthorEvent.PostEntity())
		assert.Equal(t, postAuthorEntity, postAuthorEvent.PostAuthorEntity())
	})
}
