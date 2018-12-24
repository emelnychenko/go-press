package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostVideoApi interface {
		ChangePostVideo(postId *models.PostId, postVideoId *models.FileId) common.Error
		RemovePostVideo(postId *models.PostId) common.Error
	}
)
