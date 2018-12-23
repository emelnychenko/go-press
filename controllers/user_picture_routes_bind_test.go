package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestUserPictureRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindUserPictureRoutes", func(t *testing.T) {
		var userPictureController contracts.UserPictureController = new(userPictureControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodPut, "/v0/user/:userId/picture/:fileId", gomock.AssignableToTypeOf(userPictureController.ChangeUserPicture))
		router.EXPECT().AddRoute(http.MethodDelete, "/v0/user/:userId/picture", gomock.AssignableToTypeOf(userPictureController.RemoveUserPicture))

		BindUserPictureRoutes(router, userPictureController)
	})
}
