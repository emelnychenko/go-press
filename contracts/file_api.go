package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
	"io"
)

type (
	PrepareFileDestination func(file *models.File) (fileDestination io.Writer)

	FileApi interface {
		ListFiles(
			filePaginationQuery *models.FilePaginationQuery,
		) (paginationResult *models.PaginationResult, err common.Error)
		GetFile(fileId *models.FileId) (file *models.File, err common.Error)
		UploadFile(fileSource io.Reader, data *models.FileUpload) (file *models.File, err common.Error)
		DownloadFile(fileId *models.FileId, prepareFileDestination PrepareFileDestination) (err common.Error)
		UpdateFile(fileId *models.FileId, data *models.FileUpdate) (err common.Error)
		DeleteFile(fileId *models.FileId) (err common.Error)
	}
)
