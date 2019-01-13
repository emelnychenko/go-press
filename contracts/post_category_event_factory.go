package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	PostCategoryEventFactory interface {
		CreatePostCategoryAddedEvent(*entities.PostEntity, *entities.CategoryEntity) PostCategoryEvent
		CreatePostCategoryRemovedEvent(*entities.PostEntity, *entities.CategoryEntity) PostCategoryEvent
	}
)
