package models

import "github.com/emelnychenko/go-press/enums"

type (
	Subject interface {
		SubjectId() *ModelId
		SubjectType() enums.SubjectType
	}
)
