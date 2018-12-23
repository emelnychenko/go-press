package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/helpers"
	"net/http"
)

func BindFileRoutes(r contracts.Router, c contracts.FileController) {
	r.AddRoute(http.MethodGet, "/v0/files", c.ListFiles)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/v0/file/:%s", helpers.FileIdParameterName), c.GetFile)
	r.AddRoute(http.MethodGet, fmt.Sprintf("/v0/download/:%s", helpers.FileIdParameterName), c.DownloadFile)
	r.AddRoute(http.MethodPost, "/v0/upload", c.UploadFile)
	r.AddRoute(http.MethodPost, fmt.Sprintf("/v0/file/:%s", helpers.FileIdParameterName), c.UpdateFile)
	r.AddRoute(http.MethodDelete, fmt.Sprintf("/v0/file/:%s", helpers.FileIdParameterName), c.DeleteFile)
}
