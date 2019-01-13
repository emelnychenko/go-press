package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	postCategoryEventFactoryImpl struct {
	}
)

//NewPostCategoryEventFactory
func NewPostCategoryEventFactory() contracts.PostCategoryEventFactory {
	return new(postCategoryEventFactoryImpl)
}

//CreatePostCategoryAddedEvent
func (*postCategoryEventFactoryImpl) CreatePostCategoryAddedEvent(
	postEntity *entities.PostEntity, categoryEntity *entities.CategoryEntity,
) contracts.PostCategoryEvent {
	return events.NewPostCategoryAddedEvent(postEntity, categoryEntity)
}

//CreatePostCategoryRemovedEvent
func (*postCategoryEventFactoryImpl) CreatePostCategoryRemovedEvent(
	postEntity *entities.PostEntity, categoryEntity *entities.CategoryEntity,
) contracts.PostCategoryEvent {
	return events.NewPostCategoryRemovedEvent(postEntity, categoryEntity)
}
