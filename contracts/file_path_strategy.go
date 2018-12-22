package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
)

type (
	FilePathStrategy interface {
		BuildPath(fileEntity *entities.FileEntity) (filePath string, err common.Error)
	}
)
