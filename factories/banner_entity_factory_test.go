package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerEntityFactory(t *testing.T) {
	t.Run("NewBannerEntityFactory", func(t *testing.T) {
		_, isBannerEntityFactory := NewBannerEntityFactory().(*bannerEntityFactoryImpl)

		assert.True(t, isBannerEntityFactory)
	})

	t.Run("CreateBannerEntity", func(t *testing.T) {
		bannerEntityFactory := new(bannerEntityFactoryImpl)
		assert.IsType(t, new(entities.BannerEntity), bannerEntityFactory.CreateBannerEntity())
	})
}
