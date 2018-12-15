package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type userAggregatorImpl struct {
}

func NewUserAggregator() contracts.UserAggregator {
	return new(userAggregatorImpl)
}

func (*userAggregatorImpl) AggregateObject(e *entities.UserEntity) *models.User {
	m := new(models.User)
	m.Id = e.Id
	m.FirstName = e.FirstName
	m.LastName = e.LastName
	m.Email = e.Email
	m.Verified = e.Verified
	return m
}

func (c *userAggregatorImpl) AggregateCollection(e []*entities.UserEntity) []*models.User {
	m := make([]*models.User, len(e))

	for k, v := range e {
		m[k] = c.AggregateObject(v)
	}

	return m
}

