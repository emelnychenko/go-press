package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostAuthorEvents(t *testing.T) {
	t.Run("NewPostAuthorChangedEvent", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postAuthorEntity := new(entities.UserEntity)

		postAuthorEvent, isPostAuthorEvent := NewPostAuthorChangedEvent(postEntity, postAuthorEntity).(*PostAuthorEvent)

		assert.True(t, isPostAuthorEvent)
		assert.Equal(t, postEntity, postAuthorEvent.postEntity)
		assert.Equal(t, postAuthorEntity, postAuthorEvent.postAuthorEntity)
		assert.Equal(t, PostAuthorChangedEventName, postAuthorEvent.name)
	})

	t.Run("PostAuthorEntity", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		postAuthorEntity := new(entities.UserEntity)

		postAuthorEvent := &PostAuthorEvent{postEntity: postEntity, postAuthorEntity: postAuthorEntity}

		assert.Equal(t, postEntity, postAuthorEvent.PostEntity())
		assert.Equal(t, postAuthorEntity, postAuthorEvent.PostAuthorEntity())
	})
}
