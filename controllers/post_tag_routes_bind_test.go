package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestPostTagRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindPostTagRoutes", func(t *testing.T) {
		var postTagController contracts.PostTagController = new(postTagControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/post/:postId/tags", gomock.Any())
		router.EXPECT().AddRoute(http.MethodPut, "/post/:postId/tag/:tagId", gomock.Any())
		router.EXPECT().AddRoute(http.MethodDelete, "/post/:postId/tag/:tagId", gomock.Any())

		BindPostTagRoutes(router, postTagController)
	})
}
