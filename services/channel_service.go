package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	channelServiceImpl struct {
		channelEntityFactory contracts.ChannelEntityFactory
		channelRepository    contracts.ChannelRepository
	}
)

func NewChannelService(
	channelEntityFactory contracts.ChannelEntityFactory,
	channelRepository contracts.ChannelRepository,
) (channelService contracts.ChannelService) {
	return &channelServiceImpl{
		channelEntityFactory,
		channelRepository,
	}
}

func (s *channelServiceImpl) ListChannels(
	channelPaginationQuery *models.ChannelPaginationQuery,
) (*models.PaginationResult, common.Error) {
	return s.channelRepository.ListChannels(channelPaginationQuery)
}

func (s *channelServiceImpl) GetChannel(channelId *models.ChannelId) (*entities.ChannelEntity, common.Error) {
	return s.channelRepository.GetChannel(channelId)
}

func (s *channelServiceImpl) CreateChannel(data *models.ChannelCreate) (
	channelEntity *entities.ChannelEntity, err common.Error,
) {
	channelEntity = s.channelEntityFactory.CreateChannelEntity()
	channelEntity.Name = data.Name

	err = s.channelRepository.SaveChannel(channelEntity)
	return
}

func (s *channelServiceImpl) UpdateChannel(channelEntity *entities.ChannelEntity, data *models.ChannelUpdate) common.Error {
	channelEntity.Name = data.Name

	updated := time.Now().UTC()
	channelEntity.Updated = &updated

	return s.channelRepository.SaveChannel(channelEntity)
}

func (s *channelServiceImpl) DeleteChannel(channelEntity *entities.ChannelEntity) common.Error {
	return s.channelRepository.RemoveChannel(channelEntity)
}
