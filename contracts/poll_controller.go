package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	PollController interface {
		ListPolls(httpContext HttpContext) (response interface{}, err errors.Error)
		GetPoll(httpContext HttpContext) (response interface{}, err errors.Error)
		CreatePoll(httpContext HttpContext) (response interface{}, err errors.Error)
		UpdatePoll(httpContext HttpContext) (_ interface{}, err errors.Error)
		DeletePoll(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
