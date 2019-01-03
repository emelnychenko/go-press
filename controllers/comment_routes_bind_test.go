package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestCommentRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindCommentRoutes", func(t *testing.T) {
		var commentController contracts.CommentController = new(commentControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/comments", gomock.AssignableToTypeOf(commentController.ListComments))
		router.EXPECT().AddRoute(http.MethodGet, "/comment/:commentId", gomock.AssignableToTypeOf(commentController.GetComment))
		router.EXPECT().AddRoute(http.MethodPost, "/comments", gomock.AssignableToTypeOf(commentController.CreateComment))
		router.EXPECT().AddRoute(http.MethodPost, "/comment/:commentId", gomock.AssignableToTypeOf(commentController.UpdateComment))
		router.EXPECT().AddRoute(http.MethodDelete, "/comment/:commentId", gomock.AssignableToTypeOf(commentController.DeleteComment))

		BindCommentRoutes(router, commentController)
	})
}
