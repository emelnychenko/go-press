package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostController", func(t *testing.T) {
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postApi := mocks.NewMockPostApi(ctrl)
		postController, isPostController := NewPostController(postHttpHelper, postModelFactory, postApi).(*postControllerImpl)

		assert.True(t, isPostController)
		assert.Equal(t, postHttpHelper, postController.postHttpHelper)
		assert.Equal(t, postModelFactory, postController.postModelFactory)
		assert.Equal(t, postApi, postController.postApi)
	})

	t.Run("ListPosts", func(t *testing.T) {
		var posts []*models.Post
		postApi := mocks.NewMockPostApi(ctrl)
		postApi.EXPECT().ListPosts().Return(posts, nil)

		postController := &postControllerImpl{postApi: postApi}
		response, err := postController.ListPosts(nil)

		assert.Equal(t, posts, response)
		assert.Nil(t, err)
	})

	t.Run("ListPosts:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		postApi := mocks.NewMockPostApi(ctrl)
		postApi.EXPECT().ListPosts().Return(nil, systemErr)

		postController := &postControllerImpl{postApi: postApi}
		response, err := postController.ListPosts(nil)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetPost", func(t *testing.T) {
		postId := new(models.PostId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		var post *models.Post
		postApi := mocks.NewMockPostApi(ctrl)
		postApi.EXPECT().GetPost(postId).Return(post, nil)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postController := &postControllerImpl{postHttpHelper: postHttpHelper, postApi: postApi}
		response, err := postController.GetPost(httpContext)

		assert.Equal(t, post, response)
		assert.Nil(t, err)
	})

	t.Run("GetPost:ParserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postController := &postControllerImpl{postHttpHelper: postHttpHelper}
		response, err := postController.GetPost(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetPost:ApiError", func(t *testing.T) {
		postId := new(models.PostId)
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postApi := mocks.NewMockPostApi(ctrl)
		postApi.EXPECT().GetPost(postId).Return(nil, systemErr)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postController := &postControllerImpl{postHttpHelper: postHttpHelper, postApi: postApi}
		response, err := postController.GetPost(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreatePost", func(t *testing.T) {
		post := new(models.Post)
		data := new(models.PostCreate)
		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postModelFactory.EXPECT().CreatePostCreate().Return(data)

		postApi := mocks.NewMockPostApi(ctrl)
		// TODO: Change any
		postApi.EXPECT().CreatePost(gomock.Any(), data).Return(post, nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		postController := &postControllerImpl{
			postModelFactory: postModelFactory,
			postApi:          postApi,
		}
		response, err := postController.CreatePost(httpContext)

		assert.Equal(t, post, response)
		assert.Nil(t, err)
	})

	t.Run("CreatePost:BindPostUpdateError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.PostCreate)

		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postModelFactory.EXPECT().CreatePostCreate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		postController := &postControllerImpl{
			postModelFactory: postModelFactory,
		}
		_, err := postController.CreatePost(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("CreatePost:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.PostCreate)

		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postModelFactory.EXPECT().CreatePostCreate().Return(data)

		postApi := mocks.NewMockPostApi(ctrl)
		postApi.EXPECT().CreatePost(gomock.Any(), data).Return(nil, systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		postController := &postControllerImpl{
			postModelFactory: postModelFactory,
			postApi:          postApi,
		}
		_, err := postController.CreatePost(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdatePost", func(t *testing.T) {
		postId := new(models.PostId)
		data := new(models.PostUpdate)
		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postModelFactory.EXPECT().CreatePostUpdate().Return(data)

		postApi := mocks.NewMockPostApi(ctrl)
		postApi.EXPECT().UpdatePost(postId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postController := &postControllerImpl{
			postHttpHelper:   postHttpHelper,
			postModelFactory: postModelFactory,
			postApi:          postApi,
		}
		_, err := postController.UpdatePost(httpContext)

		assert.Nil(t, err)
	})

	t.Run("UpdatePost:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postController := &postControllerImpl{postHttpHelper: postHttpHelper}
		_, err := postController.UpdatePost(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdatePost:BindPostUpdateError", func(t *testing.T) {
		postId := new(models.PostId)
		systemErr := common.NewUnknownError()
		data := new(models.PostUpdate)
		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postModelFactory.EXPECT().CreatePostUpdate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postController := &postControllerImpl{
			postHttpHelper:   postHttpHelper,
			postModelFactory: postModelFactory,
		}
		_, err := postController.UpdatePost(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdatePost:ApiError", func(t *testing.T) {
		postId := new(models.PostId)
		systemErr := common.NewUnknownError()

		data := new(models.PostUpdate)
		postModelFactory := mocks.NewMockPostModelFactory(ctrl)
		postModelFactory.EXPECT().CreatePostUpdate().Return(data)

		postApi := mocks.NewMockPostApi(ctrl)
		postApi.EXPECT().UpdatePost(postId, data).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postController := &postControllerImpl{
			postHttpHelper:   postHttpHelper,
			postModelFactory: postModelFactory,
			postApi:          postApi,
		}
		_, err := postController.UpdatePost(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeletePost", func(t *testing.T) {
		postId := new(models.PostId)

		postApi := mocks.NewMockPostApi(ctrl)
		postApi.EXPECT().DeletePost(postId).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postController := &postControllerImpl{postHttpHelper: postHttpHelper, postApi: postApi}
		_, err := postController.DeletePost(httpContext)

		assert.Nil(t, err)
	})

	t.Run("DeletePost:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postController := &postControllerImpl{postHttpHelper: postHttpHelper}
		_, err := postController.DeletePost(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeletePost:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		postId := new(models.PostId)

		postApi := mocks.NewMockPostApi(ctrl)
		postApi.EXPECT().DeletePost(postId).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		postController := &postControllerImpl{postHttpHelper: postHttpHelper, postApi: postApi}
		_, err := postController.DeletePost(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
