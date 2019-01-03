package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	BannerEventFactory interface {
		CreateBannerCreatedEvent(bannerEntity *entities.BannerEntity) BannerEvent
		CreateBannerUpdatedEvent(bannerEntity *entities.BannerEntity) BannerEvent
		CreateBannerDeletedEvent(bannerEntity *entities.BannerEntity) BannerEvent
	}
)
