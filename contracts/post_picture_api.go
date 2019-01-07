package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostPictureApi interface {
		ChangePostPicture(postId *models.PostId, postPictureId *models.FileId) errors.Error
		RemovePostPicture(postId *models.PostId) errors.Error
	}
)
