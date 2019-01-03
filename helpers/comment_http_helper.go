package helpers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
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

func (*commentHttpHelperImpl) ParseCommentId(httpContext contracts.HttpContext) (*models.CommentId, common.Error) {
	return common.ParseModelId(httpContext.Parameter(CommentIdParameterName))
}

