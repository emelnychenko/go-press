package validators

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"gopkg.in/go-playground/validator.v9"
)

type (
	modelValidatorImpl struct {
		validate *validator.Validate
	}
)

func NewModelValidator() (modelValidator contracts.ModelValidator) {
	return &modelValidatorImpl{validate: validator.New()}
}

func (v *modelValidatorImpl) ValidateModel(model interface{}) (err errors.Error) {
	validateErr := v.validate.Struct(model)

	if nil != validateErr {
		err = errors.NewBadRequestErrorFromBuiltin(validateErr)
	}

	return
}
