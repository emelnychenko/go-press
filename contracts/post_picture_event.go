package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostPictureEvent interface {
		Event
		PostEntity() *entities.PostEntity
		PostPictureEntity() *entities.FileEntity
	}
)