package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryAggregator interface {
		AggregateCategory(categoryEntity *entities.CategoryEntity) *models.Category
		AggregateCategoryTree(categoryEntityTree *entities.CategoryEntityTree) *models.CategoryTree
		AggregateCategories(categoryEntities []*entities.CategoryEntity) []*models.Category
		AggregateCategoriesTree(categoryEntityTree *entities.CategoryEntityTree) []*models.CategoryTree
		AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult
	}
)
