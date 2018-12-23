package models

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"time"
)

type (
	User struct {
		Id        *UserId    `json:"id" xml:"id"`
		FirstName string     `json:"firstName" xml:"firstName"`
		LastName  string     `json:"lastName" xml:"lastName"`
		Email     string     `json:"email" xml:"email"`
		Picture   *File      `json:"picture" xml:"picture"`
		Verified  bool       `json:"verified" xml:"verified"`
		Created   *time.Time `json:"created" xml:"created"`
		Updated   *time.Time `json:"updated" xml:"updated"`
	}

	UserCreate struct {
		FirstName string `json:"firstName" xml:"firstName"`
		LastName  string `json:"lastName" xml:"lastName"`
		Email     string `json:"email" xml:"email"`
		Password  string `json:"password" xml:"password"`
	}

	UserUpdate struct {
		FirstName string `json:"firstName" xml:"firstName"`
		LastName  string `json:"lastName" xml:"lastName"`
	}

	UserChangePassword struct {
		OldPassword string `json:"oldPassword" xml:"oldPassword"`
		NewPassword string `json:"newPassword" xml:"newPassword"`
	}

	UserChangeIdentity struct {
		Email string `json:"email" xml:"email"`
	}
)

func (c *User) SubjectId() *common.ModelId {
	return c.Id
}

func (*User) SubjectType() enums.SubjectType {
	return enums.UserSubjectType
}
