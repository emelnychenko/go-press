package dispatchers

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/events"
	eventsLib "github.com/liuggio/events"
)

type (
	eventDispatcherImpl struct {
		dispatcher eventsLib.Dispatcher
	}
)

func NewEventDispatcher() contracts.EventDispatcher {
	dispatcher := eventsLib.New()
	return &eventDispatcherImpl{dispatcher: dispatcher}
}

func (d *eventDispatcherImpl) Dispatch(event contracts.Event) {
	d.dispatcher.Raise(event.Name(), event)
}

func (d *eventDispatcherImpl) Subscribe(eventName string, eventSubscriberFunc contracts.EventSubscriberFunc) {
	d.dispatcher.On(eventName, func(data interface{}) error {
		return eventSubscriberFunc(data.(*events.Event))
	})
}
