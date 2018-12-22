package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	PostTable = "posts"
)

type (
	PostEntity struct {
		Id          *models.PostId `gorm:"primary_key;type:char(36);column:id"`
		AuthorId    *common.ModelId
		AuthorType  enums.SubjectType
		Title       string
		Description string `gorm:"type:text"`
		Content     string
		Status      enums.PostStatus
		Privacy     enums.PostPrivacy
		Views       int
		Published   *time.Time
		Created     *time.Time
		Updated     *time.Time
	}
)

func NewPostEntity() *PostEntity {
	created := time.Now().UTC()
	return &PostEntity{
		Id:      common.NewModelId(),
		Status:  enums.PostDraftStatus,
		Privacy: enums.PostPublicPrivacy,
		Created: &created,
	}
}

func (*PostEntity) TableName() string {
	return PostTable
}

func (c *PostEntity) SetAuthor(postAuthor common.Subject) {
	c.AuthorId = postAuthor.SubjectId()
	c.AuthorType = postAuthor.SubjectType()
}
