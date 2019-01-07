package helpers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

const (
	CommentIdParameterName = "commentId"
)

type (
	commentHttpHelperImpl struct {
	}
)

func NewCommentHttpHelper() contracts.CommentHttpHelper {
	return new(commentHttpHelperImpl)
}

func (*commentHttpHelperImpl) ParseCommentId(httpContext contracts.HttpContext) (*models.CommentId, errors.Error) {
	return models.ParseModelId(httpContext.Parameter(CommentIdParameterName))
}
