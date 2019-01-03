package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	ChannelEventFactory interface {
		CreateChannelCreatedEvent(channelEntity *entities.ChannelEntity) ChannelEvent
		CreateChannelUpdatedEvent(channelEntity *entities.ChannelEntity) ChannelEvent
		CreateChannelDeletedEvent(channelEntity *entities.ChannelEntity) ChannelEvent
	}
)
