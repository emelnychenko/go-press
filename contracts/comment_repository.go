package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	CommentRepository interface {
		ListComments(commentPaginationQuery *models.CommentPaginationQuery) (*models.PaginationResult, common.Error)
		GetComment(commentId *models.CommentId) (commentEntity *entities.CommentEntity, err common.Error)
		SaveComment(commentEntity *entities.CommentEntity) (err common.Error)
		RemoveComment(commentEntity *entities.CommentEntity) (err common.Error)
	}
)
