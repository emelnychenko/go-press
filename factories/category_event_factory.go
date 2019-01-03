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

func NewCategoryEventFactory() contracts.CategoryEventFactory {
	return new(categoryEventFactoryImpl)
}

func (*categoryEventFactoryImpl) CreateCategoryCreatedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	return events.NewCategoryCreatedEvent(categoryEntity)
}

func (*categoryEventFactoryImpl) CreateCategoryUpdatedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	return events.NewCategoryUpdatedEvent(categoryEntity)
}

func (*categoryEventFactoryImpl) CreateCategoryDeletedEvent(categoryEntity *entities.CategoryEntity) contracts.CategoryEvent {
	return events.NewCategoryDeletedEvent(categoryEntity)
}
