package controllers

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostAuthorController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostAuthorController", func(t *testing.T) {
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		postAuthorApi := mocks.NewMockPostAuthorApi(ctrl)

		postAuthorController, isPostAuthorController := NewPostAuthorController(
			postHttpHelper, userHttpHelper, postAuthorApi,
		).(*postAuthorControllerImpl)

		assert.True(t, isPostAuthorController)
		assert.Equal(t, postHttpHelper, postAuthorController.postHttpHelper)
		assert.Equal(t, userHttpHelper, postAuthorController.userHttpHelper)
		assert.Equal(t, postAuthorApi, postAuthorController.postAuthorApi)
	})

	t.Run("ChangePostAuthor", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postAuthorId := new(models.UserId)
		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(postAuthorId, nil)

		postAuthorApi := mocks.NewMockPostAuthorApi(ctrl)
		postAuthorApi.EXPECT().ChangePostAuthor(postId, postAuthorId).Return(nil)

		postAuthorController := &postAuthorControllerImpl{
			postHttpHelper: postHttpHelper,
			userHttpHelper: userHttpHelper,
			postAuthorApi:  postAuthorApi,
		}
		response, err := postAuthorController.ChangePostAuthor(httpContext)

		assert.Nil(t, response, err)
	})

	t.Run("ChangePostAuthor:ParsePostIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postAuthorController := &postAuthorControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		_, err := postAuthorController.ChangePostAuthor(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostAuthor:ParseUserIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(nil, systemErr)

		postAuthorController := &postAuthorControllerImpl{
			postHttpHelper: postHttpHelper,
			userHttpHelper: userHttpHelper,
		}
		_, err := postAuthorController.ChangePostAuthor(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangePostAuthor:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postAuthorId := new(models.UserId)
		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(postAuthorId, nil)

		postAuthorApi := mocks.NewMockPostAuthorApi(ctrl)
		postAuthorApi.EXPECT().ChangePostAuthor(postId, postAuthorId).Return(systemErr)

		postAuthorController := &postAuthorControllerImpl{
			postHttpHelper: postHttpHelper,
			userHttpHelper: userHttpHelper,
			postAuthorApi:  postAuthorApi,
		}
		_, err := postAuthorController.ChangePostAuthor(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
