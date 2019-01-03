package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	ChannelCreatedEventName = "ChannelCreatedEvent"
	ChannelUpdatedEventName = "ChannelUpdatedEvent"
	ChannelDeletedEventName = "ChannelDeletedEvent"
)

type (
	ChannelEvent struct {
		*Event
		channelEntity *entities.ChannelEntity
	}
)

func (e *ChannelEvent) ChannelEntity() *entities.ChannelEntity {
	return e.channelEntity
}

func NewChannelCreatedEvent(channelEntity *entities.ChannelEntity) contracts.ChannelEvent {
	event := &Event{name: ChannelCreatedEventName}
	return &ChannelEvent{channelEntity: channelEntity, Event: event}
}

func NewChannelUpdatedEvent(channelEntity *entities.ChannelEntity) contracts.ChannelEvent {
	event := &Event{name: ChannelUpdatedEventName}
	return &ChannelEvent{channelEntity: channelEntity, Event: event}
}

func NewChannelDeletedEvent(channelEntity *entities.ChannelEntity) contracts.ChannelEvent {
	event := &Event{name: ChannelDeletedEventName}
	return &ChannelEvent{channelEntity: channelEntity, Event: event}
}
