package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCategoryService", func(t *testing.T) {
		categoryEntityFactory := mocks.NewMockCategoryEntityFactory(ctrl)
		categoryRepository := mocks.NewMockCategoryRepository(ctrl)

		categoryService, isCategoryService := NewCategoryService(
			categoryEntityFactory, categoryRepository,
		).(*categoryServiceImpl)

		assert.True(t, isCategoryService)
		assert.Equal(t, categoryEntityFactory, categoryService.categoryEntityFactory)
		assert.Equal(t, categoryRepository, categoryService.categoryRepository)
	})

	t.Run("ListCategories", func(t *testing.T) {
		categoryPaginationQuery := new(models.CategoryPaginationQuery)

		var categoryEntities *models.PaginationResult
		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().ListCategories(categoryPaginationQuery).Return(categoryEntities, nil)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		response, err := categoryService.ListCategories(categoryPaginationQuery)

		assert.Equal(t, categoryEntities, response)
		assert.Nil(t, err)
	})

	t.Run("CreateCategory", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryEntityFactory := mocks.NewMockCategoryEntityFactory(ctrl)
		categoryEntityFactory.EXPECT().CreateCategoryEntity().Return(categoryEntity)

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().SaveCategory(categoryEntity).Return(nil)

		data := &models.CategoryCreate{
			Name: "0",
		}
		categoryService := &categoryServiceImpl{
			categoryEntityFactory: categoryEntityFactory,
			categoryRepository:    categoryRepository,
		}
		response, err := categoryService.CreateCategory(data)

		assert.IsType(t, categoryEntity, response)
		assert.Nil(t, err)
		assert.Equal(t, data.Name, categoryEntity.Name)
	})

	t.Run("GetCategory", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)
		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategory(categoryId).Return(categoryEntity, nil)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		response, err := categoryService.GetCategory(categoryId)

		assert.Equal(t, categoryEntity, response)
		assert.Nil(t, err)
	})

	t.Run("UpdateCategory", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().SaveCategory(categoryEntity).Return(nil)

		data := &models.CategoryUpdate{
			Name: "0",
		}
		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		assert.Nil(t, categoryService.UpdateCategory(categoryEntity, data))

		assert.Equal(t, data.Name, categoryEntity.Name)
		assert.NotNil(t, categoryEntity.Updated)
	})

	t.Run("DeleteCategory", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().RemoveCategory(categoryEntity).Return(nil)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		assert.Nil(t, categoryService.DeleteCategory(categoryEntity))
	})
}
