package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	TagRepository interface {
		ListTags(tagPaginationQuery *models.TagPaginationQuery) (*models.PaginationResult, errors.Error)
		GetTag(tagId *models.TagId) (tagEntity *entities.TagEntity, err errors.Error)
		SaveTag(tagEntity *entities.TagEntity) (err errors.Error)
		RemoveTag(tagEntity *entities.TagEntity) (err errors.Error)

		GetTagXrefs(*entities.TagEntity) ([]*entities.TagXrefEntity, errors.Error)
		GetTagObjectXrefs(models.Object) ([]*entities.TagXrefEntity, errors.Error)
		GetTagXref(*entities.TagEntity, models.Object) (*entities.TagXrefEntity, errors.Error)
		SaveTagXref(*entities.TagXrefEntity) errors.Error
		RemoveTagXref(*entities.TagXrefEntity) errors.Error

		ListObjectTags(models.Object, *models.TagPaginationQuery) (*models.PaginationResult, errors.Error)
	}
)
