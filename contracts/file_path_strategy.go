package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	FilePathStrategy interface {
		BuildPath(fileEntity *entities.FileEntity) (filePath string, err errors.Error)
	}
)
