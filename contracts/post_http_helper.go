package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostHttpHelper interface {
		ParsePostId(httpContext HttpContext) (postId *models.PostId, err common.Error)
	}
)
