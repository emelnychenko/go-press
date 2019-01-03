package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	PollHttpHelper interface {
		ParsePollId(httpContext HttpContext) (*models.PollId, common.Error)
	}
)
