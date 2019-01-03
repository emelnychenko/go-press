package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	channelModelFactoryImpl struct {
		paginationModelFactory contracts.PaginationModelFactory
	}
)

func NewChannelModelFactory(paginationModelFactory contracts.PaginationModelFactory) contracts.ChannelModelFactory {
	return &channelModelFactoryImpl{paginationModelFactory}
}

func (f *channelModelFactoryImpl) CreateChannelPaginationQuery() *models.ChannelPaginationQuery {
	paginationQuery := f.paginationModelFactory.CreatePaginationQuery()
	return &models.ChannelPaginationQuery{PaginationQuery: paginationQuery}
}

func (*channelModelFactoryImpl) CreateChannel() *models.Channel {
	return new(models.Channel)
}

func (*channelModelFactoryImpl) CreateChannelCreate() *models.ChannelCreate {
	return new(models.ChannelCreate)
}

func (*channelModelFactoryImpl) CreateChannelUpdate() *models.ChannelUpdate {
	return new(models.ChannelUpdate)
}
