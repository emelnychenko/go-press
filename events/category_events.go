package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	CategoryCreatedEventName = "CategoryCreatedEvent"
	CategoryUpdatedEventName = "CategoryUpdatedEvent"
	CategoryDeletedEventName = "CategoryDeletedEvent"
)

type (
	CategoryEvent struct {
		*Event
		categoryEntity *entities.CategoryEntity
	}
)

func (e *CategoryEvent) CategoryEntity() *entities.CategoryEntity {
	return e.categoryEntity
}

func NewCategoryCreatedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	event := &Event{name: CategoryCreatedEventName}
	return &CategoryEvent{categoryEntity: categoryEntity, Event: event}
}

func NewCategoryUpdatedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	event := &Event{name: CategoryUpdatedEventName}
	return &CategoryEvent{categoryEntity: categoryEntity, Event: event}
}

func NewCategoryDeletedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	event := &Event{name: CategoryDeletedEventName}
	return &CategoryEvent{categoryEntity: categoryEntity, Event: event}
}
