package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	FileEntityFactory interface {
		CreateFileEntity() (fileEntity *entities.FileEntity)
	}
)
