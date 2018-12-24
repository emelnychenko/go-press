package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestFileRoutesBind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("BindFileRoutes", func(t *testing.T) {
		var fileController contracts.FileController = new(fileControllerImpl)
		router := mocks.NewMockRouter(ctrl)

		router.EXPECT().AddRoute(http.MethodGet, "/files", gomock.AssignableToTypeOf(fileController.ListFiles))
		router.EXPECT().AddRoute(http.MethodGet, "/file/:fileId", gomock.AssignableToTypeOf(fileController.GetFile))
		router.EXPECT().AddRoute(http.MethodGet, "/download/:fileId", gomock.AssignableToTypeOf(fileController.DownloadFile))
		router.EXPECT().AddRoute(http.MethodPost, "/upload", gomock.AssignableToTypeOf(fileController.UploadFile))
		router.EXPECT().AddRoute(http.MethodPost, "/file/:fileId", gomock.AssignableToTypeOf(fileController.UpdateFile))
		router.EXPECT().AddRoute(http.MethodDelete, "/file/:fileId", gomock.AssignableToTypeOf(fileController.DeleteFile))

		BindFileRoutes(router, fileController)
	})
}
