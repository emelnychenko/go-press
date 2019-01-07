package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
)

type channelControllerImpl struct {
	channelHttpHelper   contracts.ChannelHttpHelper
	channelModelFactory contracts.ChannelModelFactory
	channelApi          contracts.ChannelApi
}

func NewChannelController(
	channelHttpHelper contracts.ChannelHttpHelper,
	channelModelFactory contracts.ChannelModelFactory,
	channelApi contracts.ChannelApi,
) (channelController contracts.ChannelController) {
	return &channelControllerImpl{
		channelHttpHelper,
		channelModelFactory,
		channelApi,
	}
}

func (c *channelControllerImpl) ListChannels(httpContext contracts.HttpContext) (paginationResult interface{}, err errors.Error) {
	channelPaginationQuery := c.channelModelFactory.CreateChannelPaginationQuery()

	if err = httpContext.BindModel(channelPaginationQuery.PaginationQuery); err != nil {
		return
	}

	if err = httpContext.BindModel(channelPaginationQuery); err != nil {
		return
	}

	paginationResult, err = c.channelApi.ListChannels(channelPaginationQuery)
	return
}

func (c *channelControllerImpl) GetChannel(httpContext contracts.HttpContext) (channel interface{}, err errors.Error) {
	channelId, err := c.channelHttpHelper.ParseChannelId(httpContext)

	if err != nil {
		return
	}

	channel, err = c.channelApi.GetChannel(channelId)
	return
}

func (c *channelControllerImpl) CreateChannel(httpContext contracts.HttpContext) (channel interface{}, err errors.Error) {
	data := c.channelModelFactory.CreateChannelCreate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	channel, err = c.channelApi.CreateChannel(data)
	return
}

func (c *channelControllerImpl) UpdateChannel(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
	channelId, err := c.channelHttpHelper.ParseChannelId(httpContext)

	if err != nil {
		return
	}

	data := c.channelModelFactory.CreateChannelUpdate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	err = c.channelApi.UpdateChannel(channelId, data)
	return
}

func (c *channelControllerImpl) DeleteChannel(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
	channelId, err := c.channelHttpHelper.ParseChannelId(httpContext)

	if err != nil {
		return
	}

	err = c.channelApi.DeleteChannel(channelId)
	return
}
