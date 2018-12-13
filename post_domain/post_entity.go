package post_domain

import (
	"../common"
)

type (
	PostEntity struct {
		Id *PostId `gorm:"primary_key;type:char(36);column:id"`
		Content string `json:"content" xml:"content"`
	}
)

func NewPostEntity() *PostEntity {
	return &PostEntity{Id: common.NewModelId()}
}

func (*PostEntity) TableName() string {
	return "posts"
}
