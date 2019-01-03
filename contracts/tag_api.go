package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	TagApi interface {
		ListTags(tagPaginationQuery *models.TagPaginationQuery) (*models.PaginationResult, common.Error)
		GetTag(tagId *models.TagId) (tag *models.Tag, err common.Error)
		CreateTag(data *models.TagCreate) (tag *models.Tag, err common.Error)
		UpdateTag(tagId *models.TagId, data *models.TagUpdate) (err common.Error)
		DeleteTag(tagId *models.TagId) (err common.Error)
	}
)
