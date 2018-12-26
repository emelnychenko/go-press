package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostNormalizer interface {
		NormalizePostEntity(postEntity *entities.PostEntity)
	}
)
