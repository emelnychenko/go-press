package user

import (
	"../user_domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserAggregator(t *testing.T) {
	t.Run("AggregateObject", func(t *testing.T) {
		aggregator := NewUserAggregator()
		model := aggregator.AggregateObject(&user_domain.UserEntity{})

		assert.IsType(t, &user_domain.User{}, model)
	})

	t.Run("AggregateCollection", func(t *testing.T) {
		aggregator := NewUserAggregator()
		entities := []*user_domain.UserEntity{user_domain.NewUserEntity()}
		models := aggregator.AggregateCollection(entities)
		assert.IsType(t, []*user_domain.User{}, models)
		assert.Equal(t, len(entities), len(models))
	})
}