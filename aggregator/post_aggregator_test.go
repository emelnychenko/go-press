package aggregator

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostAggregator(t *testing.T) {
	t.Run("AggregateObject", func(t *testing.T) {
		postAggregator := NewPostAggregator()
		post := postAggregator.AggregateObject(new(entities.PostEntity))

		assert.IsType(t, new(models.Post), post)
	})

	t.Run("AggregateCollection", func(t *testing.T) {
		postAggregator := NewPostAggregator()
		postEntities := []*entities.PostEntity{entities.NewPostEntity()}
		posts := postAggregator.AggregateCollection(postEntities)
		assert.IsType(t, []*models.Post{}, posts)
		assert.Equal(t, len(postEntities), len(posts))
	})
}