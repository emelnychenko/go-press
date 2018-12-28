package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostStatusValidator interface {
		ValidatePostCreate(data *models.PostCreate) common.Error
		ValidatePostUpdate(data *models.PostUpdate) common.Error
	}
)
