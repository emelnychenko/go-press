package paginators

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
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
) (err errors.Error) {
	countRecords := func(db *gorm.DB, model interface{}, err chan errors.Error, paginationTotal *int) {
		if gormErr := db.Model(model).Count(paginationTotal).Error; nil != gormErr {
			err <- errors.NewSystemErrorFromBuiltin(gormErr)
			return
		}

		err <- nil
	}

	countErr := make(chan errors.Error, 1)
	go countRecords(db, paginationData, countErr, paginationTotal)

	paginationOffset := paginationQuery.Offset()
	selectErr := db.Limit(paginationQuery.Limit).Offset(paginationOffset).Find(paginationData).Error

	err = <-countErr
	close(countErr)

	if nil != err {
		return
	}

	if nil != selectErr {
		err = errors.NewSystemErrorFromBuiltin(selectErr)
		return
	}
	return
}
