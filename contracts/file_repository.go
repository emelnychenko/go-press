package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	FileRepository interface {
		ListFiles(filePaginationQuery *models.FilePaginationQuery) (*models.PaginationResult, errors.Error)
		GetFile(fileId *models.FileId) (fileEntity *entities.FileEntity, err errors.Error)
		SaveFile(fileEntity *entities.FileEntity) (err errors.Error)
		RemoveFile(fileEntity *entities.FileEntity) (err errors.Error)
	}
)
