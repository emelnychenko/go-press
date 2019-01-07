package controllers

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCommentController", func(t *testing.T) {
		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentApi := mocks.NewMockCommentApi(ctrl)
		commentController, isCommentController := NewCommentController(
			commentHttpHelper,
			commentModelFactory,
			commentApi,
		).(*commentControllerImpl)

		assert.True(t, isCommentController)
		assert.Equal(t, commentHttpHelper, commentController.commentHttpHelper)
		assert.Equal(t, commentModelFactory, commentController.commentModelFactory)
		assert.Equal(t, commentApi, commentController.commentApi)
	})

	t.Run("ListComments", func(t *testing.T) {
		commentPaginationQuery := new(models.CommentPaginationQuery)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateCommentPaginationQuery().Return(commentPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(commentPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(commentPaginationQuery).Return(nil)

		var paginationResult *models.PaginationResult
		commentApi := mocks.NewMockCommentApi(ctrl)
		commentApi.EXPECT().ListComments(commentPaginationQuery).Return(paginationResult, nil)

		commentController := &commentControllerImpl{
			commentModelFactory: commentModelFactory,
			commentApi:          commentApi,
		}
		response, err := commentController.ListComments(httpContext)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListComments:BindPaginationQueryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		commentPaginationQuery := new(models.CommentPaginationQuery)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateCommentPaginationQuery().Return(commentPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(commentPaginationQuery.PaginationQuery).Return(systemErr)

		commentController := &commentControllerImpl{
			commentModelFactory: commentModelFactory,
		}
		response, err := commentController.ListComments(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListComments:BindCommentPaginationQueryError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		commentPaginationQuery := new(models.CommentPaginationQuery)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateCommentPaginationQuery().Return(commentPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(commentPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(commentPaginationQuery).Return(systemErr)

		commentController := &commentControllerImpl{
			commentModelFactory: commentModelFactory,
		}
		response, err := commentController.ListComments(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetComment", func(t *testing.T) {
		commentId := new(models.CommentId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		var comment *models.Comment
		commentApi := mocks.NewMockCommentApi(ctrl)
		commentApi.EXPECT().GetComment(commentId).Return(comment, nil)

		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentHttpHelper.EXPECT().ParseCommentId(httpContext).Return(commentId, nil)

		commentController := &commentControllerImpl{commentHttpHelper: commentHttpHelper, commentApi: commentApi}
		response, err := commentController.GetComment(httpContext)

		assert.Equal(t, comment, response)
		assert.Nil(t, err)
	})

	t.Run("GetComment:ParserError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentHttpHelper.EXPECT().ParseCommentId(httpContext).Return(nil, systemErr)

		commentController := &commentControllerImpl{commentHttpHelper: commentHttpHelper}
		response, err := commentController.GetComment(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetComment:ApiError", func(t *testing.T) {
		commentId := new(models.CommentId)
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		commentApi := mocks.NewMockCommentApi(ctrl)
		commentApi.EXPECT().GetComment(commentId).Return(nil, systemErr)

		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentHttpHelper.EXPECT().ParseCommentId(httpContext).Return(commentId, nil)

		commentController := &commentControllerImpl{commentHttpHelper: commentHttpHelper, commentApi: commentApi}
		response, err := commentController.GetComment(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateComment", func(t *testing.T) {
		comment := new(models.Comment)
		data := new(models.CommentCreate)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateCommentCreate().Return(data)

		commentApi := mocks.NewMockCommentApi(ctrl)
		commentApi.EXPECT().CreateComment(data).Return(comment, nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		commentController := &commentControllerImpl{
			commentModelFactory: commentModelFactory,
			commentApi:          commentApi,
		}
		response, err := commentController.CreateComment(httpContext)

		assert.Equal(t, comment, response)
		assert.Nil(t, err)
	})

	t.Run("CreateComment:BindCommentUpdateError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		data := new(models.CommentCreate)

		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateCommentCreate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		commentController := &commentControllerImpl{
			commentModelFactory: commentModelFactory,
		}
		_, err := commentController.CreateComment(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateComment:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		data := new(models.CommentCreate)

		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateCommentCreate().Return(data)

		commentApi := mocks.NewMockCommentApi(ctrl)
		commentApi.EXPECT().CreateComment(data).Return(nil, systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		commentController := &commentControllerImpl{
			commentModelFactory: commentModelFactory,
			commentApi:          commentApi,
		}
		_, err := commentController.CreateComment(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateComment", func(t *testing.T) {
		commentId := new(models.CommentId)
		data := new(models.CommentUpdate)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateCommentUpdate().Return(data)

		commentApi := mocks.NewMockCommentApi(ctrl)
		commentApi.EXPECT().UpdateComment(commentId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentHttpHelper.EXPECT().ParseCommentId(httpContext).Return(commentId, nil)

		commentController := &commentControllerImpl{
			commentHttpHelper:   commentHttpHelper,
			commentModelFactory: commentModelFactory,
			commentApi:          commentApi,
		}
		_, err := commentController.UpdateComment(httpContext)

		assert.Nil(t, err)
	})

	t.Run("UpdateComment:ParseError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentHttpHelper.EXPECT().ParseCommentId(httpContext).Return(nil, systemErr)

		commentController := &commentControllerImpl{commentHttpHelper: commentHttpHelper}
		_, err := commentController.UpdateComment(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateComment:BindCommentUpdateError", func(t *testing.T) {
		commentId := new(models.CommentId)
		systemErr := errors.NewUnknownError()
		data := new(models.CommentUpdate)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateCommentUpdate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentHttpHelper.EXPECT().ParseCommentId(httpContext).Return(commentId, nil)

		commentController := &commentControllerImpl{
			commentHttpHelper:   commentHttpHelper,
			commentModelFactory: commentModelFactory,
		}
		_, err := commentController.UpdateComment(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateComment:ApiError", func(t *testing.T) {
		commentId := new(models.CommentId)
		systemErr := errors.NewUnknownError()

		data := new(models.CommentUpdate)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateCommentUpdate().Return(data)

		commentApi := mocks.NewMockCommentApi(ctrl)
		commentApi.EXPECT().UpdateComment(commentId, data).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentHttpHelper.EXPECT().ParseCommentId(httpContext).Return(commentId, nil)

		commentController := &commentControllerImpl{
			commentHttpHelper:   commentHttpHelper,
			commentModelFactory: commentModelFactory,
			commentApi:          commentApi,
		}
		_, err := commentController.UpdateComment(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteComment", func(t *testing.T) {
		commentId := new(models.CommentId)

		commentApi := mocks.NewMockCommentApi(ctrl)
		commentApi.EXPECT().DeleteComment(commentId).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)

		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentHttpHelper.EXPECT().ParseCommentId(httpContext).Return(commentId, nil)

		commentController := &commentControllerImpl{commentHttpHelper: commentHttpHelper, commentApi: commentApi}
		_, err := commentController.DeleteComment(httpContext)

		assert.Nil(t, err)
	})

	t.Run("DeleteComment:ParseError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		httpContext := mocks.NewMockHttpContext(ctrl)

		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentHttpHelper.EXPECT().ParseCommentId(httpContext).Return(nil, systemErr)

		commentController := &commentControllerImpl{commentHttpHelper: commentHttpHelper}
		_, err := commentController.DeleteComment(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteComment:ApiError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		commentId := new(models.CommentId)

		commentApi := mocks.NewMockCommentApi(ctrl)
		commentApi.EXPECT().DeleteComment(commentId).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)

		commentHttpHelper := mocks.NewMockCommentHttpHelper(ctrl)
		commentHttpHelper.EXPECT().ParseCommentId(httpContext).Return(commentId, nil)

		commentController := &commentControllerImpl{commentHttpHelper: commentHttpHelper, commentApi: commentApi}
		_, err := commentController.DeleteComment(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
