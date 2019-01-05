package aggregators

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagAggregator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewTagAggregator", func(t *testing.T) {
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagAggregator, isTagAggregator := NewTagAggregator(tagModelFactory).(*tagAggregatorImpl)

		assert.True(t, isTagAggregator)
		assert.Equal(t, tagModelFactory, tagAggregator.tagModelFactory)
	})

	t.Run("AggregateTag", func(t *testing.T) {
		tag := new(models.Tag)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTag().Return(tag)

		tagAggregator := &tagAggregatorImpl{tagModelFactory: tagModelFactory}
		response := tagAggregator.AggregateTag(new(entities.TagEntity))

		assert.Equal(t, tag, response)
	})

	t.Run("AggregateTags", func(t *testing.T) {
		tags := new(models.Tag)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTag().Return(tags)

		tagAggregator := &tagAggregatorImpl{tagModelFactory: tagModelFactory}
		tagEntities := []*entities.TagEntity{new(entities.TagEntity)}
		response := tagAggregator.AggregateTags(tagEntities)

		assert.IsType(t, []*models.Tag{}, response)
		assert.Equal(t, len(tagEntities), len(response))
	})

	t.Run("AggregatePaginationResult", func(t *testing.T) {
		tag := new(models.Tag)
		tagModelFactory := mocks.NewMockTagModelFactory(ctrl)
		tagModelFactory.EXPECT().CreateTag().Return(tag)

		tagEntities := []*entities.TagEntity{entities.NewTagEntity()}
		tagAggregator := &tagAggregatorImpl{tagModelFactory: tagModelFactory}

		entityPaginationResult := &models.PaginationResult{Data: tagEntities}
		paginationResult := tagAggregator.AggregatePaginationResult(entityPaginationResult)

		assert.IsType(t, []*models.Tag{}, paginationResult.Data)
		assert.Equal(t, len(tagEntities), len(paginationResult.Data.([]*models.Tag)))
	})
}
