package contracts

import "github.com/emelnychenko/go-press/entities"

type (
	BannerEntityFactory interface {
		CreateBannerEntity() *entities.BannerEntity
	}
)
