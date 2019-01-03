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

func TestPollApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPollApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		pollEventFactory := mocks.NewMockPollEventFactory(ctrl)
		pollService := mocks.NewMockPollService(ctrl)
		pollAggregator := mocks.NewMockPollAggregator(ctrl)

		pollApi, isPollApi := NewPollApi(
			eventDispatcher, pollEventFactory, pollService, pollAggregator,
		).(*pollApiImpl)

		assert.True(t, isPollApi)
		assert.Equal(t, eventDispatcher, pollApi.eventDispatcher)
		assert.Equal(t, pollEventFactory, pollApi.pollEventFactory)
		assert.Equal(t, pollService, pollApi.pollService)
		assert.Equal(t, pollAggregator, pollApi.pollAggregator)
	})

	t.Run("ListPolls", func(t *testing.T) {
		paginationQuery := new(models.PollPaginationQuery)
		entityPaginationResult := new(models.PaginationResult)
		paginationResult := new(models.PaginationResult)

		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().ListPolls(paginationQuery).Return(entityPaginationResult, nil)

		pollAggregator := mocks.NewMockPollAggregator(ctrl)
		pollAggregator.EXPECT().AggregatePaginationResult(entityPaginationResult).Return(paginationResult)

		pollApi := &pollApiImpl{pollService: pollService, pollAggregator: pollAggregator}
		response, err := pollApi.ListPolls(paginationQuery)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListPolls:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		paginationQuery := new(models.PollPaginationQuery)

		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().ListPolls(paginationQuery).Return(nil, systemErr)

		pollApi := &pollApiImpl{pollService: pollService}
		response, err := pollApi.ListPolls(paginationQuery)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetPoll", func(t *testing.T) {
		pollId := new(models.PollId)
		pollEntity := new(entities.PollEntity)
		poll := new(models.Poll)

		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().GetPoll(pollId).Return(pollEntity, nil)

		pollAggregator := mocks.NewMockPollAggregator(ctrl)
		pollAggregator.EXPECT().AggregatePoll(pollEntity).Return(poll)

		pollApi := &pollApiImpl{pollService: pollService, pollAggregator: pollAggregator}
		response, err := pollApi.GetPoll(pollId)

		assert.Equal(t, poll, response)
		assert.Nil(t, err)
	})

	t.Run("GetPoll:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		pollId := new(models.PollId)
		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().GetPoll(pollId).Return(nil, systemErr)

		pollApi := &pollApiImpl{pollService: pollService}
		response, err := pollApi.GetPoll(pollId)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreatePoll", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		poll := new(models.Poll)
		data := new(models.PollCreate)

		pollEvent := new(events.PollEvent)
		pollEventFactory := mocks.NewMockPollEventFactory(ctrl)
		pollEventFactory.EXPECT().CreatePollCreatedEvent(pollEntity).Return(pollEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(pollEvent)

		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().CreatePoll(data).Return(pollEntity, nil)

		pollAggregator := mocks.NewMockPollAggregator(ctrl)
		pollAggregator.EXPECT().AggregatePoll(pollEntity).Return(poll)

		pollApi := &pollApiImpl{
			eventDispatcher: eventDispatcher,
			pollEventFactory: pollEventFactory,
			pollService: pollService,
			pollAggregator: pollAggregator,
		}
		response, err := pollApi.CreatePoll(data)

		assert.Equal(t, poll, response)
		assert.Nil(t, err)
	})

	t.Run("CreatePoll:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		data := new(models.PollCreate)
		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().CreatePoll(data).Return(nil, systemErr)

		pollApi := &pollApiImpl{pollService: pollService}
		response, err := pollApi.CreatePoll(data)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdatePoll", func(t *testing.T) {
		pollId := new(models.PollId)
		pollEntity := new(entities.PollEntity)
		data := new(models.PollUpdate)

		pollEvent := new(events.PollEvent)
		pollEventFactory := mocks.NewMockPollEventFactory(ctrl)
		pollEventFactory.EXPECT().CreatePollUpdatedEvent(pollEntity).Return(pollEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(pollEvent)

		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().GetPoll(pollId).Return(pollEntity, nil)
		pollService.EXPECT().UpdatePoll(pollEntity, data).Return(nil)

		pollApi := &pollApiImpl{
			eventDispatcher: eventDispatcher,
			pollEventFactory: pollEventFactory,
			pollService: pollService,
		}
		assert.Nil(t, pollApi.UpdatePoll(pollId, data))
	})

	t.Run("UpdatePoll:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		pollId := new(models.PollId)
		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().GetPoll(pollId).Return(nil, systemErr)

		data := new(models.PollUpdate)
		pollApi := &pollApiImpl{pollService: pollService}
		assert.Equal(t, systemErr, pollApi.UpdatePoll(pollId, data))
	})

	t.Run("UpdatePoll:UpdatePollError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		pollId := new(models.PollId)
		pollEntity := new(entities.PollEntity)
		data := new(models.PollUpdate)

		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().GetPoll(pollId).Return(pollEntity, nil)
		pollService.EXPECT().UpdatePoll(pollEntity, data).Return(systemErr)

		pollApi := &pollApiImpl{
			pollService: pollService,
		}

		err := pollApi.UpdatePoll(pollId, data)
		assert.Equal(t, systemErr, err)
	})

	t.Run("DeletePoll", func(t *testing.T) {
		pollId := new(models.PollId)
		pollEntity := new(entities.PollEntity)

		pollEvent := new(events.PollEvent)
		pollEventFactory := mocks.NewMockPollEventFactory(ctrl)
		pollEventFactory.EXPECT().CreatePollDeletedEvent(pollEntity).Return(pollEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(pollEvent)

		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().GetPoll(pollId).Return(pollEntity, nil)
		pollService.EXPECT().DeletePoll(pollEntity).Return(nil)

		pollApi := &pollApiImpl{
			eventDispatcher: eventDispatcher,
			pollEventFactory: pollEventFactory,
			pollService: pollService,
		}
		assert.Nil(t, pollApi.DeletePoll(pollId))
	})

	t.Run("DeletePoll:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		pollId := new(models.PollId)
		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().GetPoll(pollId).Return(nil, systemErr)

		pollApi := &pollApiImpl{pollService: pollService}
		assert.Equal(t, systemErr, pollApi.DeletePoll(pollId))
	})

	t.Run("DeletePoll:DeletePollError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		pollId := new(models.PollId)
		pollEntity := new(entities.PollEntity)

		pollService := mocks.NewMockPollService(ctrl)
		pollService.EXPECT().GetPoll(pollId).Return(pollEntity, nil)
		pollService.EXPECT().DeletePoll(pollEntity).Return(systemErr)

		pollApi := &pollApiImpl{
			pollService: pollService,
		}

		err := pollApi.DeletePoll(pollId)
		assert.Equal(t, systemErr, err)
	})
}
