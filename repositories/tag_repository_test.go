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
	"net/http"
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

	tagId := models.NewModelId()
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
		systemErr := errors.NewUnknownError()
		tagPaginationQuery := &models.TagPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL
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
		assert.Error(t, err, errors.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("SaveTag", func(t *testing.T) {
		mocket.Catcher.Reset()

		tagEntity := entities.NewTagEntity()
		assert.Nil(t, tagRepository.SaveTag(tagEntity))
	})

	t.Run("SaveTag:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		tagEntity := new(entities.TagEntity)
		assert.Error(t, tagRepository.SaveTag(tagEntity))
	})

	t.Run("RemoveTag", func(t *testing.T) {
		mocket.Catcher.Reset()

		tagEntity := new(entities.TagEntity)
		assert.Nil(t, tagRepository.RemoveTag(tagEntity))
	})

	t.Run("RemoveTag:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		tagEntity := new(entities.TagEntity)
		assert.Error(t, tagRepository.RemoveTag(tagEntity))
	})

	t.Run("GetTagXrefs", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		tagEntity := &entities.TagEntity{Id: new(models.TagId)}

		results, err := tagRepository.GetTagXrefs(tagEntity)
		assert.NotNil(t, results)
		assert.Nil(t, err)
	})

	t.Run("GetTagXrefs:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT *`).WithError(gorm.ErrInvalidSQL)
		tagEntity := &entities.TagEntity{Id: new(models.TagId)}

		results, err := tagRepository.GetTagXrefs(tagEntity)
		assert.NotNil(t, results)
		assert.Error(t, err)
	})

	t.Run("GetTagObjectXrefs", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		postEntity := &entities.PostEntity{Id: new(models.PostId)}

		results, err := tagRepository.GetTagObjectXrefs(postEntity)
		assert.NotNil(t, results)
		assert.Nil(t, err)
	})

	t.Run("GetTagObjectXrefs:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT *`).WithError(gorm.ErrInvalidSQL)
		postEntity := &entities.PostEntity{Id: new(models.PostId)}

		results, err := tagRepository.GetTagObjectXrefs(postEntity)
		assert.NotNil(t, results)
		assert.Error(t, err)
	})

	t.Run("GetTagXref", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		tagId := new(models.TagId)
		tagEntity := &entities.TagEntity{Id: tagId}
		postId := new(models.PostId)
		postEntity := &entities.PostEntity{Id: postId}

		tagXrefEntity, err := tagRepository.GetTagXref(tagEntity, postEntity)
		assert.IsType(t, new(entities.TagXrefEntity), tagXrefEntity)
		assert.Nil(t, err)
	})

	t.Run("GetTagXref:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		tagId := new(models.TagId)
		tagEntity := &entities.TagEntity{Id: tagId}
		postId := new(models.PostId)
		postEntity := &entities.PostEntity{Id: postId}

		tagXrefEntity, err := tagRepository.GetTagXref(tagEntity, postEntity)
		assert.NotNil(t, tagXrefEntity)
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("GetTagXref:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		tagId := new(models.TagId)
		tagEntity := &entities.TagEntity{Id: tagId}
		postId := new(models.PostId)
		postEntity := &entities.PostEntity{Id: postId}

		tagXrefEntity, err := tagRepository.GetTagXref(tagEntity, postEntity)
		assert.NotNil(t, tagXrefEntity)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.Code())
	})

	t.Run("SaveTagXref", func(t *testing.T) {
		mocket.Catcher.Reset()

		tagXrefEntity := &entities.TagXrefEntity{TagId: new(models.TagId), ObjectId: new(models.ObjectId)}
		assert.Nil(t, tagRepository.SaveTagXref(tagXrefEntity))
	})

	t.Run("SaveTagXref:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		tagXrefEntity := new(entities.TagXrefEntity)
		assert.Error(t, tagRepository.SaveTagXref(tagXrefEntity))
	})

	t.Run("RemoveTagXref", func(t *testing.T) {
		mocket.Catcher.Reset()

		tagXrefEntity := new(entities.TagXrefEntity)
		assert.Nil(t, tagRepository.RemoveTagXref(tagXrefEntity))
	})

	t.Run("RemoveTagXref:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		tagXrefEntity := new(entities.TagXrefEntity)
		assert.Error(t, tagRepository.RemoveTagXref(tagXrefEntity))
	})
}
