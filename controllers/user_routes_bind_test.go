package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestUserRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindUserRoutes", func(t *testing.T) {
		var userController contracts.UserController = new(userControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/users", gomock.AssignableToTypeOf(userController.ListUsers))
		router.EXPECT().AddRoute(http.MethodGet, "/user/:userId", gomock.AssignableToTypeOf(userController.GetUser))
		router.EXPECT().AddRoute(http.MethodPost, "/users", gomock.AssignableToTypeOf(userController.CreateUser))
		router.EXPECT().AddRoute(http.MethodPost, "/user/:userId", gomock.AssignableToTypeOf(userController.UpdateUser))
		router.EXPECT().AddRoute(http.MethodPost, "/user/:userId/identity", gomock.AssignableToTypeOf(userController.ChangeUserIdentity))
		router.EXPECT().AddRoute(http.MethodPost, "/user/:userId/password", gomock.AssignableToTypeOf(userController.ChangeUserPassword))
		router.EXPECT().AddRoute(http.MethodDelete, "/user/:userId", gomock.AssignableToTypeOf(userController.DeleteUser))

		BindUserRoutes(router, userController)
	})
}
