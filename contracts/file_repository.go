package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	FileRepository interface {
		ListFiles(
			filePaginationQuery *models.FilePaginationQuery,
		) (paginationResult *models.PaginationResult, err common.Error)
		GetFile(fileId *models.FileId) (fileEntity *entities.FileEntity, err common.Error)
		SaveFile(fileEntity *entities.FileEntity) (err common.Error)
		RemoveFile(fileEntity *entities.FileEntity) (err common.Error)
	}
)
