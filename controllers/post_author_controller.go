package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
)

type postAuthorControllerImpl struct {
	postHttpHelper contracts.PostHttpHelper
	userHttpHelper contracts.UserHttpHelper
	postAuthorApi  contracts.PostAuthorApi
}

func NewPostAuthorController(
	postHttpHelper contracts.PostHttpHelper,
	userHttpHelper contracts.UserHttpHelper,
	postAuthorApi contracts.PostAuthorApi,
) contracts.PostAuthorController {
	return &postAuthorControllerImpl{
		postHttpHelper,
		userHttpHelper,
		postAuthorApi,
	}
}

func (c *postAuthorControllerImpl) ChangePostAuthor(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	userId, err := c.userHttpHelper.ParseUserId(httpContext)

	if err != nil {
		return
	}

	err = c.postAuthorApi.ChangePostAuthor(postId, userId)

	return
}
