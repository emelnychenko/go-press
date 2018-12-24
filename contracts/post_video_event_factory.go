package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	PostVideoEventFactory interface {
		CreatePostVideoChangedEvent(postEntity *entities.PostEntity, postVideo *entities.FileEntity) PostVideoEvent
		CreatePostVideoRemovedEvent(postEntity *entities.PostEntity) PostVideoEvent
	}
)
