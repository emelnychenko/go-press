package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostCategoryService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostCategoryService", func(t *testing.T) {
		categoryService := mocks.NewMockCategoryService(ctrl)
		postCategoryService, isPostCategoryService := NewPostCategoryService(categoryService).(*postCategoryServiceImpl)

		assert.True(t, isPostCategoryService)
		assert.Equal(t, categoryService, postCategoryService.categoryService)
	})

	t.Run("ListPostCategories", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		categoryPaginationQuery := new(models.CategoryPaginationQuery)
		paginationResult := new(models.PaginationResult)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().ListObjectCategories(postEntity, categoryPaginationQuery).Return(paginationResult, nil)

		postCategoryService := &postCategoryServiceImpl{categoryService: categoryService}
		result, err := postCategoryService.ListPostCategories(postEntity, categoryPaginationQuery)

		assert.Nil(t, err)
		assert.Equal(t, result, paginationResult)
	})

	t.Run("AddPostCategory", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		categoryEntity := new(entities.CategoryEntity)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().CreateCategoryXref(categoryEntity, postEntity).Return(nil, nil)

		postCategoryService := &postCategoryServiceImpl{categoryService: categoryService}
		err := postCategoryService.AddPostCategory(postEntity, categoryEntity)

		assert.Nil(t, err)
	})

	t.Run("RemovePostCategory", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		categoryEntity := new(entities.CategoryEntity)
		categoryXrefEntity := new(entities.CategoryXrefEntity)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategoryXref(categoryEntity, postEntity).Return(categoryXrefEntity, nil)
		categoryService.EXPECT().DeleteCategoryXref(categoryXrefEntity).Return(nil)

		postCategoryService := &postCategoryServiceImpl{categoryService: categoryService}
		err := postCategoryService.RemovePostCategory(postEntity, categoryEntity)

		assert.Nil(t, err)
	})

	t.Run("RemovePostCategory:GetCategoryXrefError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		postEntity := new(entities.PostEntity)
		categoryEntity := new(entities.CategoryEntity)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategoryXref(categoryEntity, postEntity).Return(nil, systemErr)

		postCategoryService := &postCategoryServiceImpl{categoryService: categoryService}
		err := postCategoryService.RemovePostCategory(postEntity, categoryEntity)

		assert.Equal(t, systemErr, err)
	})
}
