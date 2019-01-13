package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	PostCategoryAddedEventName = "PostCategoryAddedEvent"
	PostCategoryRemovedEventName = "PostCategoryRemovedEvent"
)

type (
	PostCategoryEvent struct {
		*Event
		postEntity  *entities.PostEntity
		categoryEntity *entities.CategoryEntity
	}
)

//PostEntity
func (e *PostCategoryEvent) PostEntity() *entities.PostEntity {
	return e.postEntity
}

//CategoryEntity
func (e *PostCategoryEvent) CategoryEntity() *entities.CategoryEntity {
	return e.categoryEntity
}

//NewPostCategoryAddedEvent
func NewPostCategoryAddedEvent(
	postEntity *entities.PostEntity, categoryEntity *entities.CategoryEntity,
) contracts.PostCategoryEvent {
	event := &Event{name: PostCategoryAddedEventName}
	return &PostCategoryEvent{postEntity: postEntity, categoryEntity: categoryEntity, Event: event}
}

//NewPostCategoryRemovedEvent
func NewPostCategoryRemovedEvent(
	postEntity *entities.PostEntity, categoryEntity *entities.CategoryEntity,
) contracts.PostCategoryEvent {
	event := &Event{name: PostCategoryRemovedEventName}
	return &PostCategoryEvent{postEntity: postEntity, categoryEntity: categoryEntity, Event: event}
}
