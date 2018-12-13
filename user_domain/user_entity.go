package user_domain

import (
	"../common"
)

type (
	UserEntity struct {
		Id *UserId `gorm:"primary_key;type:char(36);column:id"`
		FirstName string
		LastName string
		Email string `gorm:"type:varchar(100);unique_index"`
		Password string
		Verified bool
	}
)

func NewUserEntity() *UserEntity {
	return &UserEntity{Id: common.NewModelId(), Verified: false}
}

func (*UserEntity) TableName() string {
	return UserTable
}
