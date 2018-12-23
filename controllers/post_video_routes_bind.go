package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindPostVideoRoutes(r contracts.Router, c contracts.PostVideoController) {
	r.AddRoute(http.MethodPut, fmt.Sprintf("/v0/post/:%s/video/:%s", helpers.PostIdParameterName, helpers.FileIdParameterName), c.ChangePostVideo)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/v0/post/:%s/video", helpers.PostIdParameterName), c.RemovePostVideo)
}
