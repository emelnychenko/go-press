package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

//BindPostCategoryRoutes
func BindPostCategoryRoutes(router contracts.Router, postCategoryController contracts.PostCategoryController) {
	router.AddRoute(
		http.MethodGet,
		fmt.Sprintf("/post/:%s/categories", helpers.PostIdParameterName),
		postCategoryController.ListPostCategories,
	)

	router.AddRoute(
		http.MethodPut,
		fmt.Sprintf("/post/:%s/category/:%s", helpers.PostIdParameterName, helpers.CategoryIdParameterName),
		postCategoryController.AddPostCategory,
	)

	router.AddRoute(
		http.MethodDelete,
		fmt.Sprintf("/post/:%s/category/:%s", helpers.PostIdParameterName, helpers.CategoryIdParameterName),
		postCategoryController.RemovePostCategory,
	)
}
