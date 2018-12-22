package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostEntityFactory interface {
		CreatePostEntity() (postEntity *entities.PostEntity)
	}
)
