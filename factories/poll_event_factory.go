package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	pollEventFactoryImpl struct {
	}
)

func NewPollEventFactory() contracts.PollEventFactory {
	return new(pollEventFactoryImpl)
}

func (*pollEventFactoryImpl) CreatePollCreatedEvent(pollEntity *entities.PollEntity) contracts.PollEvent {
	return events.NewPollCreatedEvent(pollEntity)
}

func (*pollEventFactoryImpl) CreatePollUpdatedEvent(pollEntity *entities.PollEntity) contracts.PollEvent {
	return events.NewPollUpdatedEvent(pollEntity)
}

func (*pollEventFactoryImpl) CreatePollDeletedEvent(pollEntity *entities.PollEntity) contracts.PollEvent {
	return events.NewPollDeletedEvent(pollEntity)
}
