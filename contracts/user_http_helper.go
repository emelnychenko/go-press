package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserHttpHelper interface {
		ParseUserId(httpContext HttpContext) (userId *models.UserId, err errors.Error)
	}
)
