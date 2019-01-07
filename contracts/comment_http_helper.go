package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	CommentHttpHelper interface {
		ParseCommentId(httpContext HttpContext) (*models.CommentId, errors.Error)
	}
)
