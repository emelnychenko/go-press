package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostTagEvent interface {
		Event
		PostEntity() *entities.PostEntity
		TagEntity() *entities.TagEntity
	}
)
