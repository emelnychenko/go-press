package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	TagCreatedEventName = "TagCreatedEvent"
	TagUpdatedEventName = "TagUpdatedEvent"
	TagDeletedEventName = "TagDeletedEvent"
)

type (
	TagEvent struct {
		*Event
		tagEntity *entities.TagEntity
	}
)

func (e *TagEvent) TagEntity() *entities.TagEntity {
	return e.tagEntity
}

func NewTagCreatedEvent(tagEntity *entities.TagEntity) contracts.TagEvent {
	event := &Event{name: TagCreatedEventName}
	return &TagEvent{tagEntity: tagEntity, Event: event}
}

func NewTagUpdatedEvent(tagEntity *entities.TagEntity) contracts.TagEvent {
	event := &Event{name: TagUpdatedEventName}
	return &TagEvent{tagEntity: tagEntity, Event: event}
}

func NewTagDeletedEvent(tagEntity *entities.TagEntity) contracts.TagEvent {
	event := &Event{name: TagDeletedEventName}
	return &TagEvent{tagEntity: tagEntity, Event: event}
}
