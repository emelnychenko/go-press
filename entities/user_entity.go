package entities

import (
	"github.com/emelnychenko/go-press/models"
)

const (
	UserTable = "users"
)

type (
	UserEntity struct {
		Id *models.UserId `gorm:"primary_key;type:char(36);column:id"`
		FirstName string
		LastName string
		Email string `gorm:"type:varchar(100);unique_index"`
		Password string
		Verified bool
	}
)

func NewUserEntity() *UserEntity {
	return &UserEntity{Id: models.NewModelId(), Verified: false}
}

func (*UserEntity) TableName() string {
	return UserTable
}
