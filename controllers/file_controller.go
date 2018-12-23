package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/labstack/echo"
)

type fileControllerImpl struct {
	fileHttpHelper   contracts.FileHttpHelper
	fileModelFactory contracts.FileModelFactory
	fileApi          contracts.FileApi
}

func NewFileController(
	fileHttpHelper contracts.FileHttpHelper,
	fileModelFactory contracts.FileModelFactory,
	fileApi contracts.FileApi,
) (fileController contracts.FileController) {
	return &fileControllerImpl{
		fileHttpHelper,
		fileModelFactory,
		fileApi,
	}
}

func (c *fileControllerImpl) ListFiles(httpContext contracts.HttpContext) (files interface{}, err common.Error) {
	files, err = c.fileApi.ListFiles()
	return
}

func (c *fileControllerImpl) GetFile(httpContext contracts.HttpContext) (file interface{}, err common.Error) {
	fileId, err := c.fileHttpHelper.ParseFileId(httpContext)

	if err != nil {
		return
	}

	file, err = c.fileApi.GetFile(fileId)
	return
}

func (c *fileControllerImpl) UploadFile(httpContext contracts.HttpContext) (file interface{}, err common.Error) {
	fileHeader, err := c.fileHttpHelper.GetFileHeader(httpContext)

	if err != nil {
		return
	}

	data := c.fileModelFactory.CreateFileUpload()
	data.Name = fileHeader.Filename
	data.Size = fileHeader.Size
	data.Type = fileHeader.Header.Get(echo.HeaderContentType)

	fileSource, err := c.fileHttpHelper.OpenFormFile(fileHeader)

	if err != nil {
		return
	}

	defer fileSource.Close()

	file, err = c.fileApi.UploadFile(fileSource, data)
	return
}

func (c *fileControllerImpl) DownloadFile(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	fileId, err := c.fileHttpHelper.ParseFileId(httpContext)

	if err != nil {
		return
	}

	err = c.fileApi.DownloadFile(fileId, c.fileHttpHelper.PrepareFileDestination(httpContext))
	return
}

func (c *fileControllerImpl) UpdateFile(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	fileId, err := c.fileHttpHelper.ParseFileId(httpContext)

	if err != nil {
		return
	}

	data := c.fileModelFactory.CreateFileUpdate()

	if err = httpContext.BindModel(data); err != nil {
		return
	}

	err = c.fileApi.UpdateFile(fileId, data)
	return
}

func (c *fileControllerImpl) DeleteFile(httpContext contracts.HttpContext) (_ interface{}, err common.Error) {
	fileId, err := c.fileHttpHelper.ParseFileId(httpContext)

	if err != nil {
		return
	}

	err = c.fileApi.DeleteFile(fileId)
	return
}
