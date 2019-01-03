package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	bannerEntityFactoryImpl struct {
	}
)

func NewBannerEntityFactory() contracts.BannerEntityFactory {
	return new(bannerEntityFactoryImpl)
}

func (*bannerEntityFactoryImpl) CreateBannerEntity() *entities.BannerEntity {
	return entities.NewBannerEntity()
}
