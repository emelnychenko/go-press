package jobs

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryEdgesBuilderJob(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCategoryEdgesBuilderJob", func(t *testing.T) {
		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryNestedSetBuilder := mocks.NewMockCategoryNestedSetBuilder(ctrl)

		categoryEdgesBuilderJob, isCategoryEdgesBuilderJob := NewCategoryEdgesBuilderJob(
			categoryRepository,
			categoryNestedSetBuilder,
		).(*categoryEdgesBuilderJobImpl)

		assert.True(t, isCategoryEdgesBuilderJob)
		assert.Equal(t, categoryRepository, categoryEdgesBuilderJob.categoryRepository)
		assert.Equal(t, categoryNestedSetBuilder, categoryEdgesBuilderJob.categoryNestedSetBuilder)
	})

	t.Run("BuildCategoriesEdges", func(t *testing.T) {
		categoryEntity := &entities.CategoryEntity{Left: 1, Right: 2}
		categoryEntities := []*entities.CategoryEntity{categoryEntity}

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategories().Return(categoryEntities, nil)
		categoryRepository.EXPECT().SaveCategory(categoryEntity).Return(nil)

		categoryEntityNestedSetNode := &entities.CategoryEntityNestedSetNode{
			Left: 1, Right: 4, CategoryEntity: categoryEntity,
		}
		categoryEntityNestedSetNodes := []*entities.CategoryEntityNestedSetNode{categoryEntityNestedSetNode}
		categoryEntityNestedSet := &entities.CategoryEntityNestedSet{Nodes: categoryEntityNestedSetNodes}

		categoryNestedSetBuilder := mocks.NewMockCategoryNestedSetBuilder(ctrl)
		categoryNestedSetBuilder.EXPECT().BuildCategoryEntityNestedSet(categoryEntities).
			Return(categoryEntityNestedSet, nil)

		categoryEdgesBuilderJob := &categoryEdgesBuilderJobImpl{
			categoryRepository:       categoryRepository,
			categoryNestedSetBuilder: categoryNestedSetBuilder,
		}
		err := categoryEdgesBuilderJob.BuildCategoriesEdges()

		assert.Nil(t, err)
		assert.Equal(t, 1, categoryEntity.Left)
		assert.Equal(t, 4, categoryEntity.Right)
	})

	t.Run("BuildCategoriesEdges:GetCategoriesError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		categoryEntity := &entities.CategoryEntity{Left: 1, Right: 2}
		categoryEntities := []*entities.CategoryEntity{categoryEntity}

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategories().Return(categoryEntities, systemErr)

		categoryEdgesBuilderJob := &categoryEdgesBuilderJobImpl{
			categoryRepository: categoryRepository,
		}
		err := categoryEdgesBuilderJob.BuildCategoriesEdges()

		assert.Equal(t, systemErr, err)
	})

	t.Run("BuildCategoriesEdges:BuildCategoryEntityNestedSetError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		categoryEntity := &entities.CategoryEntity{Left: 1, Right: 2}
		categoryEntities := []*entities.CategoryEntity{categoryEntity}

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategories().Return(categoryEntities, nil)

		categoryNestedSetBuilder := mocks.NewMockCategoryNestedSetBuilder(ctrl)
		categoryNestedSetBuilder.EXPECT().BuildCategoryEntityNestedSet(categoryEntities).Return(nil, systemErr)

		categoryEdgesBuilderJob := &categoryEdgesBuilderJobImpl{
			categoryRepository:       categoryRepository,
			categoryNestedSetBuilder: categoryNestedSetBuilder,
		}
		err := categoryEdgesBuilderJob.BuildCategoriesEdges()

		assert.Equal(t, systemErr, err)
	})

	t.Run("BuildCategoriesEdges:SaveCategoryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		categoryEntity := &entities.CategoryEntity{Left: 1, Right: 2}
		categoryEntities := []*entities.CategoryEntity{categoryEntity}

		categoryRepository := mocks.NewMockCategoryRepository(ctrl)
		categoryRepository.EXPECT().GetCategories().Return(categoryEntities, nil)
		categoryRepository.EXPECT().SaveCategory(categoryEntity).Return(systemErr)

		categoryEntityNestedSetNode := &entities.CategoryEntityNestedSetNode{
			Left: 1, Right: 4, CategoryEntity: categoryEntity,
		}
		categoryEntityNestedSetNodes := []*entities.CategoryEntityNestedSetNode{categoryEntityNestedSetNode}
		categoryEntityNestedSet := &entities.CategoryEntityNestedSet{Nodes: categoryEntityNestedSetNodes}

		categoryNestedSetBuilder := mocks.NewMockCategoryNestedSetBuilder(ctrl)
		categoryNestedSetBuilder.EXPECT().BuildCategoryEntityNestedSet(categoryEntities).
			Return(categoryEntityNestedSet, nil)

		categoryEdgesBuilderJob := &categoryEdgesBuilderJobImpl{
			categoryRepository:       categoryRepository,
			categoryNestedSetBuilder: categoryNestedSetBuilder,
		}
		err := categoryEdgesBuilderJob.BuildCategoriesEdges()

		assert.Equal(t, systemErr, err)
	})
}
