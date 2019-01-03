package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	CategoryAggregator interface {
		AggregateCategory(categoryEntity *entities.CategoryEntity) *models.Category
		AggregateCategories(categoryEntities []*entities.CategoryEntity) []*models.Category
		AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult
	}
)
