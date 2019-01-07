package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	CommentApi interface {
		ListComments(commentPaginationQuery *models.CommentPaginationQuery) (*models.PaginationResult, errors.Error)
		GetComment(commentId *models.CommentId) (comment *models.Comment, err errors.Error)
		CreateComment(data *models.CommentCreate) (comment *models.Comment, err errors.Error)
		UpdateComment(commentId *models.CommentId, data *models.CommentUpdate) (err errors.Error)
		DeleteComment(commentId *models.CommentId) (err errors.Error)
	}
)
