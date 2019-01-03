package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	CommentModelFactory interface {
		CreateCommentPaginationQuery() *models.CommentPaginationQuery
		CreateComment() *models.Comment
		CreateCommentCreate() *models.CommentCreate
		CreateCommentUpdate() *models.CommentUpdate
	}
)
