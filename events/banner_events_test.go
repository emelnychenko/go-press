package events

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerEvents(t *testing.T) {
	t.Run("NewBannerCreatedEvent", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		bannerEvent, isBannerEvent := NewBannerCreatedEvent(bannerEntity).(*BannerEvent)

		assert.True(t, isBannerEvent)
		assert.Equal(t, bannerEntity, bannerEvent.bannerEntity)
		assert.Equal(t, BannerCreatedEventName, bannerEvent.name)
	})

	t.Run("NewBannerUpdatedEvent", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		bannerEvent, isBannerEvent := NewBannerUpdatedEvent(bannerEntity).(*BannerEvent)

		assert.True(t, isBannerEvent)
		assert.Equal(t, bannerEntity, bannerEvent.bannerEntity)
		assert.Equal(t, BannerUpdatedEventName, bannerEvent.name)
	})

	t.Run("NewBannerDeletedEvent", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		bannerEvent, isBannerEvent := NewBannerDeletedEvent(bannerEntity).(*BannerEvent)

		assert.True(t, isBannerEvent)
		assert.Equal(t, bannerEntity, bannerEvent.bannerEntity)
		assert.Equal(t, BannerDeletedEventName, bannerEvent.name)
	})

	t.Run("BannerEntity", func(t *testing.T) {
		bannerEntity := new(entities.BannerEntity)
		bannerEvent := &BannerEvent{bannerEntity: bannerEntity}

		assert.Equal(t, bannerEntity, bannerEvent.BannerEntity())
	})
}
