package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	CommentController interface {
		ListComments(httpContext HttpContext) (response interface{}, err common.Error)
		GetComment(httpContext HttpContext) (response interface{}, err common.Error)
		CreateComment(httpContext HttpContext) (response interface{}, err common.Error)
		UpdateComment(httpContext HttpContext) (_ interface{}, err common.Error)
		DeleteComment(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
