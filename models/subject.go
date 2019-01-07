package models

type (
	SubjectType string
	SubjectId   = ModelId

	Subject interface {
		SubjectId() *SubjectId
		SubjectType() SubjectType
	}
)
