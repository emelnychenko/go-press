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

func TestCommentApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCommentApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		commentEventFactory := mocks.NewMockCommentEventFactory(ctrl)
		commentService := mocks.NewMockCommentService(ctrl)
		commentAggregator := mocks.NewMockCommentAggregator(ctrl)

		commentApi, isCommentApi := NewCommentApi(
			eventDispatcher, commentEventFactory, commentService, commentAggregator,
		).(*commentApiImpl)

		assert.True(t, isCommentApi)
		assert.Equal(t, eventDispatcher, commentApi.eventDispatcher)
		assert.Equal(t, commentEventFactory, commentApi.commentEventFactory)
		assert.Equal(t, commentService, commentApi.commentService)
		assert.Equal(t, commentAggregator, commentApi.commentAggregator)
	})

	t.Run("ListComments", func(t *testing.T) {
		paginationQuery := new(models.CommentPaginationQuery)
		entityPaginationResult := new(models.PaginationResult)
		paginationResult := new(models.PaginationResult)

		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().ListComments(paginationQuery).Return(entityPaginationResult, nil)

		commentAggregator := mocks.NewMockCommentAggregator(ctrl)
		commentAggregator.EXPECT().AggregatePaginationResult(entityPaginationResult).Return(paginationResult)

		commentApi := &commentApiImpl{commentService: commentService, commentAggregator: commentAggregator}
		response, err := commentApi.ListComments(paginationQuery)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListComments:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		paginationQuery := new(models.CommentPaginationQuery)

		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().ListComments(paginationQuery).Return(nil, systemErr)

		commentApi := &commentApiImpl{commentService: commentService}
		response, err := commentApi.ListComments(paginationQuery)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetComment", func(t *testing.T) {
		commentId := new(models.CommentId)
		commentEntity := new(entities.CommentEntity)
		comment := new(models.Comment)

		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().GetComment(commentId).Return(commentEntity, nil)

		commentAggregator := mocks.NewMockCommentAggregator(ctrl)
		commentAggregator.EXPECT().AggregateComment(commentEntity).Return(comment)

		commentApi := &commentApiImpl{commentService: commentService, commentAggregator: commentAggregator}
		response, err := commentApi.GetComment(commentId)

		assert.Equal(t, comment, response)
		assert.Nil(t, err)
	})

	t.Run("GetComment:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		commentId := new(models.CommentId)
		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().GetComment(commentId).Return(nil, systemErr)

		commentApi := &commentApiImpl{commentService: commentService}
		response, err := commentApi.GetComment(commentId)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateComment", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		comment := new(models.Comment)
		data := new(models.CommentCreate)

		commentEvent := new(events.CommentEvent)
		commentEventFactory := mocks.NewMockCommentEventFactory(ctrl)
		commentEventFactory.EXPECT().CreateCommentCreatedEvent(commentEntity).Return(commentEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(commentEvent)

		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().CreateComment(data).Return(commentEntity, nil)

		commentAggregator := mocks.NewMockCommentAggregator(ctrl)
		commentAggregator.EXPECT().AggregateComment(commentEntity).Return(comment)

		commentApi := &commentApiImpl{
			eventDispatcher:     eventDispatcher,
			commentEventFactory: commentEventFactory,
			commentService:      commentService,
			commentAggregator:   commentAggregator,
		}
		response, err := commentApi.CreateComment(data)

		assert.Equal(t, comment, response)
		assert.Nil(t, err)
	})

	t.Run("CreateComment:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		data := new(models.CommentCreate)
		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().CreateComment(data).Return(nil, systemErr)

		commentApi := &commentApiImpl{commentService: commentService}
		response, err := commentApi.CreateComment(data)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateComment", func(t *testing.T) {
		commentId := new(models.CommentId)
		commentEntity := new(entities.CommentEntity)
		data := new(models.CommentUpdate)

		commentEvent := new(events.CommentEvent)
		commentEventFactory := mocks.NewMockCommentEventFactory(ctrl)
		commentEventFactory.EXPECT().CreateCommentUpdatedEvent(commentEntity).Return(commentEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(commentEvent)

		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().GetComment(commentId).Return(commentEntity, nil)
		commentService.EXPECT().UpdateComment(commentEntity, data).Return(nil)

		commentApi := &commentApiImpl{
			eventDispatcher:     eventDispatcher,
			commentEventFactory: commentEventFactory,
			commentService:      commentService,
		}
		assert.Nil(t, commentApi.UpdateComment(commentId, data))
	})

	t.Run("UpdateComment:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		commentId := new(models.CommentId)
		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().GetComment(commentId).Return(nil, systemErr)

		data := new(models.CommentUpdate)
		commentApi := &commentApiImpl{commentService: commentService}
		assert.Equal(t, systemErr, commentApi.UpdateComment(commentId, data))
	})

	t.Run("UpdateComment:UpdateCommentError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		commentId := new(models.CommentId)
		commentEntity := new(entities.CommentEntity)
		data := new(models.CommentUpdate)

		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().GetComment(commentId).Return(commentEntity, nil)
		commentService.EXPECT().UpdateComment(commentEntity, data).Return(systemErr)

		commentApi := &commentApiImpl{
			commentService: commentService,
		}

		err := commentApi.UpdateComment(commentId, data)
		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteComment", func(t *testing.T) {
		commentId := new(models.CommentId)
		commentEntity := new(entities.CommentEntity)

		commentEvent := new(events.CommentEvent)
		commentEventFactory := mocks.NewMockCommentEventFactory(ctrl)
		commentEventFactory.EXPECT().CreateCommentDeletedEvent(commentEntity).Return(commentEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(commentEvent)

		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().GetComment(commentId).Return(commentEntity, nil)
		commentService.EXPECT().DeleteComment(commentEntity).Return(nil)

		commentApi := &commentApiImpl{
			eventDispatcher:     eventDispatcher,
			commentEventFactory: commentEventFactory,
			commentService:      commentService,
		}
		assert.Nil(t, commentApi.DeleteComment(commentId))
	})

	t.Run("DeleteComment:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		commentId := new(models.CommentId)
		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().GetComment(commentId).Return(nil, systemErr)

		commentApi := &commentApiImpl{commentService: commentService}
		assert.Equal(t, systemErr, commentApi.DeleteComment(commentId))
	})

	t.Run("DeleteComment:DeleteCommentError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		commentId := new(models.CommentId)
		commentEntity := new(entities.CommentEntity)

		commentService := mocks.NewMockCommentService(ctrl)
		commentService.EXPECT().GetComment(commentId).Return(commentEntity, nil)
		commentService.EXPECT().DeleteComment(commentEntity).Return(systemErr)

		commentApi := &commentApiImpl{
			commentService: commentService,
		}

		err := commentApi.DeleteComment(commentId)
		assert.Equal(t, systemErr, err)
	})
}
