package aggregators

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryAggregator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCategoryAggregator", func(t *testing.T) {
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryAggregator, isCategoryAggregator := NewCategoryAggregator(
			categoryModelFactory,
		).(*categoryAggregatorImpl)

		assert.True(t, isCategoryAggregator)
		assert.Equal(t, categoryModelFactory, categoryAggregator.categoryModelFactory)
	})

	t.Run("AggregateCategory", func(t *testing.T) {
		category := new(models.Category)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategory().Return(category)

		categoryAggregator := &categoryAggregatorImpl{categoryModelFactory: categoryModelFactory}
		response := categoryAggregator.AggregateCategory(new(entities.CategoryEntity))

		assert.Equal(t, category, response)
	})

	t.Run("AggregateCategoryTree", func(t *testing.T) {
		category0 := new(models.Category)
		category0category0 := new(models.Category)
		category0category1 := new(models.Category)
		category0category1category0 := new(models.Category)
		category0category1category1 := new(models.Category)

		categoryTree0 := new(models.CategoryTree)
		categoryTree0categoryTree0 := new(models.CategoryTree)
		categoryTree0categoryTree1 := new(models.CategoryTree)
		categoryTree0categoryTree1categoryTree0 := new(models.CategoryTree)
		categoryTree0categoryTree1categoryTree1 := new(models.CategoryTree)

		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)

		categoryModelFactory.EXPECT().CreateCategory().Return(category0)
		categoryModelFactory.EXPECT().CreateCategory().Return(category0category0)
		categoryModelFactory.EXPECT().CreateCategory().Return(category0category1)
		categoryModelFactory.EXPECT().CreateCategory().Return(category0category1category0)
		categoryModelFactory.EXPECT().CreateCategory().Return(category0category1category1)

		categoryModelFactory.EXPECT().CreateCategoryTree().Return(categoryTree0)
		categoryModelFactory.EXPECT().CreateCategoryTree().Return(categoryTree0categoryTree0)
		categoryModelFactory.EXPECT().CreateCategoryTree().Return(categoryTree0categoryTree1)
		categoryModelFactory.EXPECT().CreateCategoryTree().Return(categoryTree0categoryTree1categoryTree0)
		categoryModelFactory.EXPECT().CreateCategoryTree().Return(categoryTree0categoryTree1categoryTree1)

		categoryEntityTree := &entities.CategoryEntityTree{
			Roots: []*entities.CategoryEntityTreeBranch{
				{
					CategoryEntity: new(entities.CategoryEntity),
					Children: []*entities.CategoryEntityTreeBranch{
						{
							CategoryEntity: new(entities.CategoryEntity),
							Children:       []*entities.CategoryEntityTreeBranch{},
						},
						{
							CategoryEntity: new(entities.CategoryEntity),
							Children: []*entities.CategoryEntityTreeBranch{
								{
									CategoryEntity: new(entities.CategoryEntity),
									Children:       []*entities.CategoryEntityTreeBranch{},
								},
								{
									CategoryEntity: new(entities.CategoryEntity),
									Children:       []*entities.CategoryEntityTreeBranch{},
								},
							},
						},
					},
				},
			},
		}

		categoryAggregator := &categoryAggregatorImpl{categoryModelFactory: categoryModelFactory}
		result := categoryAggregator.AggregateCategoryTree(categoryEntityTree)

		assert.Equal(t, categoryTree0, result)
		assert.Equal(t, category0, result.Category)
		assert.Equal(t, categoryTree0categoryTree0, result.Categories[0])
		assert.Equal(t, category0category0, result.Categories[0].Category)
		assert.Equal(t, categoryTree0categoryTree1, result.Categories[1])
		assert.Equal(t, category0category1, result.Categories[1].Category)
		assert.Equal(t, categoryTree0categoryTree1categoryTree0, result.Categories[1].Categories[0])
		assert.Equal(t, category0category1category0, result.Categories[1].Categories[0].Category)
		assert.Equal(t, categoryTree0categoryTree1categoryTree1, result.Categories[1].Categories[1])
		assert.Equal(t, category0category1category1, result.Categories[1].Categories[1].Category)
	})

	t.Run("AggregateCategories", func(t *testing.T) {
		categories := new(models.Category)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategory().Return(categories)

		categoryAggregator := &categoryAggregatorImpl{categoryModelFactory: categoryModelFactory}
		categoryEntities := []*entities.CategoryEntity{new(entities.CategoryEntity)}
		response := categoryAggregator.AggregateCategories(categoryEntities)

		assert.IsType(t, []*models.Category{}, response)
		assert.Equal(t, len(categoryEntities), len(response))
	})

	t.Run("AggregateCategoriesTree", func(t *testing.T) {
		category0 := new(models.Category)
		category0category0 := new(models.Category)
		category0category1 := new(models.Category)
		category1 := new(models.Category)

		categoryTree0 := new(models.CategoryTree)
		categoryTree0categoryTree0 := new(models.CategoryTree)
		categoryTree0categoryTree1 := new(models.CategoryTree)
		categoryTree1 := new(models.CategoryTree)

		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)

		categoryModelFactory.EXPECT().CreateCategory().Return(category0)
		categoryModelFactory.EXPECT().CreateCategory().Return(category0category0)
		categoryModelFactory.EXPECT().CreateCategory().Return(category0category1)
		categoryModelFactory.EXPECT().CreateCategory().Return(category1)

		categoryModelFactory.EXPECT().CreateCategoryTree().Return(categoryTree0)
		categoryModelFactory.EXPECT().CreateCategoryTree().Return(categoryTree0categoryTree0)
		categoryModelFactory.EXPECT().CreateCategoryTree().Return(categoryTree0categoryTree1)
		categoryModelFactory.EXPECT().CreateCategoryTree().Return(categoryTree1)

		categoryEntityTree := &entities.CategoryEntityTree{
			Roots: []*entities.CategoryEntityTreeBranch{
				{
					CategoryEntity: new(entities.CategoryEntity),
					Children: []*entities.CategoryEntityTreeBranch{
						{
							CategoryEntity: new(entities.CategoryEntity),
						},
						{
							CategoryEntity: new(entities.CategoryEntity),
						},
					},
				},
				{
					CategoryEntity: new(entities.CategoryEntity),
				},
			},
		}

		categoryAggregator := &categoryAggregatorImpl{categoryModelFactory: categoryModelFactory}
		results := categoryAggregator.AggregateCategoriesTree(categoryEntityTree)

		assert.Equal(t, categoryTree0, results[0])
		assert.Equal(t, category0, results[0].Category)
		assert.Equal(t, categoryTree0categoryTree0, results[0].Categories[0])
		assert.Equal(t, category0category0, results[0].Categories[0].Category)
		assert.Equal(t, categoryTree0categoryTree1, results[0].Categories[1])
		assert.Equal(t, category0category1, results[0].Categories[1].Category)
		assert.Equal(t, categoryTree1, results[1])
		assert.Equal(t, category1, results[1].Category)
	})

	t.Run("AggregatePaginationResult", func(t *testing.T) {
		category := new(models.Category)
		categoryModelFactory := mocks.NewMockCategoryModelFactory(ctrl)
		categoryModelFactory.EXPECT().CreateCategory().Return(category)

		categoryEntities := []*entities.CategoryEntity{entities.NewCategoryEntity()}
		categoryAggregator := &categoryAggregatorImpl{categoryModelFactory: categoryModelFactory}

		entityPaginationResult := &models.PaginationResult{Data: categoryEntities}
		paginationResult := categoryAggregator.AggregatePaginationResult(entityPaginationResult)

		assert.IsType(t, []*models.Category{}, paginationResult.Data)
		assert.Equal(t, len(categoryEntities), len(paginationResult.Data.([]*models.Category)))
	})
}
