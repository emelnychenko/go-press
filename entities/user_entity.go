package entities

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/models"
	"time"
)

const (
	UserTable = "users"
)

type (
	UserEntity struct {
		Id        *models.UserId `gorm:"primary_key;type:char(36);column:id"`
		FirstName string
		LastName  string
		Email     string `gorm:"type:varchar(100);unique_index"`
		Password  string
		Verified  bool
		Created   *time.Time
		Updated   *time.Time
	}
)

func NewUserEntity() *UserEntity {
	created := time.Now().UTC()
	return &UserEntity{Id: common.NewModelId(), Created: &created}
}

func (*UserEntity) TableName() string {
	return UserTable
}

func (c *UserEntity) SubjectId() *common.ModelId {
	return c.Id
}

func (*UserEntity) SubjectType() enums.SubjectType {
	return enums.UserSubjectType
}
