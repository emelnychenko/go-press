package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	ChannelController interface {
		ListChannels(httpContext HttpContext) (response interface{}, err common.Error)
		GetChannel(httpContext HttpContext) (response interface{}, err common.Error)
		CreateChannel(httpContext HttpContext) (response interface{}, err common.Error)
		UpdateChannel(httpContext HttpContext) (_ interface{}, err common.Error)
		DeleteChannel(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
