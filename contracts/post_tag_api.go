package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostTagApi interface {
		ListPostTags(*models.PostId, *models.TagPaginationQuery) (*models.PaginationResult, errors.Error)
		AddPostTag(*models.PostId, *models.TagId) errors.Error
		RemovePostTag(*models.PostId, *models.TagId) errors.Error
	}
)
