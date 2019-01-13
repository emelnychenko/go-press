package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
)

type (
	postCategoryControllerImpl struct {
		postHttpHelper       contracts.PostHttpHelper
		categoryHttpHelper   contracts.CategoryHttpHelper
		categoryModelFactory contracts.CategoryModelFactory
		postCategoryApi      contracts.PostCategoryApi
	}
)

//NewPostCategoryController
func NewPostCategoryController(
	postHttpHelper contracts.PostHttpHelper,
	categoryHttpHelper contracts.CategoryHttpHelper,
	categoryModelFactory contracts.CategoryModelFactory,
	postCategoryApi contracts.PostCategoryApi,
) contracts.PostCategoryController {
	return &postCategoryControllerImpl{
		postHttpHelper,
		categoryHttpHelper,
		categoryModelFactory,
		postCategoryApi,
	}
}

//ListPostCategories
func (c *postCategoryControllerImpl) ListPostCategories(httpContext contracts.HttpContext) (
	paginationResult interface{}, err errors.Error,
) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	categoryPaginationQuery := c.categoryModelFactory.CreateCategoryPaginationQuery()

	if err = httpContext.BindModel(categoryPaginationQuery.PaginationQuery); err != nil {
		return
	}

	if err = httpContext.BindModel(categoryPaginationQuery); err != nil {
		return
	}

	paginationResult, err = c.postCategoryApi.ListPostCategories(postId, categoryPaginationQuery)
	return
}

//AddPostCategory
func (c *postCategoryControllerImpl) AddPostCategory(httpContext contracts.HttpContext) (
	_ interface{}, err errors.Error,
) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	categoryId, err := c.categoryHttpHelper.ParseCategoryId(httpContext)

	if err != nil {
		return
	}

	err = c.postCategoryApi.AddPostCategory(postId, categoryId)
	return
}

//RemovePostCategory
func (c *postCategoryControllerImpl) RemovePostCategory(httpContext contracts.HttpContext) (
	_ interface{}, err errors.Error,
) {
	postId, err := c.postHttpHelper.ParsePostId(httpContext)

	if err != nil {
		return
	}

	categoryId, err := c.categoryHttpHelper.ParseCategoryId(httpContext)

	if err != nil {
		return
	}

	err = c.postCategoryApi.RemovePostCategory(postId, categoryId)
	return
}
