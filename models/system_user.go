package models

import (
	"github.com/google/uuid"
)

const (
	SystemSubjectType SubjectType = "system"
)

type (
	SystemUser struct {
		Id *SubjectId `json:"id" xml:"id"`
	}
)

func NewSystemUser() *SystemUser {
	return &SystemUser{Id: &ModelId{uuid.Nil}}
}

func (c *SystemUser) SubjectId() *SubjectId {
	return c.Id
}

func (*SystemUser) SubjectType() SubjectType {
	return SystemSubjectType
}
