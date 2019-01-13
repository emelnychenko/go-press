package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
)

type (
	postTagControllerImpl struct {
		postHttpHelper       contracts.PostHttpHelper
		tagHttpHelper   contracts.TagHttpHelper
		tagModelFactory contracts.TagModelFactory
		postTagApi      contracts.PostTagApi
	}
)

//NewPostTagController
func NewPostTagController(
	postHttpHelper contracts.PostHttpHelper,
	tagHttpHelper contracts.TagHttpHelper,
	tagModelFactory contracts.TagModelFactory,
	postTagApi contracts.PostTagApi,
) contracts.PostTagController {
	return &postTagControllerImpl{
		postHttpHelper,
		tagHttpHelper,
		tagModelFactory,
		postTagApi,
	}
}

//ListPostTags
func (c *postTagControllerImpl) ListPostTags(httpContext contracts.HttpContext) (
	paginationResult interface{}, err errors.Error,
) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	tagPaginationQuery := c.tagModelFactory.CreateTagPaginationQuery()

	if err = httpContext.BindModel(tagPaginationQuery.PaginationQuery); err != nil {
		return
	}

	if err = httpContext.BindModel(tagPaginationQuery); err != nil {
		return
	}

	paginationResult, err = c.postTagApi.ListPostTags(postId, tagPaginationQuery)
	return
}

//AddPostTag
func (c *postTagControllerImpl) AddPostTag(httpContext contracts.HttpContext) (
	_ interface{}, err errors.Error,
) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	tagId, err := c.tagHttpHelper.ParseTagId(httpContext)

	if err != nil {
		return
	}

	err = c.postTagApi.AddPostTag(postId, tagId)
	return
}

//RemovePostTag
func (c *postTagControllerImpl) RemovePostTag(httpContext contracts.HttpContext) (
	_ interface{}, err errors.Error,
) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	tagId, err := c.tagHttpHelper.ParseTagId(httpContext)

	if err != nil {
		return
	}

	err = c.postTagApi.RemovePostTag(postId, tagId)
	return
}
