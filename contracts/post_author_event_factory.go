package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostAuthorEventFactory interface {
		CreatePostAuthorChangedEvent(postEntity *entities.PostEntity, postAuthorEntity *entities.UserEntity) PostAuthorEvent
	}
)
