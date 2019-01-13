package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

//BindPostTagRoutes
func BindPostTagRoutes(router contracts.Router, postTagController contracts.PostTagController) {
	router.AddRoute(
		http.MethodGet,
		fmt.Sprintf("/post/:%s/tags", helpers.PostIdParameterName),
		postTagController.ListPostTags,
	)

	router.AddRoute(
		http.MethodPut,
		fmt.Sprintf("/post/:%s/tag/:%s", helpers.PostIdParameterName, helpers.TagIdParameterName),
		postTagController.AddPostTag,
	)

	router.AddRoute(
		http.MethodDelete,
		fmt.Sprintf("/post/:%s/tag/:%s", helpers.PostIdParameterName, helpers.TagIdParameterName),
		postTagController.RemovePostTag,
	)
}
