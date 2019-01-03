package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestChannelRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindChannelRoutes", func(t *testing.T) {
		var channelController contracts.ChannelController = new(channelControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/channels", gomock.AssignableToTypeOf(channelController.ListChannels))
		router.EXPECT().AddRoute(http.MethodGet, "/channel/:channelId", gomock.AssignableToTypeOf(channelController.GetChannel))
		router.EXPECT().AddRoute(http.MethodPost, "/channels", gomock.AssignableToTypeOf(channelController.CreateChannel))
		router.EXPECT().AddRoute(http.MethodPost, "/channel/:channelId", gomock.AssignableToTypeOf(channelController.UpdateChannel))
		router.EXPECT().AddRoute(http.MethodDelete, "/channel/:channelId", gomock.AssignableToTypeOf(channelController.DeleteChannel))

		BindChannelRoutes(router, channelController)
	})
}
