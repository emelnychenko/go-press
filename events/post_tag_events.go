package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	PostTagAddedEventName = "PostTagAddedEvent"
	PostTagRemovedEventName = "PostTagRemovedEvent"
)

type (
	PostTagEvent struct {
		*Event
		postEntity  *entities.PostEntity
		tagEntity *entities.TagEntity
	}
)

//PostEntity
func (e *PostTagEvent) PostEntity() *entities.PostEntity {
	return e.postEntity
}

//TagEntity
func (e *PostTagEvent) TagEntity() *entities.TagEntity {
	return e.tagEntity
}

//NewPostTagAddedEvent
func NewPostTagAddedEvent(
	postEntity *entities.PostEntity, tagEntity *entities.TagEntity,
) contracts.PostTagEvent {
	event := &Event{name: PostTagAddedEventName}
	return &PostTagEvent{postEntity: postEntity, tagEntity: tagEntity, Event: event}
}

//NewPostTagRemovedEvent
func NewPostTagRemovedEvent(
	postEntity *entities.PostEntity, tagEntity *entities.TagEntity,
) contracts.PostTagEvent {
	event := &Event{name: PostTagRemovedEventName}
	return &PostTagEvent{postEntity: postEntity, tagEntity: tagEntity, Event: event}
}
