package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	PostController interface {
		ListPosts(httpContext HttpContext) (response interface{}, err common.Error)
		GetPost(httpContext HttpContext) (response interface{}, err common.Error)
		CreatePost(httpContext HttpContext) (response interface{}, err common.Error)
		UpdatePost(httpContext HttpContext) (_ interface{}, err common.Error)
		DeletePost(httpContext HttpContext) (_ interface{}, err common.Error)
	}
)
