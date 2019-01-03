package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type pollAggregatorImpl struct {
	pollModelFactory contracts.PollModelFactory
}

func NewPollAggregator(pollModelFactory contracts.PollModelFactory) contracts.PollAggregator {
	return &pollAggregatorImpl{pollModelFactory}
}

func (a *pollAggregatorImpl) AggregatePoll(pollEntity *entities.PollEntity) (poll *models.Poll) {
	poll = a.pollModelFactory.CreatePoll()
	poll.Id = pollEntity.Id
	poll.Title = pollEntity.Title
	poll.Created = pollEntity.Created

	return
}

func (a *pollAggregatorImpl) AggregatePolls(pollEntities []*entities.PollEntity) (polls []*models.Poll) {
	polls = make([]*models.Poll, len(pollEntities))

	for k, postEntity := range pollEntities {
		polls[k] = a.AggregatePoll(postEntity)
	}

	return
}

func (a *pollAggregatorImpl) AggregatePaginationResult(
	entityPaginationResult *models.PaginationResult,
) (
	paginationResult *models.PaginationResult,
) {
	pollEntities := entityPaginationResult.Data.([]*entities.PollEntity)
	polls := a.AggregatePolls(pollEntities)
	return &models.PaginationResult{Total: entityPaginationResult.Total, Data: polls}
}