package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	PollCreatedEventName = "PollCreatedEvent"
	PollUpdatedEventName = "PollUpdatedEvent"
	PollDeletedEventName = "PollDeletedEvent"
)

type (
	PollEvent struct {
		*Event
		pollEntity *entities.PollEntity
	}
)

func (e *PollEvent) PollEntity() *entities.PollEntity {
	return e.pollEntity
}

func NewPollCreatedEvent(pollEntity *entities.PollEntity) contracts.PollEvent {
	event := &Event{name: PollCreatedEventName}
	return &PollEvent{pollEntity: pollEntity, Event: event}
}

func NewPollUpdatedEvent(pollEntity *entities.PollEntity) contracts.PollEvent {
	event := &Event{name: PollUpdatedEventName}
	return &PollEvent{pollEntity: pollEntity, Event: event}
}

func NewPollDeletedEvent(pollEntity *entities.PollEntity) contracts.PollEvent {
	event := &Event{name: PollDeletedEventName}
	return &PollEvent{pollEntity: pollEntity, Event: event}
}
