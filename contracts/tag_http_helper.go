package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	TagHttpHelper interface {
		ParseTagId(httpContext HttpContext) (*models.TagId, common.Error)
	}
)
