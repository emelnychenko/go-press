package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerEntity(t *testing.T) {
	t.Run("NewBannerEntity", func(t *testing.T) {
		bannerEntity := NewBannerEntity()

		assert.IsType(t, new(models.BannerId), bannerEntity.Id)
		assert.NotNil(t, bannerEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		bannerEntity := new(BannerEntity)

		assert.Equal(t, BannerTableName, bannerEntity.TableName())
	})
}
