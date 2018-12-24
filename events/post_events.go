package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	PostCreatedEventName = "PostCreatedEvent"
	PostUpdatedEventName = "PostUpdatedEvent"
	PostDeletedEventName = "PostDeletedEvent"
)

type (
	PostEvent struct {
		*Event
		postEntity *entities.PostEntity
	}
)

func (e *PostEvent) PostEntity() *entities.PostEntity {
	return e.postEntity
}

func NewPostCreatedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	event := &Event{name: PostCreatedEventName}
	return &PostEvent{postEntity: postEntity, Event: event}
}

func NewPostUpdatedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	event := &Event{name: PostUpdatedEventName}
	return &PostEvent{postEntity: postEntity, Event: event}
}

func NewPostDeletedEvent(postEntity *entities.PostEntity) contracts.PostEvent {
	event := &Event{name: PostDeletedEventName}
	return &PostEvent{postEntity: postEntity, Event: event}
}
