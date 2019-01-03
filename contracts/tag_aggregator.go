package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	TagAggregator interface {
		AggregateTag(tagEntity *entities.TagEntity) *models.Tag
		AggregateTags(tagEntities []*entities.TagEntity) []*models.Tag
		AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult
	}
)
