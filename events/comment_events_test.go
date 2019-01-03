package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentEvents(t *testing.T) {
	t.Run("NewCommentCreatedEvent", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		commentEvent, isCommentEvent := NewCommentCreatedEvent(commentEntity).(*CommentEvent)

		assert.True(t, isCommentEvent)
		assert.Equal(t, commentEntity, commentEvent.commentEntity)
		assert.Equal(t, CommentCreatedEventName, commentEvent.name)
	})

	t.Run("NewCommentUpdatedEvent", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		commentEvent, isCommentEvent := NewCommentUpdatedEvent(commentEntity).(*CommentEvent)

		assert.True(t, isCommentEvent)
		assert.Equal(t, commentEntity, commentEvent.commentEntity)
		assert.Equal(t, CommentUpdatedEventName, commentEvent.name)
	})

	t.Run("NewCommentDeletedEvent", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		commentEvent, isCommentEvent := NewCommentDeletedEvent(commentEntity).(*CommentEvent)

		assert.True(t, isCommentEvent)
		assert.Equal(t, commentEntity, commentEvent.commentEntity)
		assert.Equal(t, CommentDeletedEventName, commentEvent.name)
	})

	t.Run("CommentEntity", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		commentEvent := &CommentEvent{commentEntity: commentEntity}

		assert.Equal(t, commentEntity, commentEvent.CommentEntity())
	})
}
