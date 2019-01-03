package services

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewCommentService", func(t *testing.T) {
		commentEntityFactory := mocks.NewMockCommentEntityFactory(ctrl)
		commentRepository := mocks.NewMockCommentRepository(ctrl)

		commentService, isCommentService := NewCommentService(commentEntityFactory, commentRepository).(*commentServiceImpl)

		assert.True(t, isCommentService)
		assert.Equal(t, commentEntityFactory, commentService.commentEntityFactory)
		assert.Equal(t, commentRepository, commentService.commentRepository)
	})

	t.Run("ListComments", func(t *testing.T) {
		commentPaginationQuery := new(models.CommentPaginationQuery)

		var commentEntities *models.PaginationResult
		commentRepository := mocks.NewMockCommentRepository(ctrl)
		commentRepository.EXPECT().ListComments(commentPaginationQuery).Return(commentEntities, nil)

		commentService := &commentServiceImpl{commentRepository: commentRepository}
		response, err := commentService.ListComments(commentPaginationQuery)

		assert.Equal(t, commentEntities, response)
		assert.Nil(t, err)
	})

	t.Run("CreateComment", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		commentEntityFactory := mocks.NewMockCommentEntityFactory(ctrl)
		commentEntityFactory.EXPECT().CreateCommentEntity().Return(commentEntity)

		commentRepository := mocks.NewMockCommentRepository(ctrl)
		commentRepository.EXPECT().SaveComment(commentEntity).Return(nil)

		data := &models.CommentCreate{
			Content: "0",
		}
		commentService := &commentServiceImpl{
			commentEntityFactory: commentEntityFactory,
			commentRepository:    commentRepository,
		}
		response, err := commentService.CreateComment(data)

		assert.IsType(t, commentEntity, response)
		assert.Nil(t, err)
		assert.Equal(t, data.Content, commentEntity.Content)
	})

	t.Run("GetComment", func(t *testing.T) {
		commentId := new(models.CommentId)
		commentEntity := new(entities.CommentEntity)
		commentRepository := mocks.NewMockCommentRepository(ctrl)
		commentRepository.EXPECT().GetComment(commentId).Return(commentEntity, nil)

		commentService := &commentServiceImpl{commentRepository: commentRepository}
		response, err := commentService.GetComment(commentId)

		assert.Equal(t, commentEntity, response)
		assert.Nil(t, err)
	})

	t.Run("UpdateComment", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		commentRepository := mocks.NewMockCommentRepository(ctrl)
		commentRepository.EXPECT().SaveComment(commentEntity).Return(nil)

		data := &models.CommentUpdate{
			Content: "0",
		}
		commentService := &commentServiceImpl{commentRepository: commentRepository}
		assert.Nil(t, commentService.UpdateComment(commentEntity, data))

		assert.Equal(t, data.Content, commentEntity.Content)
		assert.NotNil(t, commentEntity.Updated)
	})

	t.Run("DeleteComment", func(t *testing.T) {
		commentEntity := new(entities.CommentEntity)
		commentRepository := mocks.NewMockCommentRepository(ctrl)
		commentRepository.EXPECT().RemoveComment(commentEntity).Return(nil)

		commentService := &commentServiceImpl{commentRepository: commentRepository}
		assert.Nil(t, commentService.DeleteComment(commentEntity))
	})
}
