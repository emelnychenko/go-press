package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	userPictureApiImpl struct {
		eventDispatcher         contracts.EventDispatcher
		userPictureEventFactory contracts.UserPictureEventFactory
		userService             contracts.UserService
		fileService             contracts.FileService
		userPictureService      contracts.UserPictureService
	}
)

func NewUserPictureApi(
	eventDispatcher contracts.EventDispatcher,
	userPictureEventFactory contracts.UserPictureEventFactory,
	userService contracts.UserService,
	fileService contracts.FileService,
	userPictureService contracts.UserPictureService,
) (userPictureApi contracts.UserPictureApi) {
	return &userPictureApiImpl{
		eventDispatcher:         eventDispatcher,
		userPictureEventFactory: userPictureEventFactory,
		userService:             userService,
		fileService:             fileService,
		userPictureService:      userPictureService,
	}
}

func (a *userPictureApiImpl) ChangeUserPicture(userId *models.UserId, userPictureId *models.FileId) (err common.Error) {
	userEntity, err := a.userService.GetUser(userId)

	if nil != err {
		return
	}

	userPictureEntity, err := a.fileService.GetFile(userPictureId)

	if nil != err {
		return
	}

	err = a.userPictureService.ChangeUserPicture(userEntity, userPictureEntity)

	userPictureEvent := a.userPictureEventFactory.CreateUserPictureChangedEvent(userEntity, userPictureEntity)
	a.eventDispatcher.Dispatch(userPictureEvent)

	return
}

func (a *userPictureApiImpl) RemoveUserPicture(userId *models.UserId) (err common.Error) {
	userEntity, err := a.userService.GetUser(userId)

	if nil != err {
		return
	}

	err = a.userPictureService.RemoveUserPicture(userEntity)

	userPictureEvent := a.userPictureEventFactory.CreateUserPictureRemovedEvent(userEntity)
	a.eventDispatcher.Dispatch(userPictureEvent)

	return
}
