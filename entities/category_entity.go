package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	CategoryTableName = "categories"
)

type (
	CategoryEntity struct {
		Id      *models.CategoryId `gorm:"primary_key;type:char(36);column:id"`
		Name    string
		Created *time.Time
		Updated *time.Time
	}
)

func NewCategoryEntity() *CategoryEntity {
	created := time.Now().UTC()
	return &CategoryEntity{
		Id:      common.NewModelId(),
		Created: &created,
	}
}

func (*CategoryEntity) TableName() string {
	return CategoryTableName
}
