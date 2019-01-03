package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	CommentEvent interface {
		Event
		CommentEntity() *entities.CommentEntity
	}
)
