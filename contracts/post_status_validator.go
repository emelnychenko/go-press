package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostStatusValidator interface {
		ValidatePostCreate(data *models.PostCreate) errors.Error
		ValidatePostUpdate(data *models.PostUpdate) errors.Error
	}
)
