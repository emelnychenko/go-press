package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	CategoryEventFactory interface {
		CreateCategoryCreatedEvent(categoryEntity *entities.CategoryEntity) CategoryEvent
		CreateCategoryUpdatedEvent(categoryEntity *entities.CategoryEntity) CategoryEvent
		CreateCategoryDeletedEvent(categoryEntity *entities.CategoryEntity) CategoryEvent
		CreateCategoryParentChangedEvent(categoryEntity *entities.CategoryEntity) CategoryEvent
		CreateCategoryParentRemovedEvent(categoryEntity *entities.CategoryEntity) CategoryEvent
	}
)
