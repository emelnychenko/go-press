package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
)

type (
	PostPublisherJob interface {
		PublishPost(postEntity *entities.PostEntity) (err common.Error)
	}
)