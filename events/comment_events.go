package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	CommentCreatedEventName = "CommentCreatedEvent"
	CommentUpdatedEventName = "CommentUpdatedEvent"
	CommentDeletedEventName = "CommentDeletedEvent"
)

type (
	CommentEvent struct {
		*Event
		commentEntity *entities.CommentEntity
	}
)

func (e *CommentEvent) CommentEntity() *entities.CommentEntity {
	return e.commentEntity
}

func NewCommentCreatedEvent(commentEntity *entities.CommentEntity) contracts.CommentEvent {
	event := &Event{name: CommentCreatedEventName}
	return &CommentEvent{commentEntity: commentEntity, Event: event}
}

func NewCommentUpdatedEvent(commentEntity *entities.CommentEntity) contracts.CommentEvent {
	event := &Event{name: CommentUpdatedEventName}
	return &CommentEvent{commentEntity: commentEntity, Event: event}
}

func NewCommentDeletedEvent(commentEntity *entities.CommentEntity) contracts.CommentEvent {
	event := &Event{name: CommentDeletedEventName}
	return &CommentEvent{commentEntity: commentEntity, Event: event}
}
