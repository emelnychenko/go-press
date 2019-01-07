package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostVideoApi interface {
		ChangePostVideo(postId *models.PostId, postVideoId *models.FileId) errors.Error
		RemovePostVideo(postId *models.PostId) errors.Error
	}
)
