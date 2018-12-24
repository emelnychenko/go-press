package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	UserPictureEventFactory interface {
		CreateUserPictureChangedEvent(userEntity *entities.UserEntity, userPictureEntity *entities.FileEntity) UserPictureEvent
		CreateUserPictureRemovedEvent(userEntity *entities.UserEntity) UserPictureEvent
	}
)
