package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostVideoApi interface {
		ChangePostVideo(userId *models.PostId, userVideoId *models.FileId) common.Error
		RemovePostVideo(userId *models.PostId) common.Error
	}
)
