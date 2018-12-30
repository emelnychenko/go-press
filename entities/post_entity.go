package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	PostTableName = "posts"
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
		PictureId   *models.FileId `gorm:"type:char(36)"`
		VideoId     *models.FileId `gorm:"type:char(36)"`
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
	return PostTableName
}

func (e *PostEntity) SetPicture(postPicture *FileEntity) {
	e.PictureId = postPicture.Id
}

func (e *PostEntity) RemovePicture() {
	e.PictureId = nil
}

func (e *PostEntity) SetVideo(postVideo *FileEntity) {
	e.VideoId = postVideo.Id
}

func (e *PostEntity) RemoveVideo() {
	e.VideoId = nil
}

func (e *PostEntity) SetAuthor(postAuthor common.Subject) {
	e.AuthorId = postAuthor.SubjectId()
	e.AuthorType = postAuthor.SubjectType()
}
