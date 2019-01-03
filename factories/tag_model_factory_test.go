package factories

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagModelFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewTagModelFactory", func(t *testing.T) {
		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		tagModelFactory, isTagModelFactory := NewTagModelFactory(paginationModelFactory).(*tagModelFactoryImpl)

		assert.True(t, isTagModelFactory)
		assert.Equal(t, paginationModelFactory, tagModelFactory.paginationModelFactory)
	})

	t.Run("CreateTagPaginationQuery", func(t *testing.T) {
		paginationQuery := new(models.PaginationQuery)

		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		paginationModelFactory.EXPECT().CreatePaginationQuery().Return(paginationQuery)

		tagModelFactory := &tagModelFactoryImpl{paginationModelFactory: paginationModelFactory}
		tagPaginationQuery := tagModelFactory.CreateTagPaginationQuery()

		assert.Equal(t, paginationQuery, tagPaginationQuery.PaginationQuery)
	})

	t.Run("CreateTag", func(t *testing.T) {
		tagModelFactory := new(tagModelFactoryImpl)
		assert.NotNil(t, tagModelFactory.CreateTag())
	})

	t.Run("CreateTagCreate", func(t *testing.T) {
		tagModelFactory := new(tagModelFactoryImpl)
		assert.NotNil(t, tagModelFactory.CreateTagCreate())
	})

	t.Run("CreateTagUpdate", func(t *testing.T) {
		tagModelFactory := new(tagModelFactoryImpl)
		assert.NotNil(t, tagModelFactory.CreateTagUpdate())
	})
}
