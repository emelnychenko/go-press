package entities

import (
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	UserTableName         = "users"
	UserEntitySubjectType models.SubjectType = "user"
)

type (
	UserEntity struct {
		Id        *models.UserId `gorm:"primary_key;type:char(36);column:id"`
		FirstName string
		LastName  string
		Email     string `gorm:"type:varchar(100);unique_index"`
		Password  string
		PictureId *models.FileId `gorm:"type:char(36)"`
		Verified  bool
		Created   *time.Time
		Updated   *time.Time
	}
)

func NewUserEntity() *UserEntity {
	created := time.Now().UTC()
	return &UserEntity{Id: models.NewModelId(), Created: &created}
}

func (*UserEntity) TableName() string {
	return UserTableName
}

func (c *UserEntity) SubjectId() *models.SubjectId {
	return c.Id
}

func (*UserEntity) SubjectType() models.SubjectType {
	return UserEntitySubjectType
}

func (e *UserEntity) SetPicture(postPicture *FileEntity) {
	e.PictureId = postPicture.Id
}

func (e *UserEntity) RemovePicture() {
	e.PictureId = nil
}
