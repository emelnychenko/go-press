package helpers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

const (
	PostIdParameterName = "postId"
)

type (
	postHttpHelperImpl struct {
	}
)

func NewPostHttpHelper() contracts.PostHttpHelper {
	return new(postHttpHelperImpl)
}

func (*postHttpHelperImpl) ParsePostId(httpContext contracts.HttpContext) (*models.PostId, errors.Error) {
	return models.ParseModelId(httpContext.Parameter(PostIdParameterName))
}
