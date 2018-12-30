package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
	"github.com/jinzhu/gorm"
)

type (
	DbPaginator interface {
		Paginate(
			db *gorm.DB,
			paginationQuery *models.PaginationQuery,
			paginationData interface{},
			paginationTotal *int,
		) (err common.Error)
	}
)
