package contracts

import (
	"github.com/emelnychenko/go-press/entities"
)

type (
	BannerEvent interface {
		Event
		BannerEntity() *entities.BannerEntity
	}
)
