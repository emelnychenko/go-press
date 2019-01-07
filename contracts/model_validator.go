package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	ModelValidator interface {
		ValidateModel(model interface{}) errors.Error
	}
)
