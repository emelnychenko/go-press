package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	PostVideoService interface {
		ChangePostVideo(postEntity *entities.PostEntity, postVideoEntity *entities.FileEntity) errors.Error
		RemovePostVideo(postEntity *entities.PostEntity) errors.Error
	}
)
