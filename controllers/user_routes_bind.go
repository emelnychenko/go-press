package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindUserRoutes(r contracts.Router, c contracts.UserController) {
	r.AddRoute(http.MethodGet, "/users", c.ListUsers)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/user/:%s", helpers.UserIdParameterName), c.GetUser)
	r.AddRoute(http.MethodPost, "/users", c.CreateUser)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/user/:%s", helpers.UserIdParameterName), c.UpdateUser)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/user/:%s/identity", helpers.UserIdParameterName), c.ChangeUserIdentity)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/user/:%s/password", helpers.UserIdParameterName), c.ChangeUserPassword)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/user/:%s", helpers.UserIdParameterName), c.DeleteUser)
}