package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostAuthorApi interface {
		ChangePostAuthor(postId *models.PostId, postAuthorId *models.UserId) errors.Error
	}
)
