package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	CommentAggregator interface {
		AggregateComment(commentEntity *entities.CommentEntity) *models.Comment
		AggregateComments(commentEntities []*entities.CommentEntity) []*models.Comment
		AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult
	}
)
