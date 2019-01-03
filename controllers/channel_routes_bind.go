package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindChannelRoutes(r contracts.Router, c contracts.ChannelController) {
	r.AddRoute(http.MethodGet, "/channels", c.ListChannels)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/channel/:%s", helpers.ChannelIdParameterName), c.GetChannel)
	r.AddRoute(http.MethodPost, "/channels", c.CreateChannel)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/channel/:%s", helpers.ChannelIdParameterName), c.UpdateChannel)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/channel/:%s", helpers.ChannelIdParameterName), c.DeleteChannel)
}
