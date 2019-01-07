package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	TagEntityFactory interface {
		CreateTagEntity() *entities.TagEntity
		CreateTagXrefEntity(*entities.TagEntity, models.Object) *entities.TagXrefEntity
	}
)
