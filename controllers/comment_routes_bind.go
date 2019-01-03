package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindCommentRoutes(r contracts.Router, c contracts.CommentController) {
	r.AddRoute(http.MethodGet, "/comments", c.ListComments)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/comment/:%s", helpers.CommentIdParameterName), c.GetComment)
	r.AddRoute(http.MethodPost, "/comments", c.CreateComment)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/comment/:%s", helpers.CommentIdParameterName), c.UpdateComment)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/comment/:%s", helpers.CommentIdParameterName), c.DeleteComment)
}
