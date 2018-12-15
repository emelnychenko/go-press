package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserAggregator interface {
		AggregateObject(object *entities.UserEntity) *models.User
		AggregateCollection(collection []*entities.UserEntity) []*models.User
	}
)
