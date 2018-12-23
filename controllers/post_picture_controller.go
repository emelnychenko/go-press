package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
)

type postPictureControllerImpl struct {
	postHttpHelper contracts.PostHttpHelper
	fileHttpHelper contracts.FileHttpHelper
	postPictureApi contracts.PostPictureApi
}

func NewPostPictureController(
	postHttpHelper contracts.PostHttpHelper,
	fileHttpHelper contracts.FileHttpHelper,
	postPictureApi contracts.PostPictureApi,
) contracts.PostPictureController {
	return &postPictureControllerImpl{
		postHttpHelper,
		fileHttpHelper,
		postPictureApi,
	}
}

func (c *postPictureControllerImpl) ChangePostPicture(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	fileId, err := c.fileHttpHelper.ParseFileId(httpContext)

	if err != nil {
		return
	}

	err = c.postPictureApi.ChangePostPicture(postId, fileId)

	return
}

func (c *postPictureControllerImpl) RemovePostPicture(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	err = c.postPictureApi.RemovePostPicture(postId)

	return
}
