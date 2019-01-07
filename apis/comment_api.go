package apis

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	commentApiImpl struct {
		eventDispatcher     contracts.EventDispatcher
		commentEventFactory contracts.CommentEventFactory
		commentService      contracts.CommentService
		commentAggregator   contracts.CommentAggregator
	}
)

func NewCommentApi(
	eventDispatcher contracts.EventDispatcher,
	commentEventFactory contracts.CommentEventFactory,
	commentService contracts.CommentService,
	commentAggregator contracts.CommentAggregator,
) (commentApi contracts.CommentApi) {
	return &commentApiImpl{
		eventDispatcher,
		commentEventFactory,
		commentService,
		commentAggregator,
	}
}

func (a *commentApiImpl) ListComments(
	commentPaginationQuery *models.CommentPaginationQuery,
) (paginationResult *models.PaginationResult, err errors.Error) {
	entityPaginationResult, err := a.commentService.ListComments(commentPaginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.commentAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}

func (a *commentApiImpl) GetComment(commentId *models.CommentId) (comment *models.Comment, err errors.Error) {
	commentEntity, err := a.commentService.GetComment(commentId)

	if nil != err {
		return
	}

	comment = a.commentAggregator.AggregateComment(commentEntity)
	return
}

func (a *commentApiImpl) CreateComment(data *models.CommentCreate) (comment *models.Comment, err errors.Error) {
	commentEntity, err := a.commentService.CreateComment(data)

	if nil != err {
		return
	}

	commentCreatedEvent := a.commentEventFactory.CreateCommentCreatedEvent(commentEntity)
	a.eventDispatcher.Dispatch(commentCreatedEvent)

	comment = a.commentAggregator.AggregateComment(commentEntity)
	return
}

func (a *commentApiImpl) UpdateComment(commentId *models.CommentId, data *models.CommentUpdate) (err errors.Error) {
	commentService := a.commentService
	commentEntity, err := commentService.GetComment(commentId)

	if nil != err {
		return
	}

	err = commentService.UpdateComment(commentEntity, data)

	if nil != err {
		return
	}

	commentUpdatedEvent := a.commentEventFactory.CreateCommentUpdatedEvent(commentEntity)
	a.eventDispatcher.Dispatch(commentUpdatedEvent)
	return
}

func (a *commentApiImpl) DeleteComment(commentId *models.CommentId) (err errors.Error) {
	commentService := a.commentService
	commentEntity, err := commentService.GetComment(commentId)

	if nil != err {
		return
	}

	err = commentService.DeleteComment(commentEntity)

	if nil != err {
		return
	}

	commentDeletedEvent := a.commentEventFactory.CreateCommentDeletedEvent(commentEntity)
	a.eventDispatcher.Dispatch(commentDeletedEvent)

	return
}
