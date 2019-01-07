package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
)

type categoryControllerImpl struct {
	categoryHttpHelper   contracts.CategoryHttpHelper
	categoryModelFactory contracts.CategoryModelFactory
	categoryApi          contracts.CategoryApi
}

//NewCategoryController
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

//ListCategories
func (c *categoryControllerImpl) ListCategories(httpContext contracts.HttpContext) (paginationResult interface{}, err errors.Error) {
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

//GetCategoriesTree
func (c *categoryControllerImpl) GetCategoriesTree(httpContext contracts.HttpContext) (
	categoriesTree interface{}, err errors.Error,
) {
	categoriesTree, err = c.categoryApi.GetCategoriesTree()
	return
}

//GetCategory
func (c *categoryControllerImpl) GetCategory(httpContext contracts.HttpContext) (
	category interface{}, err errors.Error,
) {
	categoryId, err := c.categoryHttpHelper.ParseCategoryId(httpContext)

	if err != nil {
		return
	}

	category, err = c.categoryApi.GetCategory(categoryId)
	return
}

//GetCategoryTree
func (c *categoryControllerImpl) GetCategoryTree(httpContext contracts.HttpContext) (
	categoryTree interface{}, err errors.Error,
) {
	categoryId, err := c.categoryHttpHelper.ParseCategoryId(httpContext)

	if err != nil {
		return
	}

	categoryTree, err = c.categoryApi.GetCategoryTree(categoryId)
	return
}

//CreateCategory
func (c *categoryControllerImpl) CreateCategory(httpContext contracts.HttpContext) (category interface{}, err errors.Error) {
	data := c.categoryModelFactory.CreateCategoryCreate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	category, err = c.categoryApi.CreateCategory(data)
	return
}

//UpdateCategory
func (c *categoryControllerImpl) UpdateCategory(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
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

//ChangeCategoryParent
func (c *categoryControllerImpl) ChangeCategoryParent(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
	categoryId, err := c.categoryHttpHelper.ParseCategoryId(httpContext)

	if err != nil {
		return
	}

	parentCategoryId, err := c.categoryHttpHelper.ParseParentCategoryId(httpContext)

	if err != nil {
		return
	}

	err = c.categoryApi.ChangeCategoryParent(categoryId, parentCategoryId)
	return
}

//RemoveCategoryParent
func (c *categoryControllerImpl) RemoveCategoryParent(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
	categoryId, err := c.categoryHttpHelper.ParseCategoryId(httpContext)

	if err != nil {
		return
	}

	err = c.categoryApi.RemoveCategoryParent(categoryId)
	return
}

//DeleteCategory
func (c *categoryControllerImpl) DeleteCategory(httpContext contracts.HttpContext) (_ interface{}, err errors.Error) {
	categoryId, err := c.categoryHttpHelper.ParseCategoryId(httpContext)

	if err != nil {
		return
	}

	err = c.categoryApi.DeleteCategory(categoryId)
	return
}
