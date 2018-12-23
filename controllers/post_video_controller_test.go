package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostVideoController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostVideoController", func(t *testing.T) {
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		postVideoApi := mocks.NewMockPostVideoApi(ctrl)

		postVideoController, isPostVideoController := NewPostVideoController(
			postHttpHelper, fileHttpHelper, postVideoApi,
		).(*postVideoControllerImpl)

		assert.True(t, isPostVideoController)
		assert.Equal(t, postHttpHelper, postVideoController.postHttpHelper)
		assert.Equal(t, fileHttpHelper, postVideoController.fileHttpHelper)
		assert.Equal(t, postVideoApi, postVideoController.postVideoApi)
	})

	t.Run("ChangePostVideo", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postVideoId := new(models.FileId)
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(postVideoId, nil)

		postVideoApi := mocks.NewMockPostVideoApi(ctrl)
		postVideoApi.EXPECT().ChangePostVideo(postId, postVideoId).Return(nil)

		postVideoController := &postVideoControllerImpl{
			postHttpHelper: postHttpHelper,
			fileHttpHelper: fileHttpHelper,
			postVideoApi: postVideoApi,
		}
		response, err := postVideoController.ChangePostVideo(httpContext)

		assert.Nil(t, response, err)
	})

	t.Run("ChangePostVideo:ParsePostIdError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postVideoController := &postVideoControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		_, err := postVideoController.ChangePostVideo(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostVideo:ParseFileIdError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(nil, systemErr)

		postVideoController := &postVideoControllerImpl{
			postHttpHelper: postHttpHelper,
			fileHttpHelper: fileHttpHelper,
		}
		_, err := postVideoController.ChangePostVideo(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostVideo:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postVideoId := new(models.FileId)
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(postVideoId, nil)

		postVideoApi := mocks.NewMockPostVideoApi(ctrl)
		postVideoApi.EXPECT().ChangePostVideo(postId, postVideoId).Return(systemErr)

		postVideoController := &postVideoControllerImpl{
			postHttpHelper: postHttpHelper,
			fileHttpHelper: fileHttpHelper,
			postVideoApi: postVideoApi,
		}
		_, err := postVideoController.ChangePostVideo(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostVideo", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postVideoApi := mocks.NewMockPostVideoApi(ctrl)
		postVideoApi.EXPECT().RemovePostVideo(postId).Return(nil)

		postVideoController := &postVideoControllerImpl{
			postHttpHelper: postHttpHelper,
			postVideoApi: postVideoApi,
		}
		response, err := postVideoController.RemovePostVideo(httpContext)

		assert.Nil(t, response, err)
	})

	t.Run("RemovePostVideo:ParsePostIdError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postVideoController := &postVideoControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		_, err := postVideoController.RemovePostVideo(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostVideo:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)


		postVideoApi := mocks.NewMockPostVideoApi(ctrl)
		postVideoApi.EXPECT().RemovePostVideo(postId).Return(systemErr)

		postVideoController := &postVideoControllerImpl{
			postHttpHelper: postHttpHelper,
			postVideoApi: postVideoApi,
		}
		_, err := postVideoController.RemovePostVideo(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
