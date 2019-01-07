package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"io"
)

type (
	FileService interface {
		ListFiles(*models.FilePaginationQuery) (*models.PaginationResult, errors.Error)
		GetFile(fileId *models.FileId) (fileEntity *entities.FileEntity, err errors.Error)
		UploadFile(fileSource io.Reader, data *models.FileUpload) (fileEntity *entities.FileEntity, err errors.Error)
		DownloadFile(fileEntity *entities.FileEntity, fileDestination io.Writer) (err errors.Error)
		UpdateFile(fileEntity *entities.FileEntity, data *models.FileUpdate) (err errors.Error)
		DeleteFile(fileEntity *entities.FileEntity) (err errors.Error)
	}
)
