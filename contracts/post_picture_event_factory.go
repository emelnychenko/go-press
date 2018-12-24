package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	PostPictureEventFactory interface {
		CreatePostPictureChangedEvent(postEntity *entities.PostEntity, postPictureEntity *entities.FileEntity) PostPictureEvent
		CreatePostPictureRemovedEvent(postEntity *entities.PostEntity) PostPictureEvent
	}
)
