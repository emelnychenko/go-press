package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	CommentService interface {
		ListComments(commentPaginationQuery *models.CommentPaginationQuery) (*models.PaginationResult, errors.Error)
		GetComment(commentId *models.CommentId) (commentEntity *entities.CommentEntity, err errors.Error)
		CreateComment(data *models.CommentCreate) (commentEntity *entities.CommentEntity, err errors.Error)
		UpdateComment(commentEntity *entities.CommentEntity, data *models.CommentUpdate) (err errors.Error)
		DeleteComment(commentEntity *entities.CommentEntity) (err errors.Error)
	}
)
