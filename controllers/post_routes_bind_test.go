package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestPostRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindPostRoutes", func(t *testing.T) {
		var postController contracts.PostController = new(postControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/posts", gomock.AssignableToTypeOf(postController.ListPosts))
		router.EXPECT().AddRoute(http.MethodGet, "/post/:postId", gomock.AssignableToTypeOf(postController.GetPost))
		router.EXPECT().AddRoute(http.MethodPost, "/posts", gomock.AssignableToTypeOf(postController.CreatePost))
		router.EXPECT().AddRoute(http.MethodPost, "/post/:postId", gomock.AssignableToTypeOf(postController.UpdatePost))
		router.EXPECT().AddRoute(http.MethodDelete, "/post/:postId", gomock.AssignableToTypeOf(postController.DeletePost))

		BindPostRoutes(router, postController)
	})
}
