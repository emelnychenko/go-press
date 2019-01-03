package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	TagModelFactory interface {
		CreateTagPaginationQuery() *models.TagPaginationQuery
		CreateTag() *models.Tag
		CreateTagCreate() *models.TagCreate
		CreateTagUpdate() *models.TagUpdate
	}
)
