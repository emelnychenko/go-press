package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	pollServiceImpl struct {
		pollEntityFactory contracts.PollEntityFactory
		pollRepository    contracts.PollRepository
	}
)

func NewPollService(
	pollEntityFactory contracts.PollEntityFactory,
	pollRepository contracts.PollRepository,
) (pollService contracts.PollService) {
	return &pollServiceImpl{
		pollEntityFactory,
		pollRepository,
	}
}

func (s *pollServiceImpl) ListPolls(
	pollPaginationQuery *models.PollPaginationQuery,
) (*models.PaginationResult, common.Error) {
	return s.pollRepository.ListPolls(pollPaginationQuery)
}

func (s *pollServiceImpl) GetPoll(pollId *models.PollId) (*entities.PollEntity, common.Error) {
	return s.pollRepository.GetPoll(pollId)
}

func (s *pollServiceImpl) CreatePoll(data *models.PollCreate) (pollEntity *entities.PollEntity, err common.Error) {
	pollEntity = s.pollEntityFactory.CreatePollEntity()
	pollEntity.Title = data.Title

	err = s.pollRepository.SavePoll(pollEntity)
	return
}

func (s *pollServiceImpl) UpdatePoll(pollEntity *entities.PollEntity, data *models.PollUpdate) common.Error {
	pollEntity.Title = data.Title

	updated := time.Now().UTC()
	pollEntity.Updated = &updated

	return s.pollRepository.SavePoll(pollEntity)
}

func (s *pollServiceImpl) DeletePoll(pollEntity *entities.PollEntity) common.Error {
	return s.pollRepository.RemovePoll(pollEntity)
}
