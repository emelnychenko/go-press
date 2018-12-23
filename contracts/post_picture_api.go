package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostPictureApi interface {
		ChangePostPicture(userId *models.PostId, userPictureId *models.FileId) common.Error
		RemovePostPicture(userId *models.PostId) common.Error
	}
)
