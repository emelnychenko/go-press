package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	TagRepository interface {
		ListTags(tagPaginationQuery *models.TagPaginationQuery) (*models.PaginationResult, common.Error)
		GetTag(tagId *models.TagId) (tagEntity *entities.TagEntity, err common.Error)
		SaveTag(tagEntity *entities.TagEntity) (err common.Error)
		RemoveTag(tagEntity *entities.TagEntity) (err common.Error)
	}
)
