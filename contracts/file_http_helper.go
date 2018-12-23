package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
	"mime/multipart"
)

type (
	FileHttpHelper interface {
		ParseFileId(httpContext HttpContext) (fileId *models.FileId, err common.Error)
		GetFileHeader(httpContext HttpContext) (formHeader *multipart.FileHeader, err common.Error)
		OpenFormFile(httpContext *multipart.FileHeader) (formFile multipart.File, err common.Error)
		PrepareFileDestination(httpContext HttpContext) (getFileDestination PrepareFileDestination)
	}
)
