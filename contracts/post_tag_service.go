package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostTagService interface {
		ListPostTags(*entities.PostEntity, *models.TagPaginationQuery) (*models.PaginationResult, errors.Error)
		AddPostTag(*entities.PostEntity, *entities.TagEntity) errors.Error
		RemovePostTag(*entities.PostEntity, *entities.TagEntity) errors.Error
	}
)
