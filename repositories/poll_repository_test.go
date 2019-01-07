package repositories

import (
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbPaginator := mocks.NewMockDbPaginator(ctrl)
	db, _ := gorm.Open(mocket.DriverName, "")
	pollRepository, isPollRepository := NewPollRepository(db, dbPaginator).(*pollRepositoryImpl)

	assert.True(t, isPollRepository)
	assert.Equal(t, db, pollRepository.db)
	assert.Equal(t, dbPaginator, pollRepository.dbPaginator)

	pollId := models.NewModelId()
	commonReply := []map[string]interface{}{{
		"id": pollId.String(),
	}}

	t.Run("ListPolls", func(t *testing.T) {
		pollPaginationQuery := &models.PollPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), pollPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(nil)

		paginationResult, err := pollRepository.ListPolls(pollPaginationQuery)
		assert.IsType(t, []*entities.PollEntity{}, paginationResult.Data)
		assert.Nil(t, err)
	})

	t.Run("ListPolls:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		pollPaginationQuery := &models.PollPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), pollPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		pollEntities, err := pollRepository.ListPolls(pollPaginationQuery)
		assert.Nil(t, pollEntities)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetPoll", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		pollEntity, err := pollRepository.GetPoll(pollId)
		assert.IsType(t, new(entities.PollEntity), pollEntity)
		assert.Nil(t, err)
	})

	t.Run("GetPoll:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		pollEntity, err := pollRepository.GetPoll(pollId)
		assert.NotNil(t, pollEntity)
		assert.Error(t, err)
	})

	t.Run("GetPoll:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		pollEntity, err := pollRepository.GetPoll(pollId)
		assert.NotNil(t, pollEntity)
		assert.Error(t, err, errors.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("SavePoll", func(t *testing.T) {
		mocket.Catcher.Reset()

		pollEntity := entities.NewPollEntity()
		assert.Nil(t, pollRepository.SavePoll(pollEntity))
	})

	t.Run("SavePoll:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		pollEntity := new(entities.PollEntity)
		assert.Error(t, pollRepository.SavePoll(pollEntity))
	})

	t.Run("RemovePoll", func(t *testing.T) {
		mocket.Catcher.Reset()

		pollEntity := new(entities.PollEntity)
		assert.Nil(t, pollRepository.RemovePoll(pollEntity))
	})

	t.Run("RemovePoll:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		pollEntity := new(entities.PollEntity)
		assert.Error(t, pollRepository.RemovePoll(pollEntity))
	})
}
