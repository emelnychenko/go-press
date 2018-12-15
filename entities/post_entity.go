package entities

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	PostEntity struct {
		Id *models.PostId `gorm:"primary_key;type:char(36);column:id"`
		Content string `json:"content" xml:"content"`
	}
)

func NewPostEntity() *PostEntity {
	return &PostEntity{Id: models.NewModelId()}
}

func (*PostEntity) TableName() string {
	return "posts"
}
