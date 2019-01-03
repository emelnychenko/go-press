package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	BannerTableName = "banners"
)

type (
	BannerEntity struct {
		Id      *models.BannerId `gorm:"primary_key;type:char(36);column:id"`
		Title   string
		Key     string
		Created *time.Time
		Updated *time.Time
	}
)

func NewBannerEntity() *BannerEntity {
	created := time.Now().UTC()
	return &BannerEntity{
		Id:      common.NewModelId(),
		Created: &created,
	}
}

func (*BannerEntity) TableName() string {
	return BannerTableName
}
