package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type userAggregatorImpl struct {
	userModelFactory contracts.UserModelFactory
}

func NewUserAggregator(userModelFactory contracts.UserModelFactory) contracts.UserAggregator {
	return &userAggregatorImpl{userModelFactory}
}

func (a *userAggregatorImpl) AggregateUser(e *entities.UserEntity) *models.User {
	m := a.userModelFactory.CreateUser()
	m.Id = e.Id
	m.FirstName = e.FirstName
	m.LastName = e.LastName
	m.Email = e.Email
	m.Verified = e.Verified
	return m
}

func (a *userAggregatorImpl) AggregateUsers(e []*entities.UserEntity) []*models.User {
	m := make([]*models.User, len(e))

	for k, v := range e {
		m[k] = a.AggregateUser(v)
	}

	return m
}
