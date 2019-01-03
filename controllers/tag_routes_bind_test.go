package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestTagRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindTagRoutes", func(t *testing.T) {
		var tagController contracts.TagController = new(tagControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/tags", gomock.AssignableToTypeOf(tagController.ListTags))
		router.EXPECT().AddRoute(http.MethodGet, "/tag/:tagId", gomock.AssignableToTypeOf(tagController.GetTag))
		router.EXPECT().AddRoute(http.MethodPost, "/tags", gomock.AssignableToTypeOf(tagController.CreateTag))
		router.EXPECT().AddRoute(http.MethodPost, "/tag/:tagId", gomock.AssignableToTypeOf(tagController.UpdateTag))
		router.EXPECT().AddRoute(http.MethodDelete, "/tag/:tagId", gomock.AssignableToTypeOf(tagController.DeleteTag))

		BindTagRoutes(router, tagController)
	})
}
