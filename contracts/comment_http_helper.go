package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	CommentHttpHelper interface {
		ParseCommentId(httpContext HttpContext) (*models.CommentId, common.Error)
	}
)
