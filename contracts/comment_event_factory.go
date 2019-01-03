package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	CommentEventFactory interface {
		CreateCommentCreatedEvent(commentEntity *entities.CommentEntity) CommentEvent
		CreateCommentUpdatedEvent(commentEntity *entities.CommentEntity) CommentEvent
		CreateCommentDeletedEvent(commentEntity *entities.CommentEntity) CommentEvent
	}
)
