package controllers

import (
	"fmt"
	"github.com/emelnychenko/go-press/helpers"
	"github.com/labstack/echo"
)

func BindFileController(e *echo.Echo, c *FileController) {
	e.GET("/v0/files", c.ListFiles)
	e.GET(fmt.Sprintf("/v0/file/:%s", helpers.FileIdParam), c.GetFile)
	e.GET(fmt.Sprintf("/v0/download/:%s", helpers.FileIdParam), c.DownloadFile)
	e.POST("/v0/upload", c.UploadFile)
	e.POST(fmt.Sprintf("/v0/file/:%s", helpers.FileIdParam), c.UpdateFile)
	e.DELETE(fmt.Sprintf("/v0/file/:%s", helpers.FileIdParam), c.DeleteFile)
}
