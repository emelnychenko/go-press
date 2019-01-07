package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	CategoryEventFactory interface {
		CreateCategoryCreatedEvent(*entities.CategoryEntity) CategoryEvent
		CreateCategoryUpdatedEvent(*entities.CategoryEntity) CategoryEvent
		CreateCategoryDeletedEvent(*entities.CategoryEntity) CategoryEvent
		CreateCategoryParentChangedEvent(*entities.CategoryEntity) CategoryEvent
		CreateCategoryParentRemovedEvent(*entities.CategoryEntity) CategoryEvent
	}
)
