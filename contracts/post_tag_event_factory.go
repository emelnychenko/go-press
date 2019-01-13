package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	PostTagEventFactory interface {
		CreatePostTagAddedEvent(*entities.PostEntity, *entities.TagEntity) PostTagEvent
		CreatePostTagRemovedEvent(*entities.PostEntity, *entities.TagEntity) PostTagEvent
	}
)
