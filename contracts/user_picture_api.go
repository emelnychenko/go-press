package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserPictureApi interface {
		ChangeUserPicture(userId *models.UserId, userPictureId *models.FileId) common.Error
		RemoveUserPicture(userId *models.UserId) common.Error
	}
)
