package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	TagEvent interface {
		Event
		TagEntity() *entities.TagEntity
	}
)
