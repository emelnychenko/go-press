package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	tagEventFactoryImpl struct {
	}
)

func NewTagEventFactory() contracts.TagEventFactory {
	return new(tagEventFactoryImpl)
}

func (*tagEventFactoryImpl) CreateTagCreatedEvent(tagEntity *entities.TagEntity) contracts.TagEvent {
	return events.NewTagCreatedEvent(tagEntity)
}

func (*tagEventFactoryImpl) CreateTagUpdatedEvent(tagEntity *entities.TagEntity) contracts.TagEvent {
	return events.NewTagUpdatedEvent(tagEntity)
}

func (*tagEventFactoryImpl) CreateTagDeletedEvent(tagEntity *entities.TagEntity) contracts.TagEvent {
	return events.NewTagDeletedEvent(tagEntity)
}
