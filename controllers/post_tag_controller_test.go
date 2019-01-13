package controllers

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostTagController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostTagController", func(t *testing.T) {
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		postTagApi := mocks.NewMockPostTagApi(ctrl)

		postTagController, isPostTagController := NewPostTagController(
			postHttpHelper,
			tagHttpHelper,
			tagModelFactory,
			postTagApi,
		).(*postTagControllerImpl)

		assert.True(t, isPostTagController)
		assert.Equal(t, postHttpHelper, postTagController.postHttpHelper)
		assert.Equal(t, tagHttpHelper, postTagController.tagHttpHelper)
		assert.Equal(t, tagModelFactory, postTagController.tagModelFactory)
		assert.Equal(t, postTagApi, postTagController.postTagApi)
	})

	t.Run("ListPostTags", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		tagPaginationQuery := new(models.TagPaginationQuery)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagPaginationQuery().Return(tagPaginationQuery)

		httpContext.EXPECT().BindModel(tagPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(tagPaginationQuery).Return(nil)

		paginationResult := new(models.PaginationResult)
		postTagApi := mocks.NewMockPostTagApi(ctrl)
		postTagApi.EXPECT().ListPostTags(postId, tagPaginationQuery).
			Return(paginationResult, nil)

		postTagController := &postTagControllerImpl{
			postHttpHelper:       postHttpHelper,
			tagModelFactory: tagModelFactory,
			postTagApi:      postTagApi,
		}
		result, err := postTagController.ListPostTags(httpContext)
		assert.Equal(t, paginationResult, result)
		assert.Nil(t, err)
	})

	t.Run("ListPostTags:ParsePostId", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postTagController := &postTagControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		result, err := postTagController.ListPostTags(httpContext)
		assert.Nil(t, result)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListPostTags:BindPaginationQueryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		tagPaginationQuery := new(models.TagPaginationQuery)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagPaginationQuery().Return(tagPaginationQuery)

		httpContext.EXPECT().BindModel(tagPaginationQuery.PaginationQuery).Return(systemErr)

		postTagController := &postTagControllerImpl{
			postHttpHelper:       postHttpHelper,
			tagModelFactory: tagModelFactory,
		}
		result, err := postTagController.ListPostTags(httpContext)
		assert.Nil(t, result)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListPostTags:BindTagPaginationQueryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		tagPaginationQuery := new(models.TagPaginationQuery)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagPaginationQuery().Return(tagPaginationQuery)

		httpContext.EXPECT().BindModel(tagPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(tagPaginationQuery).Return(systemErr)

		postTagController := &postTagControllerImpl{
			postHttpHelper:       postHttpHelper,
			tagModelFactory: tagModelFactory,
		}
		result, err := postTagController.ListPostTags(httpContext)
		assert.Nil(t, result)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListPostTags:ListPostTagsError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		tagPaginationQuery := new(models.TagPaginationQuery)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTagPaginationQuery().Return(tagPaginationQuery)

		httpContext.EXPECT().BindModel(tagPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(tagPaginationQuery).Return(nil)

		postTagApi := mocks.NewMockPostTagApi(ctrl)
		postTagApi.EXPECT().ListPostTags(postId, tagPaginationQuery).
			Return(nil, systemErr)

		postTagController := &postTagControllerImpl{
			postHttpHelper:       postHttpHelper,
			tagModelFactory: tagModelFactory,
			postTagApi:      postTagApi,
		}
		result, err := postTagController.ListPostTags(httpContext)
		assert.Nil(t, result)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostTag", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		tagId := new(models.TagId)
		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		postTagApi := mocks.NewMockPostTagApi(ctrl)
		postTagApi.EXPECT().AddPostTag(postId, tagId).Return(nil)

		postTagController := &postTagControllerImpl{
			postHttpHelper:     postHttpHelper,
			tagHttpHelper: tagHttpHelper,
			postTagApi:    postTagApi,
		}
		_, err := postTagController.AddPostTag(httpContext)
		assert.Nil(t, err)
	})

	t.Run("AddPostTag:ParsePostIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postTagController := &postTagControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		_, err := postTagController.AddPostTag(httpContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostTag:ParseTagIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(nil, systemErr)

		postTagController := &postTagControllerImpl{
			postHttpHelper:     postHttpHelper,
			tagHttpHelper: tagHttpHelper,
		}
		_, err := postTagController.AddPostTag(httpContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostTag:AddPostTagError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		tagId := new(models.TagId)
		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		postTagApi := mocks.NewMockPostTagApi(ctrl)
		postTagApi.EXPECT().AddPostTag(postId, tagId).Return(systemErr)

		postTagController := &postTagControllerImpl{
			postHttpHelper:     postHttpHelper,
			tagHttpHelper: tagHttpHelper,
			postTagApi:    postTagApi,
		}
		_, err := postTagController.AddPostTag(httpContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostTag", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		tagId := new(models.TagId)
		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		postTagApi := mocks.NewMockPostTagApi(ctrl)
		postTagApi.EXPECT().RemovePostTag(postId, tagId).Return(nil)

		postTagController := &postTagControllerImpl{
			postHttpHelper:     postHttpHelper,
			tagHttpHelper: tagHttpHelper,
			postTagApi:    postTagApi,
		}
		_, err := postTagController.RemovePostTag(httpContext)
		assert.Nil(t, err)
	})

	t.Run("RemovePostTag:ParsePostIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postTagController := &postTagControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		_, err := postTagController.RemovePostTag(httpContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostTag:ParseTagIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(nil, systemErr)

		postTagController := &postTagControllerImpl{
			postHttpHelper:     postHttpHelper,
			tagHttpHelper: tagHttpHelper,
		}
		_, err := postTagController.RemovePostTag(httpContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostTag:RemovePostTagError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		tagId := new(models.TagId)
		tagHttpHelper := mocks.NewMockTagHttpHelper(ctrl)
		tagHttpHelper.EXPECT().ParseTagId(httpContext).Return(tagId, nil)

		postTagApi := mocks.NewMockPostTagApi(ctrl)
		postTagApi.EXPECT().RemovePostTag(postId, tagId).Return(systemErr)

		postTagController := &postTagControllerImpl{
			postHttpHelper:     postHttpHelper,
			tagHttpHelper: tagHttpHelper,
			postTagApi:    postTagApi,
		}
		_, err := postTagController.RemovePostTag(httpContext)
		assert.Equal(t, systemErr, err)
	})
}
