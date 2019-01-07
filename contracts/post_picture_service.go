package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	PostPictureService interface {
		ChangePostPicture(postEntity *entities.PostEntity, postPictureEntity *entities.FileEntity) errors.Error
		RemovePostPicture(postEntity *entities.PostEntity) errors.Error
	}
)
