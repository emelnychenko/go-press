package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	CommentService interface {
		ListComments(commentPaginationQuery *models.CommentPaginationQuery) (*models.PaginationResult, common.Error)
		GetComment(commentId *models.CommentId) (commentEntity *entities.CommentEntity, err common.Error)
		CreateComment(data *models.CommentCreate) (commentEntity *entities.CommentEntity, err common.Error)
		UpdateComment(commentEntity *entities.CommentEntity, data *models.CommentUpdate) (err common.Error)
		DeleteComment(commentEntity *entities.CommentEntity) (err common.Error)
	}
)
