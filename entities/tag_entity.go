package entities

import (
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	TagTableName     = "tags"
	TagXrefTableName = "tag_xref"
)

type (
	TagEntity struct {
		Id      *models.TagId `gorm:"primary_key;type:char(36);column:id"`
		Name    string
		Created *time.Time
		Updated *time.Time
	}

	TagXrefEntity struct {
		TagId      *models.TagId     `gorm:"primary_key;type:char(36)"`
		ObjectType models.ObjectType `gorm:"primary_key"`
		ObjectId   *models.ObjectId  `gorm:"primary_key;type:char(36)"`
		Created    *time.Time
	}
)

func NewTagEntity() *TagEntity {
	created := time.Now().UTC()
	return &TagEntity{
		Id:      models.NewModelId(),
		Created: &created,
	}
}

//TableName
func (*TagEntity) TableName() string {
	return TagTableName
}

//NewTagXrefEntity
func NewTagXrefEntity(tagEntity *TagEntity, tagObject models.Object) *TagXrefEntity {
	created := time.Now().UTC()
	return &TagXrefEntity{
		TagId: tagEntity.Id,
		ObjectType: tagObject.ObjectType(),
		ObjectId: tagObject.ObjectId(),
		Created: &created,
	}
}

//TableName
func (*TagXrefEntity) TableName() string {
	return TagXrefTableName
}

//SetTag
func (e *TagXrefEntity) SetTag(tagEntity *TagEntity) {
	e.TagId = tagEntity.Id
}

//SetObject
func (e *TagXrefEntity) SetObject(tagObject models.Object) {
	e.ObjectId = tagObject.ObjectId()
	e.ObjectType = tagObject.ObjectType()
}