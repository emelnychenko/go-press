package factories

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostModelFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPostModelFactory", func(t *testing.T) {
		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		postModelFactory, isPostModelFactory := NewPostModelFactory(paginationModelFactory).(*postModelFactoryImpl)

		assert.True(t, isPostModelFactory)
		assert.Equal(t, paginationModelFactory, postModelFactory.paginationModelFactory)
	})

	t.Run("CreatePostPaginationQuery", func(t *testing.T) {
		paginationQuery := new(models.PaginationQuery)

		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		paginationModelFactory.EXPECT().CreatePaginationQuery().Return(paginationQuery)

		postModelFactory := &postModelFactoryImpl{paginationModelFactory: paginationModelFactory}
		postPaginationQuery := postModelFactory.CreatePostPaginationQuery()

		assert.Equal(t, paginationQuery, postPaginationQuery.PaginationQuery)
	})

	t.Run("CreatePost", func(t *testing.T) {
		postModelFactory := new(postModelFactoryImpl)
		assert.NotNil(t, postModelFactory.CreatePost())
	})

	t.Run("CreatePostCreate", func(t *testing.T) {
		postModelFactory := new(postModelFactoryImpl)
		assert.NotNil(t, postModelFactory.CreatePostCreate())
	})

	t.Run("CreatePostUpdate", func(t *testing.T) {
		postModelFactory := new(postModelFactoryImpl)
		assert.NotNil(t, postModelFactory.CreatePostUpdate())
	})
}
