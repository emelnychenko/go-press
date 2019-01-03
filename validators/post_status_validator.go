package validators

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/models"
	"gopkg.in/go-playground/validator.v9"
)

type (
	postStatusValidatorImpl struct {
		validate *validator.Validate
	}
)

func NewPostStatusValidator() (postStatusValidator contracts.PostStatusValidator) {
	return &postStatusValidatorImpl{validate: validator.New()}
}

func (v *postStatusValidatorImpl) ValidatePostCreate(data *models.PostCreate) (err common.Error) {
	if enums.PostScheduledStatus == data.Status {
		validateErr := v.validate.Var(data.Published, "required")

		if nil != validateErr {
			err = common.NewBadRequestErrorFromBuiltin(validateErr)
		}
	}

	return
}

func (v *postStatusValidatorImpl) ValidatePostUpdate(data *models.PostUpdate) (err common.Error) {
	if enums.PostScheduledStatus == data.Status {
		validateErr := v.validate.Var(data.Published, "required")

		if nil != validateErr {
			err = common.NewBadRequestErrorFromBuiltin(validateErr)
		}
	}

	return
}
