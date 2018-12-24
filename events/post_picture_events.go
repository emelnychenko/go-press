package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	PostPictureChangedEventName = "PostPictureChangedEvent"
	PostPictureRemovedEventName = "PostPictureRemovedEvent"
)

type (
	PostPictureEvent struct {
		*Event
		postEntity *entities.PostEntity
		postPicture *entities.FileEntity
	}
)

func (e *PostPictureEvent) PostEntity() *entities.PostEntity {
	return e.postEntity
}

func (e *PostPictureEvent) PostPictureEntity() *entities.FileEntity {
	return e.postPicture
}

func NewPostPictureChangedEvent(
	postEntity *entities.PostEntity,
	postPicture *entities.FileEntity,
) contracts.PostPictureEvent {
	event := &Event{name: PostPictureChangedEventName}
	return &PostPictureEvent{postEntity: postEntity, postPicture: postPicture, Event: event}
}

func NewPostPictureRemovedEvent(postEntity *entities.PostEntity) contracts.PostPictureEvent {
	event := &Event{name: PostPictureRemovedEventName}
	return &PostPictureEvent{postEntity: postEntity, postPicture: nil, Event: event}
}
