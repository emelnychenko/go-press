package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostAggregator interface {
		AggregatePost(postEntity *entities.PostEntity) (post *models.Post)
		AggregatePosts(postEntities []*entities.PostEntity) (posts []*models.Post)
		AggregatePaginationResult(
			entityPaginationResult *models.PaginationResult,
		) (paginationResult *models.PaginationResult)
	}
)
