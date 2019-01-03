package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	channelEventFactoryImpl struct {
	}
)

func NewChannelEventFactory() contracts.ChannelEventFactory {
	return new(channelEventFactoryImpl)
}

func (*channelEventFactoryImpl) CreateChannelCreatedEvent(channelEntity *entities.ChannelEntity) contracts.ChannelEvent {
	return events.NewChannelCreatedEvent(channelEntity)
}

func (*channelEventFactoryImpl) CreateChannelUpdatedEvent(channelEntity *entities.ChannelEntity) contracts.ChannelEvent {
	return events.NewChannelUpdatedEvent(channelEntity)
}

func (*channelEventFactoryImpl) CreateChannelDeletedEvent(channelEntity *entities.ChannelEntity) contracts.ChannelEvent {
	return events.NewChannelDeletedEvent(channelEntity)
}
