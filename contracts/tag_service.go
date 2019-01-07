package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	TagService interface {
		ListTags(tagPaginationQuery *models.TagPaginationQuery) (*models.PaginationResult, errors.Error)
		GetTag(tagId *models.TagId) (tagEntity *entities.TagEntity, err errors.Error)
		CreateTag(data *models.TagCreate) (tagEntity *entities.TagEntity, err errors.Error)
		UpdateTag(tagEntity *entities.TagEntity, data *models.TagUpdate) (err errors.Error)
		DeleteTag(tagEntity *entities.TagEntity) (err errors.Error)

		GetTagXrefs(*entities.TagEntity) ([]*entities.TagXrefEntity, errors.Error)
		GetTagObjectXrefs(models.Object) ([]*entities.TagXrefEntity, errors.Error)
		GetTagXref(*entities.TagEntity, models.Object) (*entities.TagXrefEntity, errors.Error)
		CreateTagXref(*entities.TagEntity, models.Object) (*entities.TagXrefEntity, errors.Error)
		DeleteTagXref(*entities.TagXrefEntity) errors.Error
	}
)
