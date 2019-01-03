package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	UserPictureEvent interface {
		Event
		UserEntity() *entities.UserEntity
		UserPictureEntity() *entities.FileEntity
	}
)
