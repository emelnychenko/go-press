package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserPictureApi interface {
		ChangeUserPicture(userId *models.UserId, userPictureId *models.FileId) errors.Error
		RemoveUserPicture(userId *models.UserId) errors.Error
	}
)
