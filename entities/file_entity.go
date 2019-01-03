package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	FileTableName = "files"
)

type (
	FileEntity struct {
		Id      *models.FileId `gorm:"primary_key;type:char(36);column:id"`
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
		Id:      common.NewModelId(),
		Created: &created,
	}
}

func (*FileEntity) TableName() string {
	return FileTableName
}
