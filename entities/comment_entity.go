package entities

import (
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	CommentTableName = "tags"
)

type (
	CommentEntity struct {
		Id      *models.CommentId `gorm:"primary_key;type:char(36);column:id"`
		Content string
		Created *time.Time
		Updated *time.Time
	}
)

func NewCommentEntity() *CommentEntity {
	created := time.Now().UTC()
	return &CommentEntity{
		Id:      models.NewModelId(),
		Created: &created,
	}
}

func (*CommentEntity) TableName() string {
	return CommentTableName
}
