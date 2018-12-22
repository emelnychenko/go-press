package helpers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
	"net/http"
)

const (
	FileIdParam  = "fileId"
	FileFormFile = "file"
)

type (
	fileEchoHelperImpl struct {
	}
)

func NewFileEchoHelper() contracts.FileEchoHelper {
	return new(fileEchoHelperImpl)
}

func (*fileEchoHelperImpl) ParseId(context echo.Context) (*models.FileId, error) {
	return common.ParseModelId(context.Param(FileIdParam))
}

func (*fileEchoHelperImpl) GetFileHeader(context echo.Context) (*multipart.FileHeader, error) {
	return context.FormFile(FileFormFile)
}

func (*fileEchoHelperImpl) OpenFormFile(formHeader *multipart.FileHeader) (multipart.File, error) {
	return formHeader.Open()
}

func (*fileEchoHelperImpl) PrepareFileDestination(context echo.Context) contracts.PrepareFileDestination {
	// TODO: Caching
	return func(file *models.File) (destination io.Writer) {
		response := context.Response()
		//request := context.Request()
		//
		//var lastModified *time.Time
		//if nil != file.Updated {
		//	lastModified = file.Updated
		//} else {
		//	lastModified = file.Created
		//}
		//response.Header().Set("Cache-Control", "max-age=290304000, public")
		//response.Header().Set(echo.HeaderLastModified, lastModified.Format(http.TimeFormat))
		//
		//ifModifiedSinceHeader := request.Header.Get(echo.HeaderIfModifiedSince)
		//if "" != ifModifiedSinceHeader {
		//	if ifModifiedSince, err := time.Parse(http.TimeFormat, ifModifiedSinceHeader); nil != err {
		//		if ifModifiedSince.Equal(*lastModified) {
		//			_ = context.NoContent(http.StatusNotModified)
		//			return
		//		}
		//	}
		//}

		response.Header().Set(echo.HeaderContentType, file.Type)
		response.Header().Set(echo.HeaderContentLength, string(file.Size))
		response.WriteHeader(http.StatusOK)
		return response
	}
}
