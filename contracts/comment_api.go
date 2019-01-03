package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	CommentApi interface {
		ListComments(commentPaginationQuery *models.CommentPaginationQuery) (*models.PaginationResult, common.Error)
		GetComment(commentId *models.CommentId) (comment *models.Comment, err common.Error)
		CreateComment(data *models.CommentCreate) (comment *models.Comment, err common.Error)
		UpdateComment(commentId *models.CommentId, data *models.CommentUpdate) (err common.Error)
		DeleteComment(commentId *models.CommentId) (err common.Error)
	}
)
