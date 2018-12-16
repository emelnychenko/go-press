package controllers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/labstack/echo"
	"net/http"
)

type FileController struct {
	fileEchoHelper   contracts.FileEchoHelper
	fileModelFactory contracts.FileModelFactory
	fileApi          contracts.FileApi
}

func NewFileController(
	fileEchoHelper contracts.FileEchoHelper,
	fileModelFactory contracts.FileModelFactory,
	fileApi contracts.FileApi,
) *FileController {
	return &FileController{
		fileEchoHelper,
		fileModelFactory,
		fileApi,
	}
}

func (c *FileController) ListFiles(context echo.Context) error {
	files, err := c.fileApi.ListFiles()

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	return context.JSON(http.StatusOK, files)
}

func (c *FileController) GetFile(context echo.Context) error {
	fileId, err := c.fileEchoHelper.ParseId(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	file, err2 := c.fileApi.GetFile(fileId)

	if err2 != nil {
		return context.JSON(err2.Code(), err2)
	}

	return context.JSON(http.StatusOK, file)
}

func (c *FileController) UploadFile(context echo.Context) error {
	fileHeader, err := c.fileEchoHelper.GetFileHeader(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	data := c.fileModelFactory.CreateFileUpload()
	data.Name = fileHeader.Filename
	data.Size = fileHeader.Size
	data.Type = fileHeader.Header.Get(echo.HeaderContentType)

	fileSource, err := c.fileEchoHelper.OpenFormFile(fileHeader)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	defer fileSource.Close()

	file, err2 := c.fileApi.UploadFile(fileSource, data)

	if err2 != nil {
		return context.JSON(err2.Code(), err2)
	}

	return context.JSON(http.StatusOK, file)
}

func (c *FileController) DownloadFile(context echo.Context) error {
	fileId, err := c.fileEchoHelper.ParseId(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	if err := c.fileApi.DownloadFile(
		fileId, c.fileEchoHelper.PrepareFileDestination(context),
	); err != nil {
		return context.JSON(err.Code(), err)
	}

	return nil
}

func (c *FileController) UpdateFile(context echo.Context) error {
	fileId, err := c.fileEchoHelper.ParseId(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	data := c.fileModelFactory.CreateFileUpdate()

	if err := context.Bind(data); err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	if err2 := c.fileApi.UpdateFile(fileId, data); err2 != nil {
		return context.JSON(err2.Code(), err2)
	}

	return context.JSON(http.StatusOK, nil)
}

func (c *FileController) DeleteFile(context echo.Context) error {
	fileId, err := c.fileEchoHelper.ParseId(context)

	if err != nil {
		return context.JSON(http.StatusBadRequest, err)
	}

	if err := c.fileApi.DeleteFile(fileId); err != nil {
		return context.JSON(err.Code(), err)
	}

	return context.JSON(http.StatusOK, nil)
}
