package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindBannerRoutes(r contracts.Router, c contracts.BannerController) {
	r.AddRoute(http.MethodGet, "/banners", c.ListBanners)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/banner/:%s", helpers.BannerIdParameterName), c.GetBanner)
	r.AddRoute(http.MethodPost, "/banners", c.CreateBanner)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/banner/:%s", helpers.BannerIdParameterName), c.UpdateBanner)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/banner/:%s", helpers.BannerIdParameterName), c.DeleteBanner)
}
