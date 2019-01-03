package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	PollEvent interface {
		Event
		PollEntity() *entities.PollEntity
	}
)