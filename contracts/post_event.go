package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostEvent interface {
		Event
		PostEntity() *entities.PostEntity
	}
)
