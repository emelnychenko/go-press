package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	FileEvent interface {
		Event
		FileEntity() *entities.FileEntity
	}
)
