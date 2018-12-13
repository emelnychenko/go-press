package user_contract

import (
	"../user_domain"
)

type (
	UserAggregator interface {
		AggregateObject(object *user_domain.UserEntity) *user_domain.User
		AggregateCollection(collection []*user_domain.UserEntity) []*user_domain.User
	}
)
