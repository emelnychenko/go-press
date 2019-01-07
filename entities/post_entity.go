package entities

import (
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	PostTableName                          = "posts"
	PostEntityObjectType models.ObjectType = "post"
)

type (
	PostEntity struct {
		Id          *models.PostId `gorm:"primary_key;type:char(36);column:id"`
		AuthorId    *models.SubjectId
		AuthorType  models.SubjectType
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
		Id:      models.NewModelId(),
		Status:  enums.PostDraftStatus,
		Privacy: enums.PostPublicPrivacy,
		Created: &created,
	}
}

func (*PostEntity) TableName() string {
	return PostTableName
}

//ObjectId
func (e *PostEntity) ObjectId() *models.ObjectId {
	return e.Id
}

//ObjectType
func (*PostEntity) ObjectType() models.ObjectType {
	return PostEntityObjectType
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

func (e *PostEntity) SetAuthor(postAuthor models.Subject) {
	e.AuthorId = postAuthor.SubjectId()
	e.AuthorType = postAuthor.SubjectType()
}
