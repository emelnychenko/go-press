package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"io"
)

type (
	PrepareFileDestination func(file *models.File) (fileDestination io.Writer)

	FileApi interface {
		ListFiles(filePaginationQuery *models.FilePaginationQuery) (*models.PaginationResult, errors.Error)
		GetFile(fileId *models.FileId) (file *models.File, err errors.Error)
		UploadFile(fileSource io.Reader, data *models.FileUpload) (file *models.File, err errors.Error)
		DownloadFile(fileId *models.FileId, prepareFileDestination PrepareFileDestination) (err errors.Error)
		UpdateFile(fileId *models.FileId, data *models.FileUpdate) (err errors.Error)
		DeleteFile(fileId *models.FileId) (err errors.Error)
	}
)
