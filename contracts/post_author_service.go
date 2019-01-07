package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	PostAuthorService interface {
		ChangePostAuthor(postEntity *entities.PostEntity, postAuthorEntity *entities.UserEntity) errors.Error
	}
)
