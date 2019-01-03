package dispatchers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/events_mocks"
	"github.com/golang/mock/gomock"
	eventsLib "github.com/liuggio/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventDispatcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewEventDispatcher", func(t *testing.T) {
		eventDispatcher, isEventDispatcher := NewEventDispatcher().(*eventDispatcherImpl)

		assert.True(t, isEventDispatcher)
		assert.NotNil(t, eventDispatcher)
	})

	t.Run("Dispatch", func(t *testing.T) {
		event := new(events.Event)

		dispatcher := events_mocks.NewMockDispatcher(ctrl)
		dispatcher.EXPECT().Raise(event.Name(), event)

		eventDispatcher := &eventDispatcherImpl{dispatcher: dispatcher}
		eventDispatcher.Dispatch(event)
	})

	t.Run("Subscribe", func(t *testing.T) {
		event := new(events.Event)

		var eventSubscriberFunc contracts.EventSubscriberFunc = func(_event contracts.Event) (err common.Error) {
			assert.Equal(t, _event, event)
			return
		}

		dispatcher := events_mocks.NewMockDispatcher(ctrl)
		dispatcher.EXPECT().On(event.Name(), gomock.Any())

		eventDispatcher := &eventDispatcherImpl{dispatcher: dispatcher}
		eventDispatcher.Subscribe(event.Name(), eventSubscriberFunc)
	})

	t.Run("Subscribe:Dispatch", func(t *testing.T) {
		event := new(events.Event)

		var eventSubscriberFunc contracts.EventSubscriberFunc = func(_event contracts.Event) (err common.Error) {
			assert.Equal(t, _event, event)
			return
		}

		dispatcher := eventsLib.New()
		eventDispatcher := &eventDispatcherImpl{dispatcher: dispatcher}
		eventDispatcher.Subscribe(event.Name(), eventSubscriberFunc)

		eventDispatcher.Dispatch(event)
	})
}
