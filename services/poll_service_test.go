package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewPollService", func(t *testing.T) {
		pollEntityFactory := mocks.NewMockPollEntityFactory(ctrl)
		pollRepository := mocks.NewMockPollRepository(ctrl)

		pollService, isPollService := NewPollService(pollEntityFactory, pollRepository).(*pollServiceImpl)

		assert.True(t, isPollService)
		assert.Equal(t, pollEntityFactory, pollService.pollEntityFactory)
		assert.Equal(t, pollRepository, pollService.pollRepository)
	})

	t.Run("ListPolls", func(t *testing.T) {
		pollPaginationQuery := new(models.PollPaginationQuery)

		var pollEntities *models.PaginationResult
		pollRepository := mocks.NewMockPollRepository(ctrl)
		pollRepository.EXPECT().ListPolls(pollPaginationQuery).Return(pollEntities, nil)

		pollService := &pollServiceImpl{pollRepository: pollRepository}
		response, err := pollService.ListPolls(pollPaginationQuery)

		assert.Equal(t, pollEntities, response)
		assert.Nil(t, err)
	})

	t.Run("CreatePoll", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		pollEntityFactory := mocks.NewMockPollEntityFactory(ctrl)
		pollEntityFactory.EXPECT().CreatePollEntity().Return(pollEntity)

		pollRepository := mocks.NewMockPollRepository(ctrl)
		pollRepository.EXPECT().SavePoll(pollEntity).Return(nil)

		data := &models.PollCreate{
			Title: "0",
		}
		pollService := &pollServiceImpl{
			pollEntityFactory: pollEntityFactory,
			pollRepository:    pollRepository,
		}
		response, err := pollService.CreatePoll(data)

		assert.IsType(t, pollEntity, response)
		assert.Nil(t, err)
		assert.Equal(t, data.Title, pollEntity.Title)
	})

	t.Run("GetPoll", func(t *testing.T) {
		pollId := new(models.PollId)
		pollEntity := new(entities.PollEntity)
		pollRepository := mocks.NewMockPollRepository(ctrl)
		pollRepository.EXPECT().GetPoll(pollId).Return(pollEntity, nil)

		pollService := &pollServiceImpl{pollRepository: pollRepository}
		response, err := pollService.GetPoll(pollId)

		assert.Equal(t, pollEntity, response)
		assert.Nil(t, err)
	})

	t.Run("UpdatePoll", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		pollRepository := mocks.NewMockPollRepository(ctrl)
		pollRepository.EXPECT().SavePoll(pollEntity).Return(nil)

		data := &models.PollUpdate{
			Title: "0",
		}
		pollService := &pollServiceImpl{pollRepository: pollRepository}
		assert.Nil(t, pollService.UpdatePoll(pollEntity, data))

		assert.Equal(t, data.Title, pollEntity.Title)
		assert.NotNil(t, pollEntity.Updated)
	})

	t.Run("DeletePoll", func(t *testing.T) {
		pollEntity := new(entities.PollEntity)
		pollRepository := mocks.NewMockPollRepository(ctrl)
		pollRepository.EXPECT().RemovePoll(pollEntity).Return(nil)

		pollService := &pollServiceImpl{pollRepository: pollRepository}
		assert.Nil(t, pollService.DeletePoll(pollEntity))
	})
}
