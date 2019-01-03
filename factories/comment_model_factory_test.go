package factories

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentModelFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCommentModelFactory", func(t *testing.T) {
		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		commentModelFactory, isCommentModelFactory := NewCommentModelFactory(paginationModelFactory).(*commentModelFactoryImpl)

		assert.True(t, isCommentModelFactory)
		assert.Equal(t, paginationModelFactory, commentModelFactory.paginationModelFactory)
	})

	t.Run("CreateCommentPaginationQuery", func(t *testing.T) {
		paginationQuery := new(models.PaginationQuery)

		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		paginationModelFactory.EXPECT().CreatePaginationQuery().Return(paginationQuery)

		commentModelFactory := &commentModelFactoryImpl{paginationModelFactory: paginationModelFactory}
		commentPaginationQuery := commentModelFactory.CreateCommentPaginationQuery()

		assert.Equal(t, paginationQuery, commentPaginationQuery.PaginationQuery)
	})

	t.Run("CreateComment", func(t *testing.T) {
		commentModelFactory := new(commentModelFactoryImpl)
		assert.NotNil(t, commentModelFactory.CreateComment())
	})

	t.Run("CreateCommentCreate", func(t *testing.T) {
		commentModelFactory := new(commentModelFactoryImpl)
		assert.NotNil(t, commentModelFactory.CreateCommentCreate())
	})

	t.Run("CreateCommentUpdate", func(t *testing.T) {
		commentModelFactory := new(commentModelFactoryImpl)
		assert.NotNil(t, commentModelFactory.CreateCommentUpdate())
	})
}
