package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindPollRoutes(r contracts.Router, c contracts.PollController) {
	r.AddRoute(http.MethodGet, "/polls", c.ListPolls)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/poll/:%s", helpers.PollIdParameterName), c.GetPoll)
	r.AddRoute(http.MethodPost, "/polls", c.CreatePoll)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/poll/:%s", helpers.PollIdParameterName), c.UpdatePoll)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/poll/:%s", helpers.PollIdParameterName), c.DeletePoll)
}
