package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostPictureApi interface {
		ChangePostPicture(postId *models.PostId, postPictureId *models.FileId) common.Error
		RemovePostPicture(postId *models.PostId) common.Error
	}
)
