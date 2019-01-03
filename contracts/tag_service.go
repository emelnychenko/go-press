package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	TagService interface {
		ListTags(tagPaginationQuery *models.TagPaginationQuery) (*models.PaginationResult, common.Error)
		GetTag(tagId *models.TagId) (tagEntity *entities.TagEntity, err common.Error)
		CreateTag(data *models.TagCreate) (tagEntity *entities.TagEntity, err common.Error)
		UpdateTag(tagEntity *entities.TagEntity, data *models.TagUpdate) (err common.Error)
		DeleteTag(tagEntity *entities.TagEntity) (err common.Error)
	}
)
