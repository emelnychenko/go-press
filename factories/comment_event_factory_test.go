package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentEventFactory(t *testing.T) {
	t.Run("NewCommentEventFactory", func(t *testing.T) {
		_, isCommentEventFactory := NewCommentEventFactory().(*commentEventFactoryImpl)

		assert.True(t, isCommentEventFactory)
	})

	t.Run("CreateCommentCreatedEvent", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		commentEventFactory := new(commentEventFactoryImpl)
		commentEvent := commentEventFactory.CreateCommentCreatedEvent(commentEntity)

		assert.Equal(t, events.CommentCreatedEventName, commentEvent.Name())
		assert.Equal(t, commentEntity, commentEvent.CommentEntity())
	})

	t.Run("CreateCommentUpdatedEvent", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		commentEventFactory := new(commentEventFactoryImpl)
		commentEvent := commentEventFactory.CreateCommentUpdatedEvent(commentEntity)

		assert.Equal(t, events.CommentUpdatedEventName, commentEvent.Name())
		assert.Equal(t, commentEntity, commentEvent.CommentEntity())
	})

	t.Run("CreateCommentDeletedEvent", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		commentEventFactory := new(commentEventFactoryImpl)
		commentEvent := commentEventFactory.CreateCommentDeletedEvent(commentEntity)

		assert.Equal(t, events.CommentDeletedEventName, commentEvent.Name())
		assert.Equal(t, commentEntity, commentEvent.CommentEntity())
	})
}
