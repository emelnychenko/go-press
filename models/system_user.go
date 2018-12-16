package models

import (
	"github.com/emelnychenko/go-press/enums"
	"github.com/google/uuid"
)

type (
	SystemUser struct {
		Id *ModelId
	}
)

func NewSystemUser() *SystemUser {
	return &SystemUser{Id: &ModelId{uuid.Nil}}
}

func (c *SystemUser) SubjectId() *ModelId {
	return c.Id
}

func (*SystemUser) SubjectType() enums.SubjectType {
	return enums.SystemSubjectType
}
