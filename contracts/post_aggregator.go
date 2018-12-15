package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostAggregator interface {
		AggregateObject(postEntity *entities.PostEntity) (post *models.Post)
		AggregateCollection(postEntities []*entities.PostEntity) (posts []*models.Post)
	}
)
