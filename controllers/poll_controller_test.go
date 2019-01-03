package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPollController", func(t *testing.T) {
		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollApi := mocks.NewMockPollApi(ctrl)
		pollController, isPollController := NewPollController(
			pollHttpHelper,
			pollModelFactory,
			pollApi,
		).(*pollControllerImpl)

		assert.True(t, isPollController)
		assert.Equal(t, pollHttpHelper, pollController.pollHttpHelper)
		assert.Equal(t, pollModelFactory, pollController.pollModelFactory)
		assert.Equal(t, pollApi, pollController.pollApi)
	})

	t.Run("ListPolls", func(t *testing.T) {
		pollPaginationQuery := new(models.PollPaginationQuery)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePollPaginationQuery().Return(pollPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(pollPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(pollPaginationQuery).Return(nil)

		var paginationResult *models.PaginationResult
		pollApi := mocks.NewMockPollApi(ctrl)
		pollApi.EXPECT().ListPolls(pollPaginationQuery).Return(paginationResult, nil)

		pollController := &pollControllerImpl{
			pollModelFactory: pollModelFactory,
			pollApi:          pollApi,
		}
		response, err := pollController.ListPolls(httpContext)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListPolls:BindPaginationQueryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		pollPaginationQuery := new(models.PollPaginationQuery)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePollPaginationQuery().Return(pollPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(pollPaginationQuery.PaginationQuery).Return(systemErr)

		pollController := &pollControllerImpl{
			pollModelFactory: pollModelFactory,
		}
		response, err := pollController.ListPolls(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListPolls:BindPollPaginationQueryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		pollPaginationQuery := new(models.PollPaginationQuery)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePollPaginationQuery().Return(pollPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(pollPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(pollPaginationQuery).Return(systemErr)

		pollController := &pollControllerImpl{
			pollModelFactory: pollModelFactory,
		}
		response, err := pollController.ListPolls(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetPoll", func(t *testing.T) {
		pollId := new(models.PollId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		var poll *models.Poll
		pollApi := mocks.NewMockPollApi(ctrl)
		pollApi.EXPECT().GetPoll(pollId).Return(poll, nil)

		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollHttpHelper.EXPECT().ParsePollId(httpContext).Return(pollId, nil)

		pollController := &pollControllerImpl{pollHttpHelper: pollHttpHelper, pollApi: pollApi}
		response, err := pollController.GetPoll(httpContext)

		assert.Equal(t, poll, response)
		assert.Nil(t, err)
	})

	t.Run("GetPoll:ParserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollHttpHelper.EXPECT().ParsePollId(httpContext).Return(nil, systemErr)

		pollController := &pollControllerImpl{pollHttpHelper: pollHttpHelper}
		response, err := pollController.GetPoll(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetPoll:ApiError", func(t *testing.T) {
		pollId := new(models.PollId)
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		pollApi := mocks.NewMockPollApi(ctrl)
		pollApi.EXPECT().GetPoll(pollId).Return(nil, systemErr)

		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollHttpHelper.EXPECT().ParsePollId(httpContext).Return(pollId, nil)

		pollController := &pollControllerImpl{pollHttpHelper: pollHttpHelper, pollApi: pollApi}
		response, err := pollController.GetPoll(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreatePoll", func(t *testing.T) {
		poll := new(models.Poll)
		data := new(models.PollCreate)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePollCreate().Return(data)

		pollApi := mocks.NewMockPollApi(ctrl)
		pollApi.EXPECT().CreatePoll(data).Return(poll, nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		pollController := &pollControllerImpl{
			pollModelFactory: pollModelFactory,
			pollApi:          pollApi,
		}
		response, err := pollController.CreatePoll(httpContext)

		assert.Equal(t, poll, response)
		assert.Nil(t, err)
	})

	t.Run("CreatePoll:BindPollUpdateError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.PollCreate)

		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePollCreate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		pollController := &pollControllerImpl{
			pollModelFactory: pollModelFactory,
		}
		_, err := pollController.CreatePoll(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("CreatePoll:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.PollCreate)

		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePollCreate().Return(data)

		pollApi := mocks.NewMockPollApi(ctrl)
		pollApi.EXPECT().CreatePoll(data).Return(nil, systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		pollController := &pollControllerImpl{
			pollModelFactory: pollModelFactory,
			pollApi:          pollApi,
		}
		_, err := pollController.CreatePoll(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdatePoll", func(t *testing.T) {
		pollId := new(models.PollId)
		data := new(models.PollUpdate)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePollUpdate().Return(data)

		pollApi := mocks.NewMockPollApi(ctrl)
		pollApi.EXPECT().UpdatePoll(pollId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollHttpHelper.EXPECT().ParsePollId(httpContext).Return(pollId, nil)

		pollController := &pollControllerImpl{
			pollHttpHelper:   pollHttpHelper,
			pollModelFactory: pollModelFactory,
			pollApi:          pollApi,
		}
		_, err := pollController.UpdatePoll(httpContext)

		assert.Nil(t, err)
	})

	t.Run("UpdatePoll:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollHttpHelper.EXPECT().ParsePollId(httpContext).Return(nil, systemErr)

		pollController := &pollControllerImpl{pollHttpHelper: pollHttpHelper}
		_, err := pollController.UpdatePoll(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdatePoll:BindPollUpdateError", func(t *testing.T) {
		pollId := new(models.PollId)
		systemErr := common.NewUnknownError()
		data := new(models.PollUpdate)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePollUpdate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollHttpHelper.EXPECT().ParsePollId(httpContext).Return(pollId, nil)

		pollController := &pollControllerImpl{
			pollHttpHelper:   pollHttpHelper,
			pollModelFactory: pollModelFactory,
		}
		_, err := pollController.UpdatePoll(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdatePoll:ApiError", func(t *testing.T) {
		pollId := new(models.PollId)
		systemErr := common.NewUnknownError()

		data := new(models.PollUpdate)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePollUpdate().Return(data)

		pollApi := mocks.NewMockPollApi(ctrl)
		pollApi.EXPECT().UpdatePoll(pollId, data).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollHttpHelper.EXPECT().ParsePollId(httpContext).Return(pollId, nil)

		pollController := &pollControllerImpl{
			pollHttpHelper:   pollHttpHelper,
			pollModelFactory: pollModelFactory,
			pollApi:          pollApi,
		}
		_, err := pollController.UpdatePoll(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeletePoll", func(t *testing.T) {
		pollId := new(models.PollId)

		pollApi := mocks.NewMockPollApi(ctrl)
		pollApi.EXPECT().DeletePoll(pollId).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)

		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollHttpHelper.EXPECT().ParsePollId(httpContext).Return(pollId, nil)

		pollController := &pollControllerImpl{pollHttpHelper: pollHttpHelper, pollApi: pollApi}
		_, err := pollController.DeletePoll(httpContext)

		assert.Nil(t, err)
	})

	t.Run("DeletePoll:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		httpContext := mocks.NewMockHttpContext(ctrl)

		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollHttpHelper.EXPECT().ParsePollId(httpContext).Return(nil, systemErr)

		pollController := &pollControllerImpl{pollHttpHelper: pollHttpHelper}
		_, err := pollController.DeletePoll(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeletePoll:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		pollId := new(models.PollId)

		pollApi := mocks.NewMockPollApi(ctrl)
		pollApi.EXPECT().DeletePoll(pollId).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)

		pollHttpHelper := mocks.NewMockPollHttpHelper(ctrl)
		pollHttpHelper.EXPECT().ParsePollId(httpContext).Return(pollId, nil)

		pollController := &pollControllerImpl{pollHttpHelper: pollHttpHelper, pollApi: pollApi}
		_, err := pollController.DeletePoll(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
