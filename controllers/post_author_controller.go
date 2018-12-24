package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
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

func (c *postAuthorControllerImpl) ChangePostAuthor(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
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
