package user

import "../user_domain"

type userAggregatorImpl struct {
}

func NewUserAggregator() *userAggregatorImpl {
	return new(userAggregatorImpl)
}

func (*userAggregatorImpl) AggregateObject(e *user_domain.UserEntity) *user_domain.User {
	m := new(user_domain.User)
	m.Id = e.Id
	m.FirstName = e.FirstName
	m.LastName = e.LastName
	m.Email = e.Email
	m.Verified = e.Verified
	return m
}

func (c *userAggregatorImpl) AggregateCollection(e []*user_domain.UserEntity) []*user_domain.User {
	m := make([]*user_domain.User, len(e))

	for k, v := range e {
		m[k] = c.AggregateObject(v)
	}

	return m
}

