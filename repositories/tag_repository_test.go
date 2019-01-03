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

func TestTagRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbPaginator := mocks.NewMockDbPaginator(ctrl)
	db, _ := gorm.Open(mocket.DriverName, "")
	tagRepository, isTagRepository := NewTagRepository(db, dbPaginator).(*tagRepositoryImpl)

	assert.True(t, isTagRepository)
	assert.Equal(t, db, tagRepository.db)
	assert.Equal(t, dbPaginator, tagRepository.dbPaginator)

	tagId := common.NewModelId()
	commonReply := []map[string]interface{}{{
		"id": tagId.String(),
	}}

	t.Run("ListTags", func(t *testing.T) {
		tagPaginationQuery := &models.TagPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), tagPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(nil)

		paginationResult, err := tagRepository.ListTags(tagPaginationQuery)
		assert.IsType(t, []*entities.TagEntity{}, paginationResult.Data)
		assert.Nil(t, err)
	})

	t.Run("ListTags:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		tagPaginationQuery := &models.TagPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = errors.New("")
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), tagPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		tagEntities, err := tagRepository.ListTags(tagPaginationQuery)
		assert.Nil(t, tagEntities)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetTag", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		tagEntity, err := tagRepository.GetTag(tagId)
		assert.IsType(t, new(entities.TagEntity), tagEntity)
		assert.Nil(t, err)
	})

	t.Run("GetTag:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		tagEntity, err := tagRepository.GetTag(tagId)
		assert.NotNil(t, tagEntity)
		assert.Error(t, err)
	})

	t.Run("GetTag:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		tagEntity, err := tagRepository.GetTag(tagId)
		assert.NotNil(t, tagEntity)
		assert.Error(t, err, common.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("SaveTag", func(t *testing.T) {
		mocket.Catcher.Reset()

		tagEntity := entities.NewTagEntity()
		assert.Nil(t, tagRepository.SaveTag(tagEntity))
	})

	t.Run("SaveTag:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		tagEntity := new(entities.TagEntity)
		assert.Error(t, tagRepository.SaveTag(tagEntity))
	})

	t.Run("RemoveTag", func(t *testing.T) {
		mocket.Catcher.Reset()

		tagEntity := new(entities.TagEntity)
		assert.Nil(t, tagRepository.RemoveTag(tagEntity))
	})

	t.Run("RemoveTag:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		tagEntity := new(entities.TagEntity)
		assert.Error(t, tagRepository.RemoveTag(tagEntity))
	})
}
