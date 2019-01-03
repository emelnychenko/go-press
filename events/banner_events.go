package events

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

const (
	BannerCreatedEventName = "BannerCreatedEvent"
	BannerUpdatedEventName = "BannerUpdatedEvent"
	BannerDeletedEventName = "BannerDeletedEvent"
)

type (
	BannerEvent struct {
		*Event
		bannerEntity *entities.BannerEntity
	}
)

func (e *BannerEvent) BannerEntity() *entities.BannerEntity {
	return e.bannerEntity
}

func NewBannerCreatedEvent(bannerEntity *entities.BannerEntity) contracts.BannerEvent {
	event := &Event{name: BannerCreatedEventName}
	return &BannerEvent{bannerEntity: bannerEntity, Event: event}
}

func NewBannerUpdatedEvent(bannerEntity *entities.BannerEntity) contracts.BannerEvent {
	event := &Event{name: BannerUpdatedEventName}
	return &BannerEvent{bannerEntity: bannerEntity, Event: event}
}

func NewBannerDeletedEvent(bannerEntity *entities.BannerEntity) contracts.BannerEvent {
	event := &Event{name: BannerDeletedEventName}
	return &BannerEvent{bannerEntity: bannerEntity, Event: event}
}
