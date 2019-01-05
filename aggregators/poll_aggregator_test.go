package aggregators

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollAggregator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPollAggregator", func(t *testing.T) {
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollAggregator, isPollAggregator := NewPollAggregator(
			pollModelFactory,
		).(*pollAggregatorImpl)

		assert.True(t, isPollAggregator)
		assert.Equal(t, pollModelFactory, pollAggregator.pollModelFactory)
	})

	t.Run("AggregatePoll", func(t *testing.T) {
		poll := new(models.Poll)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePoll().Return(poll)

		pollAggregator := &pollAggregatorImpl{pollModelFactory: pollModelFactory}
		response := pollAggregator.AggregatePoll(new(entities.PollEntity))

		assert.Equal(t, poll, response)
	})

	t.Run("AggregatePolls", func(t *testing.T) {
		polls := new(models.Poll)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePoll().Return(polls)

		pollAggregator := &pollAggregatorImpl{pollModelFactory: pollModelFactory}
		pollEntities := []*entities.PollEntity{new(entities.PollEntity)}
		response := pollAggregator.AggregatePolls(pollEntities)

		assert.IsType(t, []*models.Poll{}, response)
		assert.Equal(t, len(pollEntities), len(response))
	})

	t.Run("AggregatePaginationResult", func(t *testing.T) {
		poll := new(models.Poll)
		pollModelFactory := mocks.NewMockPollModelFactory(ctrl)
		pollModelFactory.EXPECT().CreatePoll().Return(poll)

		pollEntities := []*entities.PollEntity{entities.NewPollEntity()}
		pollAggregator := &pollAggregatorImpl{pollModelFactory: pollModelFactory}

		entityPaginationResult := &models.PaginationResult{Data: pollEntities}
		paginationResult := pollAggregator.AggregatePaginationResult(entityPaginationResult)

		assert.IsType(t, []*models.Poll{}, paginationResult.Data)
		assert.Equal(t, len(pollEntities), len(paginationResult.Data.([]*models.Poll)))
	})
}
