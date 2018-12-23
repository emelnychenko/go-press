package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestPostVideoRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindPostVideoRoutes", func(t *testing.T) {
		var postVideoController contracts.PostVideoController = new(postVideoControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodPut, "/v0/post/:postId/video/:fileId", gomock.AssignableToTypeOf(postVideoController.ChangePostVideo))
		router.EXPECT().AddRoute(http.MethodDelete, "/v0/post/:postId/video", gomock.AssignableToTypeOf(postVideoController.RemovePostVideo))

		BindPostVideoRoutes(router, postVideoController)
	})
}
