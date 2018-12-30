package contracts

import "github.com/emelnychenko/go-press/common"

type (
	ModelValidator interface {
		ValidateModel(model interface{}) common.Error
	}
)