package helpers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

const (
	PollIdParameterName = "pollId"
)

type (
	pollHttpHelperImpl struct {
	}
)

func NewPollHttpHelper() contracts.PollHttpHelper {
	return new(pollHttpHelperImpl)
}

func (*pollHttpHelperImpl) ParsePollId(httpContext contracts.HttpContext) (*models.PollId, common.Error) {
	return common.ParseModelId(httpContext.Parameter(PollIdParameterName))
}

