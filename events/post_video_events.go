package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	PostVideoChangedEventName = "PostVideoChangedEvent"
	PostVideoRemovedEventName = "PostVideoRemovedEvent"
)

type (
	PostVideoEvent struct {
		*Event
		postEntity *entities.PostEntity
		postVideo *entities.FileEntity
	}
)

func (e *PostVideoEvent) PostEntity() *entities.PostEntity {
	return e.postEntity
}

func (e *PostVideoEvent) PostVideoEntity() *entities.FileEntity {
	return e.postVideo
}

func NewPostVideoChangedEvent(
	postEntity *entities.PostEntity,
	postVideo *entities.FileEntity,
) contracts.PostVideoEvent {
	event := &Event{name: PostVideoChangedEventName}
	return &PostVideoEvent{postEntity: postEntity, postVideo: postVideo, Event: event}
}

func NewPostVideoRemovedEvent(postEntity *entities.PostEntity) contracts.PostVideoEvent {
	event := &Event{name: PostVideoRemovedEventName}
	return &PostVideoEvent{postEntity: postEntity, postVideo: nil, Event: event}
}
