package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	TagHttpHelper interface {
		ParseTagId(httpContext HttpContext) (*models.TagId, errors.Error)
	}
)
