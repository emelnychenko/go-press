package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	PollController interface {
		ListPolls(httpContext HttpContext) (response interface{}, err common.Error)
		GetPoll(httpContext HttpContext) (response interface{}, err common.Error)
		CreatePoll(httpContext HttpContext) (response interface{}, err common.Error)
		UpdatePoll(httpContext HttpContext) (_ interface{}, err common.Error)
		DeletePoll(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
