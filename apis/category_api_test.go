package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCategoryApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		categoryEventFactory := mocks.NewMockCategoryEventFactory(ctrl)
		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryAggregator := mocks.NewMockCategoryAggregator(ctrl)

		categoryApi, isCategoryApi := NewCategoryApi(
			eventDispatcher, categoryEventFactory, categoryService, categoryAggregator,
		).(*categoryApiImpl)

		assert.True(t, isCategoryApi)
		assert.Equal(t, eventDispatcher, categoryApi.eventDispatcher)
		assert.Equal(t, categoryEventFactory, categoryApi.categoryEventFactory)
		assert.Equal(t, categoryService, categoryApi.categoryService)
		assert.Equal(t, categoryAggregator, categoryApi.categoryAggregator)
	})

	t.Run("ListCategories", func(t *testing.T) {
		paginationQuery := new(models.CategoryPaginationQuery)
		entityPaginationResult := new(models.PaginationResult)
		paginationResult := new(models.PaginationResult)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().ListCategories(paginationQuery).Return(entityPaginationResult, nil)

		categoryAggregator := mocks.NewMockCategoryAggregator(ctrl)
		categoryAggregator.EXPECT().AggregatePaginationResult(entityPaginationResult).Return(paginationResult)

		categoryApi := &categoryApiImpl{categoryService: categoryService, categoryAggregator: categoryAggregator}
		response, err := categoryApi.ListCategories(paginationQuery)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListCategories:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		paginationQuery := new(models.CategoryPaginationQuery)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().ListCategories(paginationQuery).Return(nil, systemErr)

		categoryApi := &categoryApiImpl{categoryService: categoryService}
		response, err := categoryApi.ListCategories(paginationQuery)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetCategory", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)
		category := new(models.Category)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(categoryEntity, nil)

		categoryAggregator := mocks.NewMockCategoryAggregator(ctrl)
		categoryAggregator.EXPECT().AggregateCategory(categoryEntity).Return(category)

		categoryApi := &categoryApiImpl{categoryService: categoryService, categoryAggregator: categoryAggregator}
		response, err := categoryApi.GetCategory(categoryId)

		assert.Equal(t, category, response)
		assert.Nil(t, err)
	})

	t.Run("GetCategory:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		categoryId := new(models.CategoryId)
		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(nil, systemErr)

		categoryApi := &categoryApiImpl{categoryService: categoryService}
		response, err := categoryApi.GetCategory(categoryId)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateCategory", func(t *testing.T) {
		categoryEntity := new(entities.CategoryEntity)
		category := new(models.Category)
		data := new(models.CategoryCreate)

		categoryEvent := new(events.CategoryEvent)
		categoryEventFactory := mocks.NewMockCategoryEventFactory(ctrl)
		categoryEventFactory.EXPECT().CreateCategoryCreatedEvent(categoryEntity).Return(categoryEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(categoryEvent)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().CreateCategory(data).Return(categoryEntity, nil)

		categoryAggregator := mocks.NewMockCategoryAggregator(ctrl)
		categoryAggregator.EXPECT().AggregateCategory(categoryEntity).Return(category)

		categoryApi := &categoryApiImpl{
			eventDispatcher: eventDispatcher,
			categoryEventFactory: categoryEventFactory,
			categoryService: categoryService,
			categoryAggregator: categoryAggregator,
		}
		response, err := categoryApi.CreateCategory(data)

		assert.Equal(t, category, response)
		assert.Nil(t, err)
	})

	t.Run("CreateCategory:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		data := new(models.CategoryCreate)
		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().CreateCategory(data).Return(nil, systemErr)

		categoryApi := &categoryApiImpl{categoryService: categoryService}
		response, err := categoryApi.CreateCategory(data)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateCategory", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)
		data := new(models.CategoryUpdate)

		categoryEvent := new(events.CategoryEvent)
		categoryEventFactory := mocks.NewMockCategoryEventFactory(ctrl)
		categoryEventFactory.EXPECT().CreateCategoryUpdatedEvent(categoryEntity).Return(categoryEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(categoryEvent)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(categoryEntity, nil)
		categoryService.EXPECT().UpdateCategory(categoryEntity, data).Return(nil)

		categoryApi := &categoryApiImpl{
			eventDispatcher: eventDispatcher,
			categoryEventFactory: categoryEventFactory,
			categoryService: categoryService,
		}
		assert.Nil(t, categoryApi.UpdateCategory(categoryId, data))
	})

	t.Run("UpdateCategory:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		categoryId := new(models.CategoryId)
		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(nil, systemErr)

		data := new(models.CategoryUpdate)
		categoryApi := &categoryApiImpl{categoryService: categoryService}
		assert.Equal(t, systemErr, categoryApi.UpdateCategory(categoryId, data))
	})

	t.Run("UpdateCategory:UpdateCategoryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		categoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)
		data := new(models.CategoryUpdate)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(categoryEntity, nil)
		categoryService.EXPECT().UpdateCategory(categoryEntity, data).Return(systemErr)

		categoryApi := &categoryApiImpl{
			categoryService: categoryService,
		}

		err := categoryApi.UpdateCategory(categoryId, data)
		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteCategory", func(t *testing.T) {
		categoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)

		categoryEvent := new(events.CategoryEvent)
		categoryEventFactory := mocks.NewMockCategoryEventFactory(ctrl)
		categoryEventFactory.EXPECT().CreateCategoryDeletedEvent(categoryEntity).Return(categoryEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(categoryEvent)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(categoryEntity, nil)
		categoryService.EXPECT().DeleteCategory(categoryEntity).Return(nil)

		categoryApi := &categoryApiImpl{
			eventDispatcher: eventDispatcher,
			categoryEventFactory: categoryEventFactory,
			categoryService: categoryService,
		}
		assert.Nil(t, categoryApi.DeleteCategory(categoryId))
	})

	t.Run("DeleteCategory:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		categoryId := new(models.CategoryId)
		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(nil, systemErr)

		categoryApi := &categoryApiImpl{categoryService: categoryService}
		assert.Equal(t, systemErr, categoryApi.DeleteCategory(categoryId))
	})

	t.Run("DeleteCategory:DeleteCategoryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		categoryId := new(models.CategoryId)
		categoryEntity := new(entities.CategoryEntity)

		categoryService := mocks.NewMockCategoryService(ctrl)
		categoryService.EXPECT().GetCategory(categoryId).Return(categoryEntity, nil)
		categoryService.EXPECT().DeleteCategory(categoryEntity).Return(systemErr)

		categoryApi := &categoryApiImpl{
			categoryService: categoryService,
		}

		err := categoryApi.DeleteCategory(categoryId)
		assert.Equal(t, systemErr, err)
	})
}
