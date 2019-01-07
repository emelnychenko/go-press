package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserAggregator interface {
		AggregateUser(object *entities.UserEntity) *models.User
		AggregateUsers(collection []*entities.UserEntity) []*models.User
		AggregatePaginationResult(entityPaginationResult *models.PaginationResult) *models.PaginationResult
	}
)
