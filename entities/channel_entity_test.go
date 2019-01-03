package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelEntity(t *testing.T) {
	t.Run("NewChannelEntity", func(t *testing.T) {
		channelEntity := NewChannelEntity()

		assert.IsType(t, new(models.ChannelId), channelEntity.Id)
		assert.NotNil(t, channelEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		channelEntity := new(ChannelEntity)

		assert.Equal(t, ChannelTableName, channelEntity.TableName())
	})
}
