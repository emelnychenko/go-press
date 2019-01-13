package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestPostCategoryRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindPostCategoryRoutes", func(t *testing.T) {
		var postCategoryController contracts.PostCategoryController = new(postCategoryControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/post/:postId/categories", gomock.Any())
		router.EXPECT().AddRoute(http.MethodPut, "/post/:postId/category/:categoryId", gomock.Any())
		router.EXPECT().AddRoute(http.MethodDelete, "/post/:postId/category/:categoryId", gomock.Any())

		BindPostCategoryRoutes(router, postCategoryController)
	})
}
