package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"io"
)

type (
	StorageProvider interface {
		UploadFile(fileEntity *entities.FileEntity, fileSource io.Reader) (err errors.Error)
		DownloadFile(fileEntity *entities.FileEntity, fileDestination io.Writer) (err errors.Error)
		DeleteFile(fileEntity *entities.FileEntity) (err errors.Error)
	}
)
