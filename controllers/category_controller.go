package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
)

type categoryControllerImpl struct {
	categoryHttpHelper   contracts.CategoryHttpHelper
	categoryModelFactory contracts.CategoryModelFactory
	categoryApi          contracts.CategoryApi
}

func NewCategoryController(
	categoryHttpHelper contracts.CategoryHttpHelper,
	categoryModelFactory contracts.CategoryModelFactory,
	categoryApi contracts.CategoryApi,
) (categoryController contracts.CategoryController) {
	return &categoryControllerImpl{
		categoryHttpHelper,
		categoryModelFactory,
		categoryApi,
	}
}

func (c *categoryControllerImpl) ListCategories(httpContext contracts.HttpContext) (paginationResult interface{}, err common.Error) {
	categoryPaginationQuery := c.categoryModelFactory.CreateCategoryPaginationQuery()

	if err = httpContext.BindModel(categoryPaginationQuery.PaginationQuery); err != nil {
		return
	}

	if err = httpContext.BindModel(categoryPaginationQuery); err != nil {
		return
	}

	paginationResult, err = c.categoryApi.ListCategories(categoryPaginationQuery)
	return
}

func (c *categoryControllerImpl) GetCategory(httpContext contracts.HttpContext) (category interface{}, err common.Error) {
	categoryId, err := c.categoryHttpHelper.ParseCategoryId(httpContext)

	if err != nil {
		return
	}

	category, err = c.categoryApi.GetCategory(categoryId)
	return
}

func (c *categoryControllerImpl) CreateCategory(httpContext contracts.HttpContext) (category interface{}, err common.Error) {
	data := c.categoryModelFactory.CreateCategoryCreate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	category, err = c.categoryApi.CreateCategory(data)
	return
}

func (c *categoryControllerImpl) UpdateCategory(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	categoryId, err := c.categoryHttpHelper.ParseCategoryId(httpContext)

	if err != nil {
		return
	}

	data := c.categoryModelFactory.CreateCategoryUpdate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	err = c.categoryApi.UpdateCategory(categoryId, data)
	return
}

func (c *categoryControllerImpl) DeleteCategory(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	categoryId, err := c.categoryHttpHelper.ParseCategoryId(httpContext)

	if err != nil {
		return
	}

	err = c.categoryApi.DeleteCategory(categoryId)
	return
}
