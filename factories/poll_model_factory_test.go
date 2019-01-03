package factories

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollModelFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPollModelFactory", func(t *testing.T) {
		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		pollModelFactory, isPollModelFactory := NewPollModelFactory(paginationModelFactory).(*pollModelFactoryImpl)

		assert.True(t, isPollModelFactory)
		assert.Equal(t, paginationModelFactory, pollModelFactory.paginationModelFactory)
	})

	t.Run("CreatePollPaginationQuery", func(t *testing.T) {
		paginationQuery := new(models.PaginationQuery)

		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		paginationModelFactory.EXPECT().CreatePaginationQuery().Return(paginationQuery)

		pollModelFactory := &pollModelFactoryImpl{paginationModelFactory: paginationModelFactory}
		pollPaginationQuery := pollModelFactory.CreatePollPaginationQuery()

		assert.Equal(t, paginationQuery, pollPaginationQuery.PaginationQuery)
	})

	t.Run("CreatePoll", func(t *testing.T) {
		pollModelFactory := new(pollModelFactoryImpl)
		assert.NotNil(t, pollModelFactory.CreatePoll())
	})

	t.Run("CreatePollCreate", func(t *testing.T) {
		pollModelFactory := new(pollModelFactoryImpl)
		assert.NotNil(t, pollModelFactory.CreatePollCreate())
	})

	t.Run("CreatePollUpdate", func(t *testing.T) {
		pollModelFactory := new(pollModelFactoryImpl)
		assert.NotNil(t, pollModelFactory.CreatePollUpdate())
	})
}
