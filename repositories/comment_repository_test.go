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

func TestCommentRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbPaginator := mocks.NewMockDbPaginator(ctrl)
	db, _ := gorm.Open(mocket.DriverName, "")
	commentRepository, isCommentRepository := NewCommentRepository(db, dbPaginator).(*commentRepositoryImpl)

	assert.True(t, isCommentRepository)
	assert.Equal(t, db, commentRepository.db)
	assert.Equal(t, dbPaginator, commentRepository.dbPaginator)

	commentId := models.NewModelId()
	commonReply := []map[string]interface{}{{
		"id": commentId.String(),
	}}

	t.Run("ListComments", func(t *testing.T) {
		commentPaginationQuery := &models.CommentPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), commentPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(nil)

		paginationResult, err := commentRepository.ListComments(commentPaginationQuery)
		assert.IsType(t, []*entities.CommentEntity{}, paginationResult.Data)
		assert.Nil(t, err)
	})

	t.Run("ListComments:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		commentPaginationQuery := &models.CommentPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), commentPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		commentEntities, err := commentRepository.ListComments(commentPaginationQuery)
		assert.Nil(t, commentEntities)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetComment", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		commentEntity, err := commentRepository.GetComment(commentId)
		assert.IsType(t, new(entities.CommentEntity), commentEntity)
		assert.Nil(t, err)
	})

	t.Run("GetComment:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		commentEntity, err := commentRepository.GetComment(commentId)
		assert.NotNil(t, commentEntity)
		assert.Error(t, err)
	})

	t.Run("GetComment:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		commentEntity, err := commentRepository.GetComment(commentId)
		assert.NotNil(t, commentEntity)
		assert.Error(t, err, errors.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("SaveComment", func(t *testing.T) {
		mocket.Catcher.Reset()

		commentEntity := entities.NewCommentEntity()
		assert.Nil(t, commentRepository.SaveComment(commentEntity))
	})

	t.Run("SaveComment:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		commentEntity := new(entities.CommentEntity)
		assert.Error(t, commentRepository.SaveComment(commentEntity))
	})

	t.Run("RemoveComment", func(t *testing.T) {
		mocket.Catcher.Reset()

		commentEntity := new(entities.CommentEntity)
		assert.Nil(t, commentRepository.RemoveComment(commentEntity))
	})

	t.Run("RemoveComment:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		commentEntity := new(entities.CommentEntity)
		assert.Error(t, commentRepository.RemoveComment(commentEntity))
	})
}
