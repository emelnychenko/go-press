package factories

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryModelFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCategoryModelFactory", func(t *testing.T) {
		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		categoryModelFactory, isCategoryModelFactory := NewCategoryModelFactory(paginationModelFactory).(*categoryModelFactoryImpl)

		assert.True(t, isCategoryModelFactory)
		assert.Equal(t, paginationModelFactory, categoryModelFactory.paginationModelFactory)
	})

	t.Run("CreateCategoryPaginationQuery", func(t *testing.T) {
		paginationQuery := new(models.PaginationQuery)

		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		paginationModelFactory.EXPECT().CreatePaginationQuery().Return(paginationQuery)

		categoryModelFactory := &categoryModelFactoryImpl{paginationModelFactory: paginationModelFactory}
		categoryPaginationQuery := categoryModelFactory.CreateCategoryPaginationQuery()

		assert.Equal(t, paginationQuery, categoryPaginationQuery.PaginationQuery)
	})

	t.Run("CreateCategory", func(t *testing.T) {
		categoryModelFactory := new(categoryModelFactoryImpl)
		assert.NotNil(t, categoryModelFactory.CreateCategory())
	})

	t.Run("CreateCategoryCreate", func(t *testing.T) {
		categoryModelFactory := new(categoryModelFactoryImpl)
		assert.NotNil(t, categoryModelFactory.CreateCategoryCreate())
	})

	t.Run("CreateCategoryUpdate", func(t *testing.T) {
		categoryModelFactory := new(categoryModelFactoryImpl)
		assert.NotNil(t, categoryModelFactory.CreateCategoryUpdate())
	})
}
