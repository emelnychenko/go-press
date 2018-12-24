package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestPostAuthorRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindPostAuthorRoutes", func(t *testing.T) {
		var postAuthorController contracts.PostAuthorController = new(postAuthorControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodPut, "/post/:postId/author/:userId", gomock.AssignableToTypeOf(postAuthorController.ChangePostAuthor))

		BindPostAuthorRoutes(router, postAuthorController)
	})
}
