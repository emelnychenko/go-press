package contracts

import "github.com/emelnychenko/go-press/errors"

type (
	ChannelController interface {
		ListChannels(httpContext HttpContext) (response interface{}, err errors.Error)
		GetChannel(httpContext HttpContext) (response interface{}, err errors.Error)
		CreateChannel(httpContext HttpContext) (response interface{}, err errors.Error)
		UpdateChannel(httpContext HttpContext) (_ interface{}, err errors.Error)
		DeleteChannel(httpContext HttpContext) (_ interface{}, err errors.Error)
	}
)
