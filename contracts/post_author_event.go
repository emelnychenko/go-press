package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostAuthorEvent interface {
		Event
		PostEntity() *entities.PostEntity
		PostAuthorEntity() *entities.UserEntity
	}
)
