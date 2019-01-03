package aggregator

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentAggregator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCommentAggregator", func(t *testing.T) {
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentAggregator, isCommentAggregator := NewCommentAggregator(commentModelFactory).(*commentAggregatorImpl)

		assert.True(t, isCommentAggregator)
		assert.Equal(t, commentModelFactory, commentAggregator.commentModelFactory)
	})

	t.Run("AggregateComment", func(t *testing.T) {
		comment := new(models.Comment)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateComment().Return(comment)

		commentAggregator := &commentAggregatorImpl{commentModelFactory: commentModelFactory}
		response := commentAggregator.AggregateComment(new(entities.CommentEntity))

		assert.Equal(t, comment, response)
	})

	t.Run("AggregateComments", func(t *testing.T) {
		comments := new(models.Comment)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateComment().Return(comments)

		commentAggregator := &commentAggregatorImpl{commentModelFactory: commentModelFactory}
		commentEntities := []*entities.CommentEntity{new(entities.CommentEntity)}
		response := commentAggregator.AggregateComments(commentEntities)

		assert.IsType(t, []*models.Comment{}, response)
		assert.Equal(t, len(commentEntities), len(response))
	})

	t.Run("AggregatePaginationResult", func(t *testing.T) {
		comment := new(models.Comment)
		commentModelFactory := mocks.NewMockCommentModelFactory(ctrl)
		commentModelFactory.EXPECT().CreateComment().Return(comment)

		commentEntities := []*entities.CommentEntity{entities.NewCommentEntity()}
		commentAggregator := &commentAggregatorImpl{commentModelFactory: commentModelFactory}

		entityPaginationResult := &models.PaginationResult{Data: commentEntities}
		paginationResult := commentAggregator.AggregatePaginationResult(entityPaginationResult)

		assert.IsType(t, []*models.Comment{}, paginationResult.Data)
		assert.Equal(t, len(commentEntities), len(paginationResult.Data.([]*models.Comment)))
	})
}
