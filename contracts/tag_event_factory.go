package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	TagEventFactory interface {
		CreateTagCreatedEvent(tagEntity *entities.TagEntity) TagEvent
		CreateTagUpdatedEvent(tagEntity *entities.TagEntity) TagEvent
		CreateTagDeletedEvent(tagEntity *entities.TagEntity) TagEvent
	}
)
