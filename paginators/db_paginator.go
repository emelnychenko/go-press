package paginators

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
	"github.com/jinzhu/gorm"
)

type (
	dbPaginatorImpl struct {
	}
)

func NewDbPaginator() contracts.DbPaginator {
	return new(dbPaginatorImpl)
}

func (*dbPaginatorImpl) Paginate(
	db *gorm.DB,
	paginationQuery *models.PaginationQuery,
	paginationData interface{},
	paginationTotal *int,
) (err common.Error) {
	countRecords := func(db *gorm.DB, model interface{}, err chan common.Error, paginationTotal *int) {
		if gormErr := db.Model(model).Count(paginationTotal).Error; nil != gormErr {
			err <- common.NewSystemErrorFromBuiltin(gormErr)
			return
		}

		err <- nil
	}

	countErr := make(chan common.Error, 1)
	go countRecords(db, paginationData, countErr, paginationTotal)

	paginationOffset := paginationQuery.Offset()
	selectErr := db.Limit(paginationQuery.Limit).Offset(paginationOffset).Find(paginationData).Error

	err = <-countErr
	close(countErr)

	if nil != err {
		return
	}

	if nil != selectErr {
		err = common.NewSystemErrorFromBuiltin(selectErr)
		return
	}
	return
}
