package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	pollApiImpl struct {
		eventDispatcher  contracts.EventDispatcher
		pollEventFactory contracts.PollEventFactory
		pollService      contracts.PollService
		pollAggregator   contracts.PollAggregator
	}
)

func NewPollApi(
	eventDispatcher contracts.EventDispatcher,
	pollEventFactory contracts.PollEventFactory,
	pollService contracts.PollService,
	pollAggregator contracts.PollAggregator,
) (pollApi contracts.PollApi) {
	return &pollApiImpl{
		eventDispatcher,
		pollEventFactory,
		pollService,
		pollAggregator,
	}
}

func (a *pollApiImpl) ListPolls(
	pollPaginationQuery *models.PollPaginationQuery,
) (paginationResult *models.PaginationResult, err common.Error) {
	entityPaginationResult, err := a.pollService.ListPolls(pollPaginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.pollAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}

func (a *pollApiImpl) GetPoll(pollId *models.PollId) (poll *models.Poll, err common.Error) {
	pollEntity, err := a.pollService.GetPoll(pollId)

	if nil != err {
		return
	}

	poll = a.pollAggregator.AggregatePoll(pollEntity)
	return
}

func (a *pollApiImpl) CreatePoll(data *models.PollCreate) (poll *models.Poll, err common.Error) {
	pollEntity, err := a.pollService.CreatePoll(data)

	if nil != err {
		return
	}

	pollCreatedEvent := a.pollEventFactory.CreatePollCreatedEvent(pollEntity)
	a.eventDispatcher.Dispatch(pollCreatedEvent)

	poll = a.pollAggregator.AggregatePoll(pollEntity)
	return
}

func (a *pollApiImpl) UpdatePoll(pollId *models.PollId, data *models.PollUpdate) (err common.Error) {
	pollService := a.pollService
	pollEntity, err := pollService.GetPoll(pollId)

	if nil != err {
		return
	}

	err = pollService.UpdatePoll(pollEntity, data)

	if nil != err {
		return
	}

	pollUpdatedEvent := a.pollEventFactory.CreatePollUpdatedEvent(pollEntity)
	a.eventDispatcher.Dispatch(pollUpdatedEvent)
	return
}

func (a *pollApiImpl) DeletePoll(pollId *models.PollId) (err common.Error) {
	pollService := a.pollService
	pollEntity, err := pollService.GetPoll(pollId)

	if nil != err {
		return
	}

	err = pollService.DeletePoll(pollEntity)

	if nil != err {
		return
	}

	pollDeletedEvent := a.pollEventFactory.CreatePollDeletedEvent(pollEntity)
	a.eventDispatcher.Dispatch(pollDeletedEvent)

	return
}
