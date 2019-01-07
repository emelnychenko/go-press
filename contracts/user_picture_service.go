package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	UserPictureService interface {
		ChangeUserPicture(userEntity *entities.UserEntity, userPictureEntity *entities.FileEntity) errors.Error
		RemoveUserPicture(userEntity *entities.UserEntity) errors.Error
	}
)
