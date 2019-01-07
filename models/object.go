package models

type (
	ObjectType string
	ObjectId   = ModelId

	Object interface {
		ObjectId() *ObjectId
		ObjectType() ObjectType
	}
)
