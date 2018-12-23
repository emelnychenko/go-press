package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	UserHttpHelper interface {
		ParseUserId(httpContext HttpContext) (userId *models.UserId, err common.Error)
	}
)
