package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type categoryAggregatorImpl struct {
	categoryModelFactory contracts.CategoryModelFactory
}

func NewCategoryAggregator(categoryModelFactory contracts.CategoryModelFactory) contracts.CategoryAggregator {
	return &categoryAggregatorImpl{categoryModelFactory}
}

func (a *categoryAggregatorImpl) AggregateCategory(categoryEntity *entities.CategoryEntity) (category *models.Category) {
	category = a.categoryModelFactory.CreateCategory()
	category.Id = categoryEntity.Id
	category.Name = categoryEntity.Name
	category.Created = categoryEntity.Created

	return
}

func (a *categoryAggregatorImpl) AggregateCategories(categoryEntities []*entities.CategoryEntity) (categorys []*models.Category) {
	categorys = make([]*models.Category, len(categoryEntities))

	for k, postEntity := range categoryEntities {
		categorys[k] = a.AggregateCategory(postEntity)
	}

	return
}

func (a *categoryAggregatorImpl) AggregatePaginationResult(
	entityPaginationResult *models.PaginationResult,
) (
	paginationResult *models.PaginationResult,
) {
	categoryEntities := entityPaginationResult.Data.([]*entities.CategoryEntity)
	categorys := a.AggregateCategories(categoryEntities)
	return &models.PaginationResult{Total: entityPaginationResult.Total, Data: categorys}
}
