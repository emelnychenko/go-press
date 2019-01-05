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
		Id               *models.CategoryId `gorm:"primary_key;type:char(36);column:id"`
		Name             string
		ParentCategoryId *models.CategoryId `gorm:"type:char(36)"`
		Left             int
		Right            int
		Created          *time.Time
		Updated          *time.Time
	}

	CategoryEntityTree struct {
		Roots []*CategoryEntityTreeBranch
	}

	CategoryEntityTreeBranch struct {
		CategoryEntity *CategoryEntity
		Children       []*CategoryEntityTreeBranch
	}

	CategoryEntityNestedSet struct {
		Nodes []*CategoryEntityNestedSetNode
	}

	CategoryEntityNestedSetNode struct {
		CategoryEntity *CategoryEntity
		Left           int
		Right          int
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
