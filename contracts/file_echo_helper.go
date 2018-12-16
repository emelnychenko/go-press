package contracts

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/labstack/echo"
	"mime/multipart"
)

type (
	FileEchoHelper interface {
		ParseId(context echo.Context) (fileId *models.FileId, err error)
		GetFileHeader(context echo.Context) (formHeader *multipart.FileHeader, err error)
		OpenFormFile(formHeader *multipart.FileHeader) (formFile multipart.File, err error)
		PrepareFileDestination(context echo.Context) (getFileDestination PrepareFileDestination)
	}
)
