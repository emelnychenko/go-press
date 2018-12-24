package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindPostPictureRoutes(r contracts.Router, c contracts.PostPictureController) {
	r.AddRoute(http.MethodPut, fmt.Sprintf("/post/:%s/picture/:%s", helpers.PostIdParameterName, helpers.FileIdParameterName), c.ChangePostPicture)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/post/:%s/picture", helpers.PostIdParameterName), c.RemovePostPicture)
}
