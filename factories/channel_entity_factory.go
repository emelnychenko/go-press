package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	channelEntityFactoryImpl struct {
	}
)

func NewChannelEntityFactory() contracts.ChannelEntityFactory {
	return new(channelEntityFactoryImpl)
}

func (*channelEntityFactoryImpl) CreateChannelEntity() *entities.ChannelEntity {
	return entities.NewChannelEntity()
}
