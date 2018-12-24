package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	PostAuthorChangedEventName = "PostAuthorChangedEvent"
)

type (
	PostAuthorEvent struct {
		*Event
		postEntity       *entities.PostEntity
		postAuthorEntity *entities.UserEntity
	}
)

func (e *PostAuthorEvent) PostEntity() *entities.PostEntity {
	return e.postEntity
}

func (e *PostAuthorEvent) PostAuthorEntity() *entities.UserEntity {
	return e.postAuthorEntity
}

func NewPostAuthorChangedEvent(
	postEntity *entities.PostEntity,
	postAuthorEntity *entities.UserEntity,
) contracts.PostAuthorEvent {
	event := &Event{name: PostAuthorChangedEventName}
	return &PostAuthorEvent{postEntity: postEntity, postAuthorEntity: postAuthorEntity, Event: event}
}
