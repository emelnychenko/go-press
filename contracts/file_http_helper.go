package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"mime/multipart"
)

type (
	FileHttpHelper interface {
		ParseFileId(httpContext HttpContext) (fileId *models.FileId, err errors.Error)
		GetFileHeader(httpContext HttpContext) (formHeader *multipart.FileHeader, err errors.Error)
		OpenFormFile(httpContext *multipart.FileHeader) (formFile multipart.File, err errors.Error)
		PrepareFileDestination(httpContext HttpContext) (getFileDestination PrepareFileDestination)
	}
)
