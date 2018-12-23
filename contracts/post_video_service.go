package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostVideoService interface {
		ChangePostVideo(postEntity *entities.PostEntity, postVideo *entities.FileEntity) common.Error
		RemovePostVideo(postEntity *entities.PostEntity) common.Error
	}
)
