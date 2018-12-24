package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
)

type (
	UserPictureService interface {
		ChangeUserPicture(userEntity *entities.UserEntity, userPictureEntity *entities.FileEntity) common.Error
		RemoveUserPicture(userEntity *entities.UserEntity) common.Error
	}
)
