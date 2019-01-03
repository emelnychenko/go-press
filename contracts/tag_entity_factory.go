package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	TagEntityFactory interface {
		CreateTagEntity() *entities.TagEntity
	}
)
