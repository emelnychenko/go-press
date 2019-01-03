package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostVideoEvent interface {
		Event
		PostEntity() *entities.PostEntity
		PostVideoEntity() *entities.FileEntity
	}
)
