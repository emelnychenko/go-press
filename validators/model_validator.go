package validators

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
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

func (v *modelValidatorImpl) ValidateModel(model interface{}) (err common.Error) {
	validateErr := v.validate.Struct(model)

	if nil != validateErr {
		err = common.NewBadRequestErrorFromBuiltin(validateErr)
	}

	return
}
