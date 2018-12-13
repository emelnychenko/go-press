package common

import (
	"github.com/google/uuid"
)

type (
	ModelId struct {
		uuid.UUID
	}
)

func NewModelId() *ModelId {
	return &ModelId{uuid.New()}
}

func ParseModelId(value string) (*ModelId, error) {
	id, err := uuid.Parse(value)

	if nil != err {
		return nil, err
	}

	return &ModelId{id}, nil
}
