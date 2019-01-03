package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	PollTableName = "polls"
)

type (
	PollEntity struct {
		Id      *models.PollId `gorm:"primary_key;type:char(36);column:id"`
		Title   string
		Created *time.Time
		Updated *time.Time
	}
)

func NewPollEntity() *PollEntity {
	created := time.Now().UTC()
	return &PollEntity{
		Id:      common.NewModelId(),
		Created: &created,
	}
}

func (*PollEntity) TableName() string {
	return PollTableName
}
