package entities

import (
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	FileTable = "files"
)

type (
	FileEntity struct {
		Id      *models.PostId `gorm:"primary_key;type:char(36);column:id"`
		Name    string
		Path    string
		Size    int64
		Type    string
		Created *time.Time
		Updated *time.Time
	}
)

func NewFileEntity() *FileEntity {
	created := time.Now().UTC()
	return &FileEntity{
		Id:      models.NewModelId(),
		Size:    0,
		Created: &created,
	}
}

func (*FileEntity) TableName() string {
	return FileTable
}
