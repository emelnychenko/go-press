package aggregator

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserAggregator(t *testing.T) {
	t.Run("AggregateObject", func(t *testing.T) {
		userAggregator := NewUserAggregator()
		userModel := userAggregator.AggregateObject(&entities.UserEntity{})

		assert.IsType(t, &models.User{}, userModel)
	})

	t.Run("AggregateCollection", func(t *testing.T) {
		userAggregator := NewUserAggregator()
		userEntities := []*entities.UserEntity{entities.NewUserEntity()}
		userModels := userAggregator.AggregateCollection(userEntities)
		assert.IsType(t, []*models.User{}, userModels)
		assert.Equal(t, len(userEntities), len(userModels))
	})
}