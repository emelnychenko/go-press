package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostPictureService interface {
		ChangePostPicture(postEntity *entities.PostEntity, postPictureEntity *entities.FileEntity) common.Error
		RemovePostPicture(postEntity *entities.PostEntity) common.Error
	}
)
