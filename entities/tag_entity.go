package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	TagTableName = "tags"
)

type (
	TagEntity struct {
		Id      *models.TagId `gorm:"primary_key;type:char(36);column:id"`
		Name    string
		Created *time.Time
		Updated *time.Time
	}
)

func NewTagEntity() *TagEntity {
	created := time.Now().UTC()
	return &TagEntity{
		Id:      common.NewModelId(),
		Created: &created,
	}
}

func (*TagEntity) TableName() string {
	return TagTableName
}
