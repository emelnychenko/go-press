package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestPollRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindPollRoutes", func(t *testing.T) {
		var pollController contracts.PollController = new(pollControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/polls", gomock.AssignableToTypeOf(pollController.ListPolls))
		router.EXPECT().AddRoute(http.MethodGet, "/poll/:pollId", gomock.AssignableToTypeOf(pollController.GetPoll))
		router.EXPECT().AddRoute(http.MethodPost, "/polls", gomock.AssignableToTypeOf(pollController.CreatePoll))
		router.EXPECT().AddRoute(http.MethodPost, "/poll/:pollId", gomock.AssignableToTypeOf(pollController.UpdatePoll))
		router.EXPECT().AddRoute(http.MethodDelete, "/poll/:pollId", gomock.AssignableToTypeOf(pollController.DeletePoll))

		BindPollRoutes(router, pollController)
	})
}
