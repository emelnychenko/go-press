package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestCategoryRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindCategoryRoutes", func(t *testing.T) {
		var categoryController contracts.CategoryController = new(categoryControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/categories", gomock.AssignableToTypeOf(categoryController.ListCategories))
		router.EXPECT().AddRoute(http.MethodGet, "/categories/tree", gomock.AssignableToTypeOf(categoryController.GetCategoriesTree))
		router.EXPECT().AddRoute(http.MethodGet, "/category/:categoryId", gomock.AssignableToTypeOf(categoryController.GetCategory))
		router.EXPECT().AddRoute(http.MethodGet, "/category/:categoryId/tree", gomock.AssignableToTypeOf(categoryController.GetCategoryTree))
		router.EXPECT().AddRoute(http.MethodPost, "/categories", gomock.AssignableToTypeOf(categoryController.CreateCategory))
		router.EXPECT().AddRoute(http.MethodPost, "/category/:categoryId", gomock.AssignableToTypeOf(categoryController.UpdateCategory))
		router.EXPECT().AddRoute(http.MethodDelete, "/category/:categoryId", gomock.AssignableToTypeOf(categoryController.DeleteCategory))

		BindCategoryRoutes(router, categoryController)
	})
}
