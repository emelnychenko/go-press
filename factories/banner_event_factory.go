package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
)

type (
	bannerEventFactoryImpl struct {
	}
)

func NewBannerEventFactory() contracts.BannerEventFactory {
	return new(bannerEventFactoryImpl)
}

func (*bannerEventFactoryImpl) CreateBannerCreatedEvent(bannerEntity *entities.BannerEntity) contracts.BannerEvent {
	return events.NewBannerCreatedEvent(bannerEntity)
}

func (*bannerEventFactoryImpl) CreateBannerUpdatedEvent(bannerEntity *entities.BannerEntity) contracts.BannerEvent {
	return events.NewBannerUpdatedEvent(bannerEntity)
}

func (*bannerEventFactoryImpl) CreateBannerDeletedEvent(bannerEntity *entities.BannerEntity) contracts.BannerEvent {
	return events.NewBannerDeletedEvent(bannerEntity)
}
