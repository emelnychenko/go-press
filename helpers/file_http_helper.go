package helpers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"github.com/labstack/echo"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

const (
	FileIdParameterName = "fileId"
	FileFormFileName    = "file"
)

type (
	fileHttpHelperImpl struct {
	}
)

func NewFileHttpHelper() contracts.FileHttpHelper {
	return new(fileHttpHelperImpl)
}

func (*fileHttpHelperImpl) ParseFileId(httpContext contracts.HttpContext) (*models.FileId, errors.Error) {
	return models.ParseModelId(httpContext.Parameter(FileIdParameterName))
}

func (*fileHttpHelperImpl) GetFileHeader(httpContext contracts.HttpContext) (*multipart.FileHeader, errors.Error) {
	return httpContext.FormFile(FileFormFileName)
}

func (*fileHttpHelperImpl) OpenFormFile(formHeader *multipart.FileHeader) (file multipart.File, err errors.Error) {
	file, formHeaderErr := formHeader.Open()

	if nil != formHeaderErr {
		err = errors.NewSystemErrorFromBuiltin(formHeaderErr)
	}

	return
}

func (*fileHttpHelperImpl) PrepareFileDestination(httpContext contracts.HttpContext) contracts.PrepareFileDestination {
	// TODO: Caching
	return func(file *models.File) (destination io.Writer) {
		response := httpContext.Response()
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

		contentLength := strconv.FormatInt(file.Size, 10)
		response.Header().Set(echo.HeaderContentType, file.Type)
		response.Header().Set(echo.HeaderContentLength, contentLength)
		response.WriteHeader(http.StatusOK)
		return response
	}
}
