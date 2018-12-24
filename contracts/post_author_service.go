package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostAuthorService interface {
		ChangePostAuthor(postEntity *entities.PostEntity, postAuthorEntity *entities.UserEntity) common.Error
	}
)
