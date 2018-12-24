package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindFileRoutes(r contracts.Router, c contracts.FileController) {
	r.AddRoute(http.MethodGet, "/files", c.ListFiles)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/file/:%s", helpers.FileIdParameterName), c.GetFile)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/download/:%s", helpers.FileIdParameterName), c.DownloadFile)
	r.AddRoute(http.MethodPost, "/upload", c.UploadFile)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/file/:%s", helpers.FileIdParameterName), c.UpdateFile)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/file/:%s", helpers.FileIdParameterName), c.DeleteFile)
}
