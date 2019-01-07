package controllers

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostPictureController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostPictureController", func(t *testing.T) {
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		postPictureApi := mocks.NewMockPostPictureApi(ctrl)

		postPictureController, isPostPictureController := NewPostPictureController(
			postHttpHelper, fileHttpHelper, postPictureApi,
		).(*postPictureControllerImpl)

		assert.True(t, isPostPictureController)
		assert.Equal(t, postHttpHelper, postPictureController.postHttpHelper)
		assert.Equal(t, fileHttpHelper, postPictureController.fileHttpHelper)
		assert.Equal(t, postPictureApi, postPictureController.postPictureApi)
	})

	t.Run("ChangePostPicture", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postPictureId := new(models.FileId)
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(postPictureId, nil)

		postPictureApi := mocks.NewMockPostPictureApi(ctrl)
		postPictureApi.EXPECT().ChangePostPicture(postId, postPictureId).Return(nil)

		postPictureController := &postPictureControllerImpl{
			postHttpHelper: postHttpHelper,
			fileHttpHelper: fileHttpHelper,
			postPictureApi: postPictureApi,
		}
		response, err := postPictureController.ChangePostPicture(httpContext)

		assert.Nil(t, response, err)
	})

	t.Run("ChangePostPicture:ParsePostIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postPictureController := &postPictureControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		_, err := postPictureController.ChangePostPicture(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostPicture:ParseFileIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(nil, systemErr)

		postPictureController := &postPictureControllerImpl{
			postHttpHelper: postHttpHelper,
			fileHttpHelper: fileHttpHelper,
		}
		_, err := postPictureController.ChangePostPicture(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostPicture:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postPictureId := new(models.FileId)
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(postPictureId, nil)

		postPictureApi := mocks.NewMockPostPictureApi(ctrl)
		postPictureApi.EXPECT().ChangePostPicture(postId, postPictureId).Return(systemErr)

		postPictureController := &postPictureControllerImpl{
			postHttpHelper: postHttpHelper,
			fileHttpHelper: fileHttpHelper,
			postPictureApi: postPictureApi,
		}
		_, err := postPictureController.ChangePostPicture(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostPicture", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postPictureApi := mocks.NewMockPostPictureApi(ctrl)
		postPictureApi.EXPECT().RemovePostPicture(postId).Return(nil)

		postPictureController := &postPictureControllerImpl{
			postHttpHelper: postHttpHelper,
			postPictureApi: postPictureApi,
		}
		response, err := postPictureController.RemovePostPicture(httpContext)

		assert.Nil(t, response, err)
	})

	t.Run("RemovePostPicture:ParsePostIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postPictureController := &postPictureControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		_, err := postPictureController.RemovePostPicture(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostPicture:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postPictureApi := mocks.NewMockPostPictureApi(ctrl)
		postPictureApi.EXPECT().RemovePostPicture(postId).Return(systemErr)

		postPictureController := &postPictureControllerImpl{
			postHttpHelper: postHttpHelper,
			postPictureApi: postPictureApi,
		}
		_, err := postPictureController.RemovePostPicture(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
