package factories

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollEntityFactory(t *testing.T) {
	t.Run("NewPollEntityFactory", func(t *testing.T) {
		_, isPollEntityFactory := NewPollEntityFactory().(*pollEntityFactoryImpl)

		assert.True(t, isPollEntityFactory)
	})

	t.Run("CreatePollEntity", func(t *testing.T) {
		pollEntityFactory := new(pollEntityFactoryImpl)
		assert.IsType(t, new(entities.PollEntity), pollEntityFactory.CreatePollEntity())
	})
}
