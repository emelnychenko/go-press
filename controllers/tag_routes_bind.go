package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindTagRoutes(r contracts.Router, c contracts.TagController) {
	r.AddRoute(http.MethodGet, "/tags", c.ListTags)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/tag/:%s", helpers.TagIdParameterName), c.GetTag)
	r.AddRoute(http.MethodPost, "/tags", c.CreateTag)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/tag/:%s", helpers.TagIdParameterName), c.UpdateTag)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/tag/:%s", helpers.TagIdParameterName), c.DeleteTag)
}
