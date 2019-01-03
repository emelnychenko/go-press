package repositories

import (
	"errors"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChannelRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbPaginator := mocks.NewMockDbPaginator(ctrl)
	db, _ := gorm.Open(mocket.DriverName, "")
	channelRepository, isChannelRepository := NewChannelRepository(db, dbPaginator).(*channelRepositoryImpl)

	assert.True(t, isChannelRepository)
	assert.Equal(t, db, channelRepository.db)
	assert.Equal(t, dbPaginator, channelRepository.dbPaginator)

	channelId := common.NewModelId()
	commonReply := []map[string]interface{}{{
		"id": channelId.String(),
	}}

	t.Run("ListChannels", func(t *testing.T) {
		channelPaginationQuery := &models.ChannelPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), channelPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(nil)

		paginationResult, err := channelRepository.ListChannels(channelPaginationQuery)
		assert.IsType(t, []*entities.ChannelEntity{}, paginationResult.Data)
		assert.Nil(t, err)
	})

	t.Run("ListChannels:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		channelPaginationQuery := &models.ChannelPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = errors.New("")
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), channelPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		channelEntities, err := channelRepository.ListChannels(channelPaginationQuery)
		assert.Nil(t, channelEntities)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetChannel", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		channelEntity, err := channelRepository.GetChannel(channelId)
		assert.IsType(t, new(entities.ChannelEntity), channelEntity)
		assert.Nil(t, err)
	})

	t.Run("GetChannel:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		channelEntity, err := channelRepository.GetChannel(channelId)
		assert.NotNil(t, channelEntity)
		assert.Error(t, err)
	})

	t.Run("GetChannel:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		channelEntity, err := channelRepository.GetChannel(channelId)
		assert.NotNil(t, channelEntity)
		assert.Error(t, err, common.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("SaveChannel", func(t *testing.T) {
		mocket.Catcher.Reset()

		channelEntity := entities.NewChannelEntity()
		assert.Nil(t, channelRepository.SaveChannel(channelEntity))
	})

	t.Run("SaveChannel:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		channelEntity := new(entities.ChannelEntity)
		assert.Error(t, channelRepository.SaveChannel(channelEntity))
	})

	t.Run("RemoveChannel", func(t *testing.T) {
		mocket.Catcher.Reset()

		channelEntity := new(entities.ChannelEntity)
		assert.Nil(t, channelRepository.RemoveChannel(channelEntity))
	})

	t.Run("RemoveChannel:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		channelEntity := new(entities.ChannelEntity)
		assert.Error(t, channelRepository.RemoveChannel(channelEntity))
	})
}
