package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCategoryController", func(t *testing.T) {
		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryController, isCategoryController := NewCategoryController(
			categoryHttpHelper,
			categoryModelFactory,
			categoryApi,
		).(*categoryControllerImpl)

		assert.True(t, isCategoryController)
		assert.Equal(t, categoryHttpHelper, categoryController.categoryHttpHelper)
		assert.Equal(t, categoryModelFactory, categoryController.categoryModelFactory)
		assert.Equal(t, categoryApi, categoryController.categoryApi)
	})

	t.Run("ListCategories", func(t *testing.T) {
		categoryPaginationQuery := new(models.CategoryPaginationQuery)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryPaginationQuery().Return(categoryPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(categoryPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(categoryPaginationQuery).Return(nil)

		var paginationResult *models.PaginationResult
		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().ListCategories(categoryPaginationQuery).Return(paginationResult, nil)

		categoryController := &categoryControllerImpl{
			categoryModelFactory: categoryModelFactory,
			categoryApi:          categoryApi,
		}
		response, err := categoryController.ListCategories(httpContext)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListCategories:BindPaginationQueryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		categoryPaginationQuery := new(models.CategoryPaginationQuery)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryPaginationQuery().Return(categoryPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(categoryPaginationQuery.PaginationQuery).Return(systemErr)

		categoryController := &categoryControllerImpl{
			categoryModelFactory: categoryModelFactory,
		}
		response, err := categoryController.ListCategories(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListCategories:BindCategoryPaginationQueryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		categoryPaginationQuery := new(models.CategoryPaginationQuery)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryPaginationQuery().Return(categoryPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(categoryPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(categoryPaginationQuery).Return(systemErr)

		categoryController := &categoryControllerImpl{
			categoryModelFactory: categoryModelFactory,
		}
		response, err := categoryController.ListCategories(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetCategoriesTree", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		var categoriesTree []*models.CategoryTree
		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().GetCategoriesTree().Return(categoriesTree, nil)

		categoryController := &categoryControllerImpl{categoryApi: categoryApi}
		response, err := categoryController.GetCategoriesTree(httpContext)

		assert.Equal(t, categoriesTree, response)
		assert.Nil(t, err)
	})

	t.Run("GetCategory", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		var category *models.Category
		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().GetCategory(categoryId).Return(category, nil)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryController := &categoryControllerImpl{categoryHttpHelper: categoryHttpHelper, categoryApi: categoryApi}
		response, err := categoryController.GetCategory(httpContext)

		assert.Equal(t, category, response)
		assert.Nil(t, err)
	})

	t.Run("GetCategory:ParserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(nil, systemErr)

		categoryController := &categoryControllerImpl{categoryHttpHelper: categoryHttpHelper}
		response, err := categoryController.GetCategory(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetCategory:ApiError", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().GetCategory(categoryId).Return(nil, systemErr)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryController := &categoryControllerImpl{categoryHttpHelper: categoryHttpHelper, categoryApi: categoryApi}
		response, err := categoryController.GetCategory(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetCategoryTree", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		var categoryTree *models.CategoryTree
		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().GetCategoryTree(categoryId).Return(categoryTree, nil)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryController := &categoryControllerImpl{categoryHttpHelper: categoryHttpHelper, categoryApi: categoryApi}
		response, err := categoryController.GetCategoryTree(httpContext)

		assert.Equal(t, categoryTree, response)
		assert.Nil(t, err)
	})

	t.Run("GetCategoryTree:ParserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(nil, systemErr)

		categoryController := &categoryControllerImpl{categoryHttpHelper: categoryHttpHelper}
		response, err := categoryController.GetCategoryTree(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetCategoryTree:ApiError", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().GetCategoryTree(categoryId).Return(nil, systemErr)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryController := &categoryControllerImpl{categoryHttpHelper: categoryHttpHelper, categoryApi: categoryApi}
		response, err := categoryController.GetCategoryTree(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateCategory", func(t *testing.T) {
		category := new(models.Category)
		data := new(models.CategoryCreate)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryCreate().Return(data)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().CreateCategory(data).Return(category, nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		categoryController := &categoryControllerImpl{
			categoryModelFactory: categoryModelFactory,
			categoryApi:          categoryApi,
		}
		response, err := categoryController.CreateCategory(httpContext)

		assert.Equal(t, category, response)
		assert.Nil(t, err)
	})

	t.Run("CreateCategory:BindCategoryUpdateError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.CategoryCreate)

		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryCreate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		categoryController := &categoryControllerImpl{
			categoryModelFactory: categoryModelFactory,
		}
		_, err := categoryController.CreateCategory(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateCategory:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.CategoryCreate)

		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryCreate().Return(data)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().CreateCategory(data).Return(nil, systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		categoryController := &categoryControllerImpl{
			categoryModelFactory: categoryModelFactory,
			categoryApi:          categoryApi,
		}
		_, err := categoryController.CreateCategory(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateCategory", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		data := new(models.CategoryUpdate)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryUpdate().Return(data)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().UpdateCategory(categoryId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryController := &categoryControllerImpl{
			categoryHttpHelper:   categoryHttpHelper,
			categoryModelFactory: categoryModelFactory,
			categoryApi:          categoryApi,
		}
		_, err := categoryController.UpdateCategory(httpContext)

		assert.Nil(t, err)
	})

	t.Run("UpdateCategory:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(nil, systemErr)

		categoryController := &categoryControllerImpl{categoryHttpHelper: categoryHttpHelper}
		_, err := categoryController.UpdateCategory(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateCategory:BindCategoryUpdateError", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		systemErr := common.NewUnknownError()
		data := new(models.CategoryUpdate)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryUpdate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryController := &categoryControllerImpl{
			categoryHttpHelper:   categoryHttpHelper,
			categoryModelFactory: categoryModelFactory,
		}
		_, err := categoryController.UpdateCategory(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateCategory:ApiError", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		systemErr := common.NewUnknownError()

		data := new(models.CategoryUpdate)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategoryUpdate().Return(data)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().UpdateCategory(categoryId, data).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryController := &categoryControllerImpl{
			categoryHttpHelper:   categoryHttpHelper,
			categoryModelFactory: categoryModelFactory,
			categoryApi:          categoryApi,
		}
		_, err := categoryController.UpdateCategory(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeCategoryParent", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		parentCategoryId := new(models.CategoryId)

		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)
		categoryHttpHelper.EXPECT().ParseParentCategoryId(httpContext).Return(parentCategoryId, nil)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().ChangeCategoryParent(categoryId, parentCategoryId).Return(nil)

		categoryController := &categoryControllerImpl{
			categoryHttpHelper: categoryHttpHelper,
			categoryApi:        categoryApi,
		}
		_, err := categoryController.ChangeCategoryParent(httpContext)

		assert.Nil(t, err)
	})

	t.Run("ChangeCategoryParent:ParseCategoryIdError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(nil, systemErr)

		categoryController := &categoryControllerImpl{
			categoryHttpHelper: categoryHttpHelper,
		}
		_, err := categoryController.ChangeCategoryParent(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeCategoryParent:ParseParentCategoryIdError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		categoryId := new(models.CategoryId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)
		categoryHttpHelper.EXPECT().ParseParentCategoryId(httpContext).Return(nil, systemErr)

		categoryController := &categoryControllerImpl{
			categoryHttpHelper: categoryHttpHelper,
		}
		_, err := categoryController.ChangeCategoryParent(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeCategoryParent:ApiChangeCategoryParentError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		categoryId := new(models.CategoryId)
		parentCategoryId := new(models.CategoryId)

		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)
		categoryHttpHelper.EXPECT().ParseParentCategoryId(httpContext).Return(parentCategoryId, nil)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().ChangeCategoryParent(categoryId, parentCategoryId).Return(systemErr)

		categoryController := &categoryControllerImpl{
			categoryHttpHelper: categoryHttpHelper,
			categoryApi:        categoryApi,
		}
		_, err := categoryController.ChangeCategoryParent(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("RemoveCategoryParent", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().RemoveCategoryParent(categoryId).Return(nil)

		categoryController := &categoryControllerImpl{
			categoryHttpHelper: categoryHttpHelper,
			categoryApi:        categoryApi,
		}
		_, err := categoryController.RemoveCategoryParent(httpContext)

		assert.Nil(t, err)
	})

	t.Run("RemoveCategoryParent:ParseCategoryIdError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(nil, systemErr)

		categoryController := &categoryControllerImpl{
			categoryHttpHelper: categoryHttpHelper,
		}
		_, err := categoryController.RemoveCategoryParent(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("RemoveCategoryParent:ApiRemoveCategoryParentError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		categoryId := new(models.CategoryId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().RemoveCategoryParent(categoryId).Return(systemErr)

		categoryController := &categoryControllerImpl{
			categoryHttpHelper: categoryHttpHelper,
			categoryApi:        categoryApi,
		}
		_, err := categoryController.RemoveCategoryParent(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteCategory", func(t *testing.T) {
		categoryId := new(models.CategoryId)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().DeleteCategory(categoryId).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryController := &categoryControllerImpl{categoryHttpHelper: categoryHttpHelper, categoryApi: categoryApi}
		_, err := categoryController.DeleteCategory(httpContext)

		assert.Nil(t, err)
	})

	t.Run("DeleteCategory:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(nil, systemErr)

		categoryController := &categoryControllerImpl{categoryHttpHelper: categoryHttpHelper}
		_, err := categoryController.DeleteCategory(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteCategory:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		categoryId := new(models.CategoryId)

		categoryApi := mocks.NewMockCategoryApi(ctrl)
		categoryApi.EXPECT().DeleteCategory(categoryId).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)

		categoryHttpHelper := mocks.NewMockCategoryHttpHelper(ctrl)
		categoryHttpHelper.EXPECT().ParseCategoryId(httpContext).Return(categoryId, nil)

		categoryController := &categoryControllerImpl{categoryHttpHelper: categoryHttpHelper, categoryApi: categoryApi}
		_, err := categoryController.DeleteCategory(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
