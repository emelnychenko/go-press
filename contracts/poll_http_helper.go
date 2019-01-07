package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PollHttpHelper interface {
		ParsePollId(httpContext HttpContext) (*models.PollId, errors.Error)
	}
)
