package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	commentEventFactoryImpl struct {
	}
)

func NewCommentEventFactory() contracts.CommentEventFactory {
	return new(commentEventFactoryImpl)
}

func (*commentEventFactoryImpl) CreateCommentCreatedEvent(commentEntity *entities.CommentEntity) contracts.CommentEvent {
	return events.NewCommentCreatedEvent(commentEntity)
}

func (*commentEventFactoryImpl) CreateCommentUpdatedEvent(commentEntity *entities.CommentEntity) contracts.CommentEvent {
	return events.NewCommentUpdatedEvent(commentEntity)
}

func (*commentEventFactoryImpl) CreateCommentDeletedEvent(commentEntity *entities.CommentEntity) contracts.CommentEvent {
	return events.NewCommentDeletedEvent(commentEntity)
}
