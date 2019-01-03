package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type commentAggregatorImpl struct {
	commentModelFactory contracts.CommentModelFactory
}

func NewCommentAggregator(commentModelFactory contracts.CommentModelFactory) contracts.CommentAggregator {
	return &commentAggregatorImpl{commentModelFactory}
}

func (a *commentAggregatorImpl) AggregateComment(commentEntity *entities.CommentEntity) (comment *models.Comment) {
	comment = a.commentModelFactory.CreateComment()
	comment.Id = commentEntity.Id
	comment.Content = commentEntity.Content
	comment.Created = commentEntity.Created

	return
}

func (a *commentAggregatorImpl) AggregateComments(commentEntities []*entities.CommentEntity) (comments []*models.Comment) {
	comments = make([]*models.Comment, len(commentEntities))

	for k, postEntity := range commentEntities {
		comments[k] = a.AggregateComment(postEntity)
	}

	return
}

func (a *commentAggregatorImpl) AggregatePaginationResult(
	entityPaginationResult *models.PaginationResult,
) (
	paginationResult *models.PaginationResult,
) {
	commentEntities := entityPaginationResult.Data.([]*entities.CommentEntity)
	comments := a.AggregateComments(commentEntities)
	return &models.PaginationResult{Total: entityPaginationResult.Total, Data: comments}
}
