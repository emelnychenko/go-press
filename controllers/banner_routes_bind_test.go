package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestBannerRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindBannerRoutes", func(t *testing.T) {
		var bannerController contracts.BannerController = new(bannerControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/banners", gomock.AssignableToTypeOf(bannerController.ListBanners))
		router.EXPECT().AddRoute(http.MethodGet, "/banner/:bannerId", gomock.AssignableToTypeOf(bannerController.GetBanner))
		router.EXPECT().AddRoute(http.MethodPost, "/banners", gomock.AssignableToTypeOf(bannerController.CreateBanner))
		router.EXPECT().AddRoute(http.MethodPost, "/banner/:bannerId", gomock.AssignableToTypeOf(bannerController.UpdateBanner))
		router.EXPECT().AddRoute(http.MethodDelete, "/banner/:bannerId", gomock.AssignableToTypeOf(bannerController.DeleteBanner))

		BindBannerRoutes(router, bannerController)
	})
}
