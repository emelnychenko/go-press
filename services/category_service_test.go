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

func TestCategoryService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCategoryService", func(t *testing.T) {
		categoryEntityFactory := mocks.NewMockCategoryEntityFactory(ctrl)
		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryTreeBuilder := mocks.NewMockCategoryTreeBuilder(ctrl)
		categoryEdgesBuilderJob := mocks.NewMockCategoryEdgesBuilderJob(ctrl)

		categoryService, isCategoryService := NewCategoryService(
			categoryEntityFactory, categoryRepository, categoryTreeBuilder, categoryEdgesBuilderJob,
		).(*categoryServiceImpl)

		assert.True(t, isCategoryService)
		assert.Equal(t, categoryEntityFactory, categoryService.categoryEntityFactory)
		assert.Equal(t, categoryRepository, categoryService.categoryRepository)
		assert.Equal(t, categoryTreeBuilder, categoryService.categoryTreeBuilder)
		assert.Equal(t, categoryEdgesBuilderJob, categoryService.categoryEdgesBuilderJob)
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

	t.Run("GetCategoriesTree", func(t *testing.T) {
		categoryEntityTree := new(entities.CategoryEntityTree)
		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategoriesTree().Return(categoryEntityTree, nil)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		result, err := categoryService.GetCategoriesTree()

		assert.Equal(t, categoryEntityTree, result)
		assert.Nil(t, err)
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

	t.Run("GetCategoryTree", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		categoryEntityTree := new(entities.CategoryEntityTree)
		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategoryTree(categoryId).Return(categoryEntityTree, nil)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		result, err := categoryService.GetCategoryTree(categoryId)

		assert.Equal(t, categoryEntityTree, result)
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

	t.Run("ChangeCategoryParent", func(t *testing.T) {
		parentCategoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)
		parentCategoryEntity := &entities.CategoryEntity{Id: parentCategoryId}
		var categoryEntities []*entities.CategoryEntity

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategoriesExcept(categoryEntity).Return(categoryEntities, nil)
		categoryRepository.EXPECT().SaveCategory(categoryEntity).Return(nil)

		categoryEntities = append(categoryEntities, categoryEntity)
		categoryTreeBuilder := mocks.NewMockCategoryTreeBuilder(ctrl)
		categoryTreeBuilder.EXPECT().BuildCategoryEntityTree(categoryEntities).Return(nil, nil)

		categoryEdgesBuilderJob := mocks.NewMockCategoryEdgesBuilderJob(ctrl)
		categoryEdgesBuilderJob.EXPECT().BuildCategoriesEdges().Return(nil)

		categoryService := &categoryServiceImpl{
			categoryRepository:      categoryRepository,
			categoryTreeBuilder:     categoryTreeBuilder,
			categoryEdgesBuilderJob: categoryEdgesBuilderJob,
		}
		err := categoryService.ChangeCategoryParent(categoryEntity, parentCategoryEntity)

		assert.Nil(t, err)
		assert.Equal(t, parentCategoryId, categoryEntity.ParentCategoryId)
	})

	t.Run("ChangeCategoryParent:GetCategoriesExceptError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		parentCategoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)
		parentCategoryEntity := &entities.CategoryEntity{Id: parentCategoryId}

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategoriesExcept(categoryEntity).Return(nil, systemErr)

		categoryService := &categoryServiceImpl{
			categoryRepository: categoryRepository,
		}
		err := categoryService.ChangeCategoryParent(categoryEntity, parentCategoryEntity)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeCategoryParent:BuildCategoryEntityTreeError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		parentCategoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)
		parentCategoryEntity := &entities.CategoryEntity{Id: parentCategoryId}
		var categoryEntities []*entities.CategoryEntity

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategoriesExcept(categoryEntity).Return(categoryEntities, nil)

		categoryEntities = append(categoryEntities, categoryEntity)
		categoryTreeBuilder := mocks.NewMockCategoryTreeBuilder(ctrl)
		categoryTreeBuilder.EXPECT().BuildCategoryEntityTree(categoryEntities).Return(nil, systemErr)

		categoryService := &categoryServiceImpl{
			categoryRepository:  categoryRepository,
			categoryTreeBuilder: categoryTreeBuilder,
		}
		err := categoryService.ChangeCategoryParent(categoryEntity, parentCategoryEntity)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeCategoryParent:SaveCategoryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		parentCategoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)
		parentCategoryEntity := &entities.CategoryEntity{Id: parentCategoryId}
		var categoryEntities []*entities.CategoryEntity

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategoriesExcept(categoryEntity).Return(categoryEntities, nil)
		categoryRepository.EXPECT().SaveCategory(categoryEntity).Return(systemErr)

		categoryEntities = append(categoryEntities, categoryEntity)
		categoryTreeBuilder := mocks.NewMockCategoryTreeBuilder(ctrl)
		categoryTreeBuilder.EXPECT().BuildCategoryEntityTree(categoryEntities).Return(nil, nil)

		categoryService := &categoryServiceImpl{
			categoryRepository:  categoryRepository,
			categoryTreeBuilder: categoryTreeBuilder,
		}
		err := categoryService.ChangeCategoryParent(categoryEntity, parentCategoryEntity)

		assert.Equal(t, systemErr, err)
	})

	t.Run("RemoveCategoryParent", func(t *testing.T) {
		categoryEntity := &entities.CategoryEntity{ParentCategoryId: new(models.CategoryId)}

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().SaveCategory(categoryEntity).Return(nil)

		categoryEdgesBuilderJob := mocks.NewMockCategoryEdgesBuilderJob(ctrl)
		categoryEdgesBuilderJob.EXPECT().BuildCategoriesEdges().Return(nil)

		categoryService := &categoryServiceImpl{
			categoryRepository:      categoryRepository,
			categoryEdgesBuilderJob: categoryEdgesBuilderJob,
		}
		err := categoryService.RemoveCategoryParent(categoryEntity)

		assert.Nil(t, err, categoryEntity.ParentCategoryId)
	})

	t.Run("RemoveCategoryParent:SaveCategoryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		categoryEntity := &entities.CategoryEntity{ParentCategoryId: new(models.CategoryId)}

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().SaveCategory(categoryEntity).Return(systemErr)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		err := categoryService.RemoveCategoryParent(categoryEntity)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteCategory", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().RemoveCategory(categoryEntity).Return(nil)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		assert.Nil(t, categoryService.DeleteCategory(categoryEntity))
	})

	t.Run("GetCategoryXrefs", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		var categoryXrefEntities []*entities.CategoryXrefEntity

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategoryXrefs(categoryEntity).Return(categoryXrefEntities, nil)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		results, err := categoryService.GetCategoryXrefs(categoryEntity)

		assert.Equal(t, categoryXrefEntities, results)
		assert.Nil(t, err)
	})

	t.Run("GetCategoryObjectXrefs", func(t *testing.T) {
		postEntity := new(entities.PostEntity)
		var categoryXrefEntities []*entities.CategoryXrefEntity

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategoryObjectXrefs(postEntity).Return(categoryXrefEntities, nil)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		results, err := categoryService.GetCategoryObjectXrefs(postEntity)

		assert.Equal(t, categoryXrefEntities, results)
		assert.Nil(t, err)
	})

	t.Run("GetCategoryXref", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		postEntity := new(entities.PostEntity)
		categoryXrefEntity := new(entities.CategoryXrefEntity)

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategoryXref(categoryEntity, postEntity).Return(categoryXrefEntity, nil)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		result, err := categoryService.GetCategoryXref(categoryEntity, postEntity)

		assert.Equal(t, categoryXrefEntity, result)
		assert.Nil(t, err)
	})

	t.Run("CreateCategoryXref", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		postEntity := new(entities.PostEntity)
		categoryXrefEntity := new(entities.CategoryXrefEntity)

		categoryEntityFactory := mocks.NewMockCategoryEntityFactory(ctrl)
		categoryEntityFactory.EXPECT().CreateCategoryXrefEntity(categoryEntity, postEntity).Return(categoryXrefEntity)

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().SaveCategoryXref(categoryXrefEntity).Return(nil)

		categoryService := &categoryServiceImpl{
			categoryRepository: categoryRepository, categoryEntityFactory: categoryEntityFactory,
		}
		result, err := categoryService.CreateCategoryXref(categoryEntity, postEntity)

		assert.Equal(t, categoryXrefEntity, result)
		assert.Nil(t, err)
	})

	t.Run("DeleteCategoryXref", func(t *testing.T) {
		categoryXrefEntity := new(entities.CategoryXrefEntity)

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().RemoveCategoryXref(categoryXrefEntity).Return(nil)

		categoryService := &categoryServiceImpl{categoryRepository: categoryRepository}
		err := categoryService.DeleteCategoryXref(categoryXrefEntity)

		assert.Nil(t, err)
	})
}
