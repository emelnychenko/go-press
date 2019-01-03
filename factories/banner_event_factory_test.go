package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerEventFactory(t *testing.T) {
	t.Run("NewBannerEventFactory", func(t *testing.T) {
		_, isBannerEventFactory := NewBannerEventFactory().(*bannerEventFactoryImpl)

		assert.True(t, isBannerEventFactory)
	})

	t.Run("CreateBannerCreatedEvent", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		bannerEventFactory := new(bannerEventFactoryImpl)
		bannerEvent := bannerEventFactory.CreateBannerCreatedEvent(bannerEntity)

		assert.Equal(t, events.BannerCreatedEventName, bannerEvent.Name())
		assert.Equal(t, bannerEntity, bannerEvent.BannerEntity())
	})

	t.Run("CreateBannerUpdatedEvent", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		bannerEventFactory := new(bannerEventFactoryImpl)
		bannerEvent := bannerEventFactory.CreateBannerUpdatedEvent(bannerEntity)

		assert.Equal(t, events.BannerUpdatedEventName, bannerEvent.Name())
		assert.Equal(t, bannerEntity, bannerEvent.BannerEntity())
	})

	t.Run("CreateBannerDeletedEvent", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		bannerEventFactory := new(bannerEventFactoryImpl)
		bannerEvent := bannerEventFactory.CreateBannerDeletedEvent(bannerEntity)

		assert.Equal(t, events.BannerDeletedEventName, bannerEvent.Name())
		assert.Equal(t, bannerEntity, bannerEvent.BannerEntity())
	})
}
