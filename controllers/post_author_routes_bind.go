package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindPostAuthorRoutes(r contracts.Router, c contracts.PostAuthorController) {
	r.AddRoute(http.MethodPut, fmt.Sprintf("/post/:%s/author/:%s", helpers.PostIdParameterName, helpers.UserIdParameterName), c.ChangePostAuthor)
}
