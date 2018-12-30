package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type postControllerImpl struct {
	postHttpHelper         contracts.PostHttpHelper
	postModelFactory       contracts.PostModelFactory
	postStatusValidator    contracts.PostStatusValidator
	postApi                contracts.PostApi
}

func NewPostController(
	postHttpHelper contracts.PostHttpHelper,
	postModelFactory contracts.PostModelFactory,
	postStatusValidator contracts.PostStatusValidator,
	postApi contracts.PostApi,
) (postController contracts.PostController) {
	return &postControllerImpl{
		postHttpHelper,
		postModelFactory,
		postStatusValidator,
		postApi,
	}
}

func (c *postControllerImpl) ListPosts(httpContext contracts.HttpContext) (paginationResult interface{}, err common.Error) {
	postPaginationQuery := c.postModelFactory.CreatePostPaginationQuery()

	if err = httpContext.BindModel(postPaginationQuery.PaginationQuery); err != nil {
		return
	}

	if err = httpContext.BindModel(postPaginationQuery); err != nil {
		return
	}

	paginationResult, err = c.postApi.ListPosts(postPaginationQuery)
	return
}

func (c *postControllerImpl) GetPost(httpContext contracts.HttpContext) (post interface{}, err common.Error) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	post, err = c.postApi.GetPost(postId)
	return
}

func (c *postControllerImpl) CreatePost(httpContext contracts.HttpContext) (post interface{}, err common.Error) {
	data := c.postModelFactory.CreatePostCreate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	if err = c.postStatusValidator.ValidatePostCreate(data); err != nil {
		return
	}

	subject := models.NewSystemUser() // TODO: Replace stub
	post, err = c.postApi.CreatePost(subject, data)
	return
}

func (c *postControllerImpl) UpdatePost(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	data := c.postModelFactory.CreatePostUpdate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	if err = c.postStatusValidator.ValidatePostUpdate(data); err != nil {
		return
	}

	err = c.postApi.UpdatePost(postId, data)
	return
}

func (c *postControllerImpl) DeletePost(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	err = c.postApi.DeletePost(postId)
	return
}
