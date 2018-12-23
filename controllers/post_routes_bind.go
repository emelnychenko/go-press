package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindPostRoutes(r contracts.Router, c contracts.PostController) {
	r.AddRoute(http.MethodGet, "/v0/posts", c.ListPosts)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/v0/post/:%s", helpers.PostIdParameterName), c.GetPost)
	r.AddRoute(http.MethodPost, "/v0/posts", c.CreatePost)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/v0/post/:%s", helpers.PostIdParameterName), c.UpdatePost)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/v0/post/:%s", helpers.PostIdParameterName), c.DeletePost)
}
