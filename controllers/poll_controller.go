package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
)

type pollControllerImpl struct {
	pollHttpHelper   contracts.PollHttpHelper
	pollModelFactory contracts.PollModelFactory
	pollApi          contracts.PollApi
}

func NewPollController(
	pollHttpHelper contracts.PollHttpHelper,
	pollModelFactory contracts.PollModelFactory,
	pollApi contracts.PollApi,
) (pollController contracts.PollController) {
	return &pollControllerImpl{
		pollHttpHelper,
		pollModelFactory,
		pollApi,
	}
}

func (c *pollControllerImpl) ListPolls(httpContext contracts.HttpContext) (paginationResult interface{}, err errors.Error) {
	pollPaginationQuery := c.pollModelFactory.CreatePollPaginationQuery()

	if err = httpContext.BindModel(pollPaginationQuery.PaginationQuery); err != nil {
		return
	}

	if err = httpContext.BindModel(pollPaginationQuery); err != nil {
		return
	}

	paginationResult, err = c.pollApi.ListPolls(pollPaginationQuery)
	return
}

func (c *pollControllerImpl) GetPoll(httpContext contracts.HttpContext) (poll interface{}, err errors.Error) {
	pollId, err := c.pollHttpHelper.ParsePollId(httpContext)

	if err != nil {
		return
	}

	poll, err = c.pollApi.GetPoll(pollId)
	return
}

func (c *pollControllerImpl) CreatePoll(httpContext contracts.HttpContext) (poll interface{}, err errors.Error) {
	data := c.pollModelFactory.CreatePollCreate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	poll, err = c.pollApi.CreatePoll(data)
	return
}

func (c *pollControllerImpl) UpdatePoll(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
	pollId, err := c.pollHttpHelper.ParsePollId(httpContext)

	if err != nil {
		return
	}

	data := c.pollModelFactory.CreatePollUpdate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	err = c.pollApi.UpdatePoll(pollId, data)
	return
}

func (c *pollControllerImpl) DeletePoll(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
	pollId, err := c.pollHttpHelper.ParsePollId(httpContext)

	if err != nil {
		return
	}

	err = c.pollApi.DeletePoll(pollId)
	return
}
