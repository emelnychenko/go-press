package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
)

type postVideoControllerImpl struct {
	postHttpHelper contracts.PostHttpHelper
	fileHttpHelper contracts.FileHttpHelper
	postVideoApi   contracts.PostVideoApi
}

func NewPostVideoController(
	postHttpHelper contracts.PostHttpHelper,
	fileHttpHelper contracts.FileHttpHelper,
	postVideoApi contracts.PostVideoApi,
) contracts.PostVideoController {
	return &postVideoControllerImpl{
		postHttpHelper,
		fileHttpHelper,
		postVideoApi,
	}
}

func (c *postVideoControllerImpl) ChangePostVideo(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	fileId, err := c.fileHttpHelper.ParseFileId(httpContext)

	if err != nil {
		return
	}

	err = c.postVideoApi.ChangePostVideo(postId, fileId)

	return
}

func (c *postVideoControllerImpl) RemovePostVideo(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	err = c.postVideoApi.RemovePostVideo(postId)

	return
}
