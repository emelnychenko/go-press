package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindCategoryRoutes(r contracts.Router, c contracts.CategoryController) {
	r.AddRoute(http.MethodGet, "/categories", c.ListCategories)
	r.AddRoute(http.MethodGet, "/categories/tree", c.GetCategoriesTree)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/category/:%s", helpers.CategoryIdParameterName), c.GetCategory)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/category/:%s/tree", helpers.CategoryIdParameterName), c.GetCategoryTree)
	r.AddRoute(http.MethodPost, "/categories", c.CreateCategory)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/category/:%s", helpers.CategoryIdParameterName), c.UpdateCategory)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/category/:%s", helpers.CategoryIdParameterName), c.DeleteCategory)
}
