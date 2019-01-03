package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelEntityFactory(t *testing.T) {
	t.Run("NewChannelEntityFactory", func(t *testing.T) {
		_, isChannelEntityFactory := NewChannelEntityFactory().(*channelEntityFactoryImpl)

		assert.True(t, isChannelEntityFactory)
	})

	t.Run("CreateChannelEntity", func(t *testing.T) {
		channelEntityFactory := new(channelEntityFactoryImpl)
		assert.IsType(t, new(entities.ChannelEntity), channelEntityFactory.CreateChannelEntity())
	})
}
