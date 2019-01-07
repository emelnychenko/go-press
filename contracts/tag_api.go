package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	TagApi interface {
		ListTags(tagPaginationQuery *models.TagPaginationQuery) (*models.PaginationResult, errors.Error)
		GetTag(tagId *models.TagId) (tag *models.Tag, err errors.Error)
		CreateTag(data *models.TagCreate) (tag *models.Tag, err errors.Error)
		UpdateTag(tagId *models.TagId, data *models.TagUpdate) (err errors.Error)
		DeleteTag(tagId *models.TagId) (err errors.Error)
	}
)
