package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	PostEventFactory interface {
		CreatePostCreatedEvent(postEntity *entities.PostEntity) PostEvent
		CreatePostUpdatedEvent(postEntity *entities.PostEntity) PostEvent
		CreatePostDeletedEvent(postEntity *entities.PostEntity) PostEvent
	}
)
