package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestPostPictureRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindPostPictureRoutes", func(t *testing.T) {
		var postPictureController contracts.PostPictureController = new(postPictureControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodPut, "/v0/post/:postId/picture/:fileId", gomock.AssignableToTypeOf(postPictureController.ChangePostPicture))
		router.EXPECT().AddRoute(http.MethodDelete, "/v0/post/:postId/picture", gomock.AssignableToTypeOf(postPictureController.RemovePostPicture))

		BindPostPictureRoutes(router, postPictureController)
	})
}
