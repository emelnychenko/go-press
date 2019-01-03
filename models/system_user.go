package models

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/enums"
	"github.com/google/uuid"
)

type (
	SystemUser struct {
		Id *common.ModelId `json:"id" xml:"id"`
	}
)

func NewSystemUser() *SystemUser {
	return &SystemUser{Id: &common.ModelId{uuid.Nil}}
}

func (c *SystemUser) SubjectId() *common.ModelId {
	return c.Id
}

func (*SystemUser) SubjectType() enums.SubjectType {
	return enums.SystemSubjectType
}
