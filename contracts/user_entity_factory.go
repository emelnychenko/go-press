package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	UserEntityFactory interface {
		CreateUserEntity() (userEntity *entities.UserEntity)
	}
)
