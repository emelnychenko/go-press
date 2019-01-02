package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	PostPublisherWorker interface {
		Start() (err common.Error)
		Stop() (err common.Error)
	}
)
