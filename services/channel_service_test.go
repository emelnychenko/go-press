package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewChannelService", func(t *testing.T) {
		channelEntityFactory := mocks.NewMockChannelEntityFactory(ctrl)
		channelRepository := mocks.NewMockChannelRepository(ctrl)

		channelService, isChannelService := NewChannelService(
			channelEntityFactory, channelRepository,
		).(*channelServiceImpl)

		assert.True(t, isChannelService)
		assert.Equal(t, channelEntityFactory, channelService.channelEntityFactory)
		assert.Equal(t, channelRepository, channelService.channelRepository)
	})

	t.Run("ListChannels", func(t *testing.T) {
		channelPaginationQuery := new(models.ChannelPaginationQuery)

		var channelEntities *models.PaginationResult
		channelRepository := mocks.NewMockChannelRepository(ctrl)
		channelRepository.EXPECT().ListChannels(channelPaginationQuery).Return(channelEntities, nil)

		channelService := &channelServiceImpl{channelRepository: channelRepository}
		response, err := channelService.ListChannels(channelPaginationQuery)

		assert.Equal(t, channelEntities, response)
		assert.Nil(t, err)
	})

	t.Run("CreateChannel", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channelEntityFactory := mocks.NewMockChannelEntityFactory(ctrl)
		channelEntityFactory.EXPECT().CreateChannelEntity().Return(channelEntity)

		channelRepository := mocks.NewMockChannelRepository(ctrl)
		channelRepository.EXPECT().SaveChannel(channelEntity).Return(nil)

		data := &models.ChannelCreate{
			Name: "0",
		}
		channelService := &channelServiceImpl{
			channelEntityFactory: channelEntityFactory,
			channelRepository:    channelRepository,
		}
		response, err := channelService.CreateChannel(data)

		assert.IsType(t, channelEntity, response)
		assert.Nil(t, err)
		assert.Equal(t, data.Name, channelEntity.Name)
	})

	t.Run("GetChannel", func(t *testing.T) {
		channelId := new(models.ChannelId)
		channelEntity := new(entities.ChannelEntity)
		channelRepository := mocks.NewMockChannelRepository(ctrl)
		channelRepository.EXPECT().GetChannel(channelId).Return(channelEntity, nil)

		channelService := &channelServiceImpl{channelRepository: channelRepository}
		response, err := channelService.GetChannel(channelId)

		assert.Equal(t, channelEntity, response)
		assert.Nil(t, err)
	})

	t.Run("UpdateChannel", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channelRepository := mocks.NewMockChannelRepository(ctrl)
		channelRepository.EXPECT().SaveChannel(channelEntity).Return(nil)

		data := &models.ChannelUpdate{
			Name: "0",
		}
		channelService := &channelServiceImpl{channelRepository: channelRepository}
		assert.Nil(t, channelService.UpdateChannel(channelEntity, data))

		assert.Equal(t, data.Name, channelEntity.Name)
		assert.NotNil(t, channelEntity.Updated)
	})

	t.Run("DeleteChannel", func(t *testing.T) {
		channelEntity := new(entities.ChannelEntity)
		channelRepository := mocks.NewMockChannelRepository(ctrl)
		channelRepository.EXPECT().RemoveChannel(channelEntity).Return(nil)

		channelService := &channelServiceImpl{channelRepository: channelRepository}
		assert.Nil(t, channelService.DeleteChannel(channelEntity))
	})
}
