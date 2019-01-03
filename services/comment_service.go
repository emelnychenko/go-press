package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	commentServiceImpl struct {
		commentEntityFactory contracts.CommentEntityFactory
		commentRepository    contracts.CommentRepository
	}
)

func NewCommentService(
	commentEntityFactory contracts.CommentEntityFactory,
	commentRepository contracts.CommentRepository,
) (commentService contracts.CommentService) {
	return &commentServiceImpl{
		commentEntityFactory,
		commentRepository,
	}
}

func (s *commentServiceImpl) ListComments(
	commentPaginationQuery *models.CommentPaginationQuery,
) (*models.PaginationResult, common.Error) {
	return s.commentRepository.ListComments(commentPaginationQuery)
}

func (s *commentServiceImpl) GetComment(commentId *models.CommentId) (*entities.CommentEntity, common.Error) {
	return s.commentRepository.GetComment(commentId)
}

func (s *commentServiceImpl) CreateComment(data *models.CommentCreate) (
	commentEntity *entities.CommentEntity, err common.Error,
) {
	commentEntity = s.commentEntityFactory.CreateCommentEntity()
	commentEntity.Content = data.Content

	err = s.commentRepository.SaveComment(commentEntity)
	return
}

func (s *commentServiceImpl) UpdateComment(
	commentEntity *entities.CommentEntity, data *models.CommentUpdate,
) common.Error {
	commentEntity.Content = data.Content

	updated := time.Now().UTC()
	commentEntity.Updated = &updated

	return s.commentRepository.SaveComment(commentEntity)
}

func (s *commentServiceImpl) DeleteComment(commentEntity *entities.CommentEntity) common.Error {
	return s.commentRepository.RemoveComment(commentEntity)
}
