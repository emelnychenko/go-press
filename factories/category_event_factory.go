package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	categoryEventFactoryImpl struct {
	}
)

//NewCategoryEventFactory
func NewCategoryEventFactory() contracts.CategoryEventFactory {
	return new(categoryEventFactoryImpl)
}

//CreateCategoryCreatedEvent
func (*categoryEventFactoryImpl) CreateCategoryCreatedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	return events.NewCategoryCreatedEvent(categoryEntity)
}

//CreateCategoryUpdatedEvent
func (*categoryEventFactoryImpl) CreateCategoryUpdatedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	return events.NewCategoryUpdatedEvent(categoryEntity)
}

//CreateCategoryDeletedEvent
func (*categoryEventFactoryImpl) CreateCategoryDeletedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	return events.NewCategoryDeletedEvent(categoryEntity)
}

//CreateCategoryParentChangedEvent
func (*categoryEventFactoryImpl) CreateCategoryParentChangedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	return events.NewCategoryParentChangedEvent(categoryEntity)
}

//CreateCategoryParentRemovedEvent
func (*categoryEventFactoryImpl) CreateCategoryParentRemovedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	return events.NewCategoryParentRemovedEvent(categoryEntity)
}
