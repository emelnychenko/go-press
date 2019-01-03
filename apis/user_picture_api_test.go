package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUserPictureApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewUserPictureApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		userPictureEventFactory := mocks.NewMockUserPictureEventFactory(ctrl)
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		userService := mocks.NewMockUserService(ctrl)
		fileService := mocks.NewMockFileService(ctrl)
		userPictureService := mocks.NewMockUserPictureService(ctrl)

		userPictureApi, isUserPictureApi := NewUserPictureApi(
			eventDispatcher, userPictureEventFactory, contentTypeValidator, userService, fileService, userPictureService,
		).(*userPictureApiImpl)

		assert.True(t, isUserPictureApi)
		assert.Equal(t, eventDispatcher, userPictureApi.eventDispatcher)
		assert.Equal(t, userPictureEventFactory, userPictureApi.userPictureEventFactory)
		assert.Equal(t, contentTypeValidator, userPictureApi.contentTypeValidator)
		assert.Equal(t, userService, userPictureApi.userService)
		assert.Equal(t, fileService, userPictureApi.fileService)
		assert.Equal(t, userPictureService, userPictureApi.userPictureService)
	})

	t.Run("ChangeUserPicture", func(t *testing.T) {
		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)

		contentType := "image/png"
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		contentTypeValidator.EXPECT().ValidateImage(contentType).Return(nil)

		userPictureId := new(models.FileId)
		userPictureEntity := &entities.FileEntity{Type: contentType}
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(userPictureId).Return(userPictureEntity, nil)

		userPictureEvent := new(events.UserPictureEvent)
		userPictureEventFactory := mocks.NewMockUserPictureEventFactory(ctrl)
		userPictureEventFactory.EXPECT().CreateUserPictureChangedEvent(userEntity, userPictureEntity).Return(userPictureEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(userPictureEvent)

		userPictureService := mocks.NewMockUserPictureService(ctrl)
		userPictureService.EXPECT().ChangeUserPicture(userEntity, userPictureEntity).Return(nil)

		userPictureApi := &userPictureApiImpl{
			eventDispatcher:         eventDispatcher,
			userPictureEventFactory: userPictureEventFactory,
			contentTypeValidator:    contentTypeValidator,
			userService:             userService,
			fileService:             fileService,
			userPictureService:      userPictureService,
		}
		assert.Nil(t, userPictureApi.ChangeUserPicture(userId, userPictureId))
	})

	t.Run("ChangeUserPicture:GetUserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		userId := new(models.UserId)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(nil, systemErr)

		userPictureId := new(models.FileId)
		userPictureApi := &userPictureApiImpl{userService: userService}
		err := userPictureApi.ChangeUserPicture(userId, userPictureId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserPicture:GetFileError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)

		userPictureId := new(models.FileId)
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(userPictureId).Return(nil, systemErr)

		userPictureApi := &userPictureApiImpl{
			userService: userService,
			fileService: fileService,
		}
		err := userPictureApi.ChangeUserPicture(userId, userPictureId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserPicture:ValidateImageError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)

		contentType := "image/png"
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		contentTypeValidator.EXPECT().ValidateImage(contentType).Return(systemErr)

		userPictureId := new(models.FileId)
		userPictureEntity := &entities.FileEntity{Type: contentType}
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(userPictureId).Return(userPictureEntity, nil)

		userPictureApi := &userPictureApiImpl{
			contentTypeValidator: contentTypeValidator,
			userService:          userService,
			fileService:          fileService,
		}

		err := userPictureApi.ChangeUserPicture(userId, userPictureId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserPicture:ChangeUserPictureError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)

		contentType := "image/png"
		contentTypeValidator := mocks.NewMockContentTypeValidator(ctrl)
		contentTypeValidator.EXPECT().ValidateImage(contentType).Return(nil)

		userPictureId := new(models.FileId)
		userPictureEntity := &entities.FileEntity{Type: contentType}
		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(userPictureId).Return(userPictureEntity, nil)

		userPictureService := mocks.NewMockUserPictureService(ctrl)
		userPictureService.EXPECT().ChangeUserPicture(userEntity, userPictureEntity).Return(systemErr)

		userPictureApi := &userPictureApiImpl{
			contentTypeValidator: contentTypeValidator,
			userService:          userService,
			fileService:          fileService,
			userPictureService:   userPictureService,
		}

		err := userPictureApi.ChangeUserPicture(userId, userPictureId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemoveUserPicture", func(t *testing.T) {
		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)

		userPictureEvent := new(events.UserPictureEvent)
		userPictureEventFactory := mocks.NewMockUserPictureEventFactory(ctrl)
		userPictureEventFactory.EXPECT().CreateUserPictureRemovedEvent(userEntity).Return(userPictureEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(userPictureEvent)

		userPictureService := mocks.NewMockUserPictureService(ctrl)
		userPictureService.EXPECT().RemoveUserPicture(userEntity).Return(nil)

		userPictureApi := &userPictureApiImpl{
			eventDispatcher:         eventDispatcher,
			userPictureEventFactory: userPictureEventFactory,
			userService:             userService,
			userPictureService:      userPictureService,
		}
		assert.Nil(t, userPictureApi.RemoveUserPicture(userId))
	})

	t.Run("RemoveUserPicture:GetUserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		userId := new(models.UserId)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(nil, systemErr)

		userPictureApi := &userPictureApiImpl{userService: userService}
		err := userPictureApi.RemoveUserPicture(userId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemoveUserPicture:RemoveUserPictureError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)
		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)

		userPictureService := mocks.NewMockUserPictureService(ctrl)
		userPictureService.EXPECT().RemoveUserPicture(userEntity).Return(systemErr)

		userPictureApi := &userPictureApiImpl{
			userService:        userService,
			userPictureService: userPictureService,
		}

		err := userPictureApi.RemoveUserPicture(userId)
		assert.Equal(t, systemErr, err)
	})
}
