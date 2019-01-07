package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	CommentRepository interface {
		ListComments(commentPaginationQuery *models.CommentPaginationQuery) (*models.PaginationResult, errors.Error)
		GetComment(commentId *models.CommentId) (commentEntity *entities.CommentEntity, err errors.Error)
		SaveComment(commentEntity *entities.CommentEntity) (err errors.Error)
		RemoveComment(commentEntity *entities.CommentEntity) (err errors.Error)
	}
)
