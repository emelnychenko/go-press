package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	UserEvent interface {
		Event
		UserEntity() *entities.UserEntity
	}
)
