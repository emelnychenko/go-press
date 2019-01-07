package models

import (
	"github.com/emelnychenko/go-press/errors"
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

func ParseModelId(value string) (*ModelId, errors.Error) {
	id, err := uuid.Parse(value)

	if nil != err {
		return nil, errors.NewSystemErrorFromBuiltin(err)
	}

	return &ModelId{id}, nil
}
