package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
)

type userPictureControllerImpl struct {
	userHttpHelper contracts.UserHttpHelper
	fileHttpHelper contracts.FileHttpHelper
	userPictureApi contracts.UserPictureApi
}

func NewUserPictureController(
	userHttpHelper contracts.UserHttpHelper,
	fileHttpHelper contracts.FileHttpHelper,
	userPictureApi contracts.UserPictureApi,
) contracts.UserPictureController {
	return &userPictureControllerImpl{
		userHttpHelper,
		fileHttpHelper,
		userPictureApi,
	}
}

func (c *userPictureControllerImpl) ChangeUserPicture(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
	userId, err := c.userHttpHelper.ParseUserId(httpContext)

	if err != nil {
		return
	}

	fileId, err := c.fileHttpHelper.ParseFileId(httpContext)

	if err != nil {
		return
	}

	err = c.userPictureApi.ChangeUserPicture(userId, fileId)

	return
}

func (c *userPictureControllerImpl) RemoveUserPicture(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
	userId, err := c.userHttpHelper.ParseUserId(httpContext)

	if err != nil {
		return
	}

	err = c.userPictureApi.RemoveUserPicture(userId)

	return
}
