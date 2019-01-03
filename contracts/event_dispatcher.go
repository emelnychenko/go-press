package contracts

import (
	"github.com/emelnychenko/go-press/common"
)

type (
	EventSubscriberFunc func(event Event) common.Error

	EventDispatcher interface {
		Dispatch(event Event)
		Subscribe(eventName string, eventSubscriberFunc EventSubscriberFunc)
	}
)
