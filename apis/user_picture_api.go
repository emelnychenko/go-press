package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	userPictureApiImpl struct {
		userService        contracts.UserService
		fileService        contracts.FileService
		userPictureService contracts.UserPictureService
	}
)

func NewUserPictureApi(
	userService contracts.UserService,
	fileService contracts.FileService,
	userPictureService contracts.UserPictureService,
) (userPictureApi contracts.UserPictureApi) {
	return &userPictureApiImpl{
		userService:        userService,
		fileService:        fileService,
		userPictureService: userPictureService,
	}
}

func (a *userPictureApiImpl) ChangeUserPicture(userId *models.UserId, userPictureId *models.FileId) (err common.Error) {
	userEntity, err := a.userService.GetUser(userId)

	if nil != err {
		return
	}

	userPicture, err := a.fileService.GetFile(userPictureId)

	if nil != err {
		return
	}

	return a.userPictureService.ChangeUserPicture(userEntity, userPicture)
}

func (a *userPictureApiImpl) RemoveUserPicture(userId *models.UserId) (err common.Error) {
	userEntity, err := a.userService.GetUser(userId)

	if nil != err {
		return
	}

	return a.userPictureService.RemoveUserPicture(userEntity)
}
