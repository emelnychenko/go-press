package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"io"
)

type (
	FileService interface {
		ListFiles() (fileEntities []*entities.FileEntity, err common.Error)
		GetFile(fileId *models.FileId) (fileEntity *entities.FileEntity, err common.Error)
		UploadFile(fileSource io.Reader, data *models.FileUpload) (fileEntity *entities.FileEntity, err common.Error)
		DownloadFile(fileEntity *entities.FileEntity, fileDestination io.Writer) (err common.Error)
		UpdateFile(fileEntity *entities.FileEntity, data *models.FileUpdate) (err common.Error)
		DeleteFile(fileEntity *entities.FileEntity) (err common.Error)
	}
)
