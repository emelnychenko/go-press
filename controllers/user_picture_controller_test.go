package controllers

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPictureController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewUserPictureController", func(t *testing.T) {
		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		userPictureApi := mocks.NewMockUserPictureApi(ctrl)

		userPictureController, isUserPictureController := NewUserPictureController(
			userHttpHelper, fileHttpHelper, userPictureApi,
		).(*userPictureControllerImpl)

		assert.True(t, isUserPictureController)
		assert.Equal(t, userHttpHelper, userPictureController.userHttpHelper)
		assert.Equal(t, fileHttpHelper, userPictureController.fileHttpHelper)
		assert.Equal(t, userPictureApi, userPictureController.userPictureApi)
	})

	t.Run("ChangeUserPicture", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		userId := new(models.UserId)
		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userPictureId := new(models.FileId)
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(userPictureId, nil)

		userPictureApi := mocks.NewMockUserPictureApi(ctrl)
		userPictureApi.EXPECT().ChangeUserPicture(userId, userPictureId).Return(nil)

		userPictureController := &userPictureControllerImpl{
			userHttpHelper: userHttpHelper,
			fileHttpHelper: fileHttpHelper,
			userPictureApi: userPictureApi,
		}
		response, err := userPictureController.ChangeUserPicture(httpContext)

		assert.Nil(t, response, err)
	})

	t.Run("ChangeUserPicture:ParseUserIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(nil, systemErr)

		userPictureController := &userPictureControllerImpl{
			userHttpHelper: userHttpHelper,
		}
		_, err := userPictureController.ChangeUserPicture(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserPicture:ParseFileIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		userId := new(models.UserId)
		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(nil, systemErr)

		userPictureController := &userPictureControllerImpl{
			userHttpHelper: userHttpHelper,
			fileHttpHelper: fileHttpHelper,
		}
		_, err := userPictureController.ChangeUserPicture(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserPicture:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		userId := new(models.UserId)
		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userPictureId := new(models.FileId)
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(userPictureId, nil)

		userPictureApi := mocks.NewMockUserPictureApi(ctrl)
		userPictureApi.EXPECT().ChangeUserPicture(userId, userPictureId).Return(systemErr)

		userPictureController := &userPictureControllerImpl{
			userHttpHelper: userHttpHelper,
			fileHttpHelper: fileHttpHelper,
			userPictureApi: userPictureApi,
		}
		_, err := userPictureController.ChangeUserPicture(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("RemoveUserPicture", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		userId := new(models.UserId)
		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userPictureApi := mocks.NewMockUserPictureApi(ctrl)
		userPictureApi.EXPECT().RemoveUserPicture(userId).Return(nil)

		userPictureController := &userPictureControllerImpl{
			userHttpHelper: userHttpHelper,
			userPictureApi: userPictureApi,
		}
		response, err := userPictureController.RemoveUserPicture(httpContext)

		assert.Nil(t, response, err)
	})

	t.Run("RemoveUserPicture:ParseUserIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(nil, systemErr)

		userPictureController := &userPictureControllerImpl{
			userHttpHelper: userHttpHelper,
		}
		_, err := userPictureController.RemoveUserPicture(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("RemoveUserPicture:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		userId := new(models.UserId)
		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userPictureApi := mocks.NewMockUserPictureApi(ctrl)
		userPictureApi.EXPECT().RemoveUserPicture(userId).Return(systemErr)

		userPictureController := &userPictureControllerImpl{
			userHttpHelper: userHttpHelper,
			userPictureApi: userPictureApi,
		}
		_, err := userPictureController.RemoveUserPicture(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
