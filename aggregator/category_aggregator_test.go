package aggregator

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
