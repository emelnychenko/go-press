package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
)

type (
	UserPictureService interface {
		ChangeUserPicture(userEntity *entities.UserEntity, userPicture *entities.FileEntity) common.Error
		RemoveUserPicture(userEntity *entities.UserEntity) common.Error
	}
)
