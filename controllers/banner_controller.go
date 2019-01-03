package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
)

type bannerControllerImpl struct {
	bannerHttpHelper   contracts.BannerHttpHelper
	bannerModelFactory contracts.BannerModelFactory
	bannerApi          contracts.BannerApi
}

func NewBannerController(
	bannerHttpHelper contracts.BannerHttpHelper,
	bannerModelFactory contracts.BannerModelFactory,
	bannerApi contracts.BannerApi,
) (bannerController contracts.BannerController) {
	return &bannerControllerImpl{
		bannerHttpHelper,
		bannerModelFactory,
		bannerApi,
	}
}

func (c *bannerControllerImpl) ListBanners(httpContext contracts.HttpContext) (paginationResult interface{}, err common.Error) {
	bannerPaginationQuery := c.bannerModelFactory.CreateBannerPaginationQuery()

	if err = httpContext.BindModel(bannerPaginationQuery.PaginationQuery); err != nil {
		return
	}

	if err = httpContext.BindModel(bannerPaginationQuery); err != nil {
		return
	}

	paginationResult, err = c.bannerApi.ListBanners(bannerPaginationQuery)
	return
}

func (c *bannerControllerImpl) GetBanner(httpContext contracts.HttpContext) (banner interface{}, err common.Error) {
	bannerId, err := c.bannerHttpHelper.ParseBannerId(httpContext)

	if err != nil {
		return
	}

	banner, err = c.bannerApi.GetBanner(bannerId)
	return
}

func (c *bannerControllerImpl) CreateBanner(httpContext contracts.HttpContext) (banner interface{}, err common.Error) {
	data := c.bannerModelFactory.CreateBannerCreate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	banner, err = c.bannerApi.CreateBanner(data)
	return
}

func (c *bannerControllerImpl) UpdateBanner(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	bannerId, err := c.bannerHttpHelper.ParseBannerId(httpContext)

	if err != nil {
		return
	}

	data := c.bannerModelFactory.CreateBannerUpdate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	err = c.bannerApi.UpdateBanner(bannerId, data)
	return
}

func (c *bannerControllerImpl) DeleteBanner(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	bannerId, err := c.bannerHttpHelper.ParseBannerId(httpContext)

	if err != nil {
		return
	}

	err = c.bannerApi.DeleteBanner(bannerId)
	return
}
