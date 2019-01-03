package entities

import (
	"github.com/emelnychenko/go-press/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollEntity(t *testing.T) {
	t.Run("NewPollEntity", func(t *testing.T) {
		pollEntity := NewPollEntity()

		assert.IsType(t, new(models.PollId), pollEntity.Id)
		assert.NotNil(t, pollEntity.Created)
	})

	t.Run("TableName", func(t *testing.T) {
		pollEntity := new(PollEntity)

		assert.Equal(t, PollTableName, pollEntity.TableName())
	})
}
