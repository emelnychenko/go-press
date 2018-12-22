package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"io"
)

type (
	StorageProvider interface {
		UploadFile(fileEntity *entities.FileEntity, fileSource io.Reader) (err common.Error)
		DownloadFile(fileEntity *entities.FileEntity, fileDestination io.Writer) (err common.Error)
		DeleteFile(fileEntity *entities.FileEntity) (err common.Error)
	}
)
