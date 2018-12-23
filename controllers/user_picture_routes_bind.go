package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindUserPictureRoutes(r contracts.Router, c contracts.UserPictureController) {
	r.AddRoute(http.MethodPut, fmt.Sprintf("/v0/user/:%s/picture/:%s", helpers.UserIdParameterName, helpers.FileIdParameterName), c.ChangeUserPicture)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/v0/user/:%s/picture", helpers.UserIdParameterName), c.RemoveUserPicture)
}
