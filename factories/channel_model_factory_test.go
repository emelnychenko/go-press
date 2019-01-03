package factories

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelModelFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewChannelModelFactory", func(t *testing.T) {
		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		channelModelFactory, isChannelModelFactory := NewChannelModelFactory(
			paginationModelFactory,
		).(*channelModelFactoryImpl)

		assert.True(t, isChannelModelFactory)
		assert.Equal(t, paginationModelFactory, channelModelFactory.paginationModelFactory)
	})

	t.Run("CreateChannelPaginationQuery", func(t *testing.T) {
		paginationQuery := new(models.PaginationQuery)

		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		paginationModelFactory.EXPECT().CreatePaginationQuery().Return(paginationQuery)

		channelModelFactory := &channelModelFactoryImpl{paginationModelFactory: paginationModelFactory}
		channelPaginationQuery := channelModelFactory.CreateChannelPaginationQuery()

		assert.Equal(t, paginationQuery, channelPaginationQuery.PaginationQuery)
	})

	t.Run("CreateChannel", func(t *testing.T) {
		channelModelFactory := new(channelModelFactoryImpl)
		assert.NotNil(t, channelModelFactory.CreateChannel())
	})

	t.Run("CreateChannelCreate", func(t *testing.T) {
		channelModelFactory := new(channelModelFactoryImpl)
		assert.NotNil(t, channelModelFactory.CreateChannelCreate())
	})

	t.Run("CreateChannelUpdate", func(t *testing.T) {
		channelModelFactory := new(channelModelFactoryImpl)
		assert.NotNil(t, channelModelFactory.CreateChannelUpdate())
	})
}
