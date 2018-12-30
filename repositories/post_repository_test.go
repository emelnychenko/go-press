package repositories

import (
	"errors"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbPaginator := mocks.NewMockDbPaginator(ctrl)
	db, _ := gorm.Open(mocket.DriverName, "")
	postRepository, isPostRepository := NewPostRepository(db, dbPaginator).(*postRepositoryImpl)

	assert.True(t, isPostRepository)
	assert.Equal(t, db, postRepository.db)
	assert.Equal(t, dbPaginator, postRepository.dbPaginator)

	postId := common.NewModelId()
	commonReply := []map[string]interface{}{{
		"id": postId.String(),
	}}

	t.Run("ListPosts", func(t *testing.T) {
		postPaginationQuery := &models.PostPaginationQuery{
			Status: enums.PostDraftStatus,
			Privacy: enums.PostPublicPrivacy,
			Author: "Test",
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), postPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(nil)

		paginationResult, err := postRepository.ListPosts(postPaginationQuery)
		assert.IsType(t, []*entities.PostEntity{}, paginationResult.Data)
		assert.Nil(t, err)
	})

	t.Run("ListPosts:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		postPaginationQuery := &models.PostPaginationQuery{
			Status: enums.PostDraftStatus,
			Privacy: enums.PostPublicPrivacy,
			Author: "Test",
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = errors.New("")
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), postPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		postEntities, err := postRepository.ListPosts(postPaginationQuery)
		assert.Nil(t, postEntities)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetPost", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		postEntity, err := postRepository.GetPost(postId)
		assert.IsType(t, new(entities.PostEntity), postEntity)
		assert.Nil(t, err)
	})

	t.Run("GetPost:PostNotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		postEntity, err := postRepository.GetPost(postId)
		assert.NotNil(t, postEntity)
		assert.Error(t, err)
	})

	t.Run("GetPost:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		postEntity, err := postRepository.GetPost(postId)
		assert.NotNil(t, postEntity)
		assert.Error(t, err)
	})

	t.Run("SavePost", func(t *testing.T) {
		mocket.Catcher.Reset()

		postEntity := entities.NewPostEntity()
		assert.Nil(t, postRepository.SavePost(postEntity))
	})

	t.Run("SavePost:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		postEntity := new(entities.PostEntity)
		assert.Error(t, postRepository.SavePost(postEntity))
	})

	t.Run("RemovePost", func(t *testing.T) {
		mocket.Catcher.Reset()

		postEntity := new(entities.PostEntity)
		assert.Nil(t, postRepository.RemovePost(postEntity))
	})

	t.Run("RemovePost:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		postEntity := new(entities.PostEntity)
		assert.Error(t, postRepository.RemovePost(postEntity))
	})
}
