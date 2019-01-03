package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	ChannelTableName = "channels"
)

type (
	ChannelEntity struct {
		Id      *models.ChannelId `gorm:"primary_key;type:char(36);column:id"`
		Name    string
		Created *time.Time
		Updated *time.Time
	}
)

func NewChannelEntity() *ChannelEntity {
	created := time.Now().UTC()
	return &ChannelEntity{
		Id:      common.NewModelId(),
		Created: &created,
	}
}

func (*ChannelEntity) TableName() string {
	return ChannelTableName
}
