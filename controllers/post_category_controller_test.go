package controllers

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostCategoryController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostCategoryController", func(t *testing.T) {
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		postCategoryApi := mocks.NewMockPostCategoryApi(ctrl)

		postCategoryController, isPostCategoryController := NewPostCategoryController(
			postHttpHelper,
			categoryHttpHelper,
			categoryModelFactory,
			postCategoryApi,
		).(*postCategoryControllerImpl)

		assert.True(t, isPostCategoryController)
		assert.Equal(t, postHttpHelper, postCategoryController.postHttpHelper)
		assert.Equal(t, categoryHttpHelper, postCategoryController.categoryHttpHelper)
		assert.Equal(t, categoryModelFactory, postCategoryController.categoryModelFactory)
		assert.Equal(t, postCategoryApi, postCategoryController.postCategoryApi)
	})

	t.Run("ListPostCategories", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		categoryPaginationQuery := new(models.CategoryPaginationQuery)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryPaginationQuery().Return(categoryPaginationQuery)

		httpContext.EXPECT().BindModel(categoryPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(categoryPaginationQuery).Return(nil)

		paginationResult := new(models.PaginationResult)
		postCategoryApi := mocks.NewMockPostCategoryApi(ctrl)
		postCategoryApi.EXPECT().ListPostCategories(postId, categoryPaginationQuery).
			Return(paginationResult, nil)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper:       postHttpHelper,
			categoryModelFactory: categoryModelFactory,
			postCategoryApi:      postCategoryApi,
		}
		result, err := postCategoryController.ListPostCategories(httpContext)
		assert.Equal(t, paginationResult, result)
		assert.Nil(t, err)
	})

	t.Run("ListPostCategories:ParsePostId", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		result, err := postCategoryController.ListPostCategories(httpContext)
		assert.Nil(t, result)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListPostCategories:BindPaginationQueryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		categoryPaginationQuery := new(models.CategoryPaginationQuery)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryPaginationQuery().Return(categoryPaginationQuery)

		httpContext.EXPECT().BindModel(categoryPaginationQuery.PaginationQuery).Return(systemErr)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper:       postHttpHelper,
			categoryModelFactory: categoryModelFactory,
		}
		result, err := postCategoryController.ListPostCategories(httpContext)
		assert.Nil(t, result)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListPostCategories:BindCategoryPaginationQueryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		categoryPaginationQuery := new(models.CategoryPaginationQuery)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryPaginationQuery().Return(categoryPaginationQuery)

		httpContext.EXPECT().BindModel(categoryPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(categoryPaginationQuery).Return(systemErr)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper:       postHttpHelper,
			categoryModelFactory: categoryModelFactory,
		}
		result, err := postCategoryController.ListPostCategories(httpContext)
		assert.Nil(t, result)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListPostCategories:ListPostCategoriesError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		categoryPaginationQuery := new(models.CategoryPaginationQuery)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryPaginationQuery().Return(categoryPaginationQuery)

		httpContext.EXPECT().BindModel(categoryPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(categoryPaginationQuery).Return(nil)

		postCategoryApi := mocks.NewMockPostCategoryApi(ctrl)
		postCategoryApi.EXPECT().ListPostCategories(postId, categoryPaginationQuery).
			Return(nil, systemErr)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper:       postHttpHelper,
			categoryModelFactory: categoryModelFactory,
			postCategoryApi:      postCategoryApi,
		}
		result, err := postCategoryController.ListPostCategories(httpContext)
		assert.Nil(t, result)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostCategory", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		categoryId := new(models.CategoryId)
		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		postCategoryApi := mocks.NewMockPostCategoryApi(ctrl)
		postCategoryApi.EXPECT().AddPostCategory(postId, categoryId).Return(nil)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper:     postHttpHelper,
			categoryHttpHelper: categoryHttpHelper,
			postCategoryApi:    postCategoryApi,
		}
		_, err := postCategoryController.AddPostCategory(httpContext)
		assert.Nil(t, err)
	})

	t.Run("AddPostCategory:ParsePostIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		_, err := postCategoryController.AddPostCategory(httpContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostCategory:ParseCategoryIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(nil, systemErr)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper:     postHttpHelper,
			categoryHttpHelper: categoryHttpHelper,
		}
		_, err := postCategoryController.AddPostCategory(httpContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("AddPostCategory:AddPostCategoryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		categoryId := new(models.CategoryId)
		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		postCategoryApi := mocks.NewMockPostCategoryApi(ctrl)
		postCategoryApi.EXPECT().AddPostCategory(postId, categoryId).Return(systemErr)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper:     postHttpHelper,
			categoryHttpHelper: categoryHttpHelper,
			postCategoryApi:    postCategoryApi,
		}
		_, err := postCategoryController.AddPostCategory(httpContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostCategory", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		categoryId := new(models.CategoryId)
		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		postCategoryApi := mocks.NewMockPostCategoryApi(ctrl)
		postCategoryApi.EXPECT().RemovePostCategory(postId, categoryId).Return(nil)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper:     postHttpHelper,
			categoryHttpHelper: categoryHttpHelper,
			postCategoryApi:    postCategoryApi,
		}
		_, err := postCategoryController.RemovePostCategory(httpContext)
		assert.Nil(t, err)
	})

	t.Run("RemovePostCategory:ParsePostIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(nil, systemErr)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper: postHttpHelper,
		}
		_, err := postCategoryController.RemovePostCategory(httpContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostCategory:ParseCategoryIdError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(nil, systemErr)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper:     postHttpHelper,
			categoryHttpHelper: categoryHttpHelper,
		}
		_, err := postCategoryController.RemovePostCategory(httpContext)
		assert.Equal(t, systemErr, err)
	})

	t.Run("RemovePostCategory:RemovePostCategoryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		postId := new(models.PostId)
		postHttpHelper := mocks.NewMockPostHttpHelper(ctrl)
		postHttpHelper.EXPECT().ParsePostId(httpContext).Return(postId, nil)

		categoryId := new(models.CategoryId)
		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		postCategoryApi := mocks.NewMockPostCategoryApi(ctrl)
		postCategoryApi.EXPECT().RemovePostCategory(postId, categoryId).Return(systemErr)

		postCategoryController := &postCategoryControllerImpl{
			postHttpHelper:     postHttpHelper,
			categoryHttpHelper: categoryHttpHelper,
			postCategoryApi:    postCategoryApi,
		}
		_, err := postCategoryController.RemovePostCategory(httpContext)
		assert.Equal(t, systemErr, err)
	})
}
