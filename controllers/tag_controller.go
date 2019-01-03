package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
)

type tagControllerImpl struct {
	tagHttpHelper   contracts.TagHttpHelper
	tagModelFactory contracts.TagModelFactory
	tagApi          contracts.TagApi
}

func NewTagController(
	tagHttpHelper contracts.TagHttpHelper,
	tagModelFactory contracts.TagModelFactory,
	tagApi contracts.TagApi,
) (tagController contracts.TagController) {
	return &tagControllerImpl{
		tagHttpHelper,
		tagModelFactory,
		tagApi,
	}
}

func (c *tagControllerImpl) ListTags(httpContext contracts.HttpContext) (paginationResult interface{}, err common.Error) {
	tagPaginationQuery := c.tagModelFactory.CreateTagPaginationQuery()

	if err = httpContext.BindModel(tagPaginationQuery.PaginationQuery); err != nil {
		return
	}

	if err = httpContext.BindModel(tagPaginationQuery); err != nil {
		return
	}

	paginationResult, err = c.tagApi.ListTags(tagPaginationQuery)
	return
}

func (c *tagControllerImpl) GetTag(httpContext contracts.HttpContext) (tag interface{}, err common.Error) {
	tagId, err := c.tagHttpHelper.ParseTagId(httpContext)

	if err != nil {
		return
	}

	tag, err = c.tagApi.GetTag(tagId)
	return
}

func (c *tagControllerImpl) CreateTag(httpContext contracts.HttpContext) (tag interface{}, err common.Error) {
	data := c.tagModelFactory.CreateTagCreate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	tag, err = c.tagApi.CreateTag(data)
	return
}

func (c *tagControllerImpl) UpdateTag(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	tagId, err := c.tagHttpHelper.ParseTagId(httpContext)

	if err != nil {
		return
	}

	data := c.tagModelFactory.CreateTagUpdate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	err = c.tagApi.UpdateTag(tagId, data)
	return
}

func (c *tagControllerImpl) DeleteTag(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	tagId, err := c.tagHttpHelper.ParseTagId(httpContext)

	if err != nil {
		return
	}

	err = c.tagApi.DeleteTag(tagId)
	return
}
