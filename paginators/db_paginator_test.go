package paginators

import (
	"errors"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbPaginator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	db, _ := gorm.Open(mocket.DriverName, "")

	type TestEntity struct {
	}

	resultCount := 5
	entityReply := make([]map[string]interface{}, resultCount)
	countReply := []map[string]interface{}{{"count(*)": resultCount}}

	t.Run("NewDbPaginator", func(t *testing.T) {
		_, isDbPaginator := NewDbPaginator().(*dbPaginatorImpl)

		assert.True(t, isDbPaginator)
	})

	t.Run("Paginate", func(t *testing.T) {
		mocket.Catcher.Reset()
		mocket.Catcher.NewMock().WithQuery("SELECT *").WithReply(entityReply)
		mocket.Catcher.NewMock().WithQuery("SELECT count(*)").WithReply(countReply)

		paginationQuery := &models.PaginationQuery{Limit: 10, Page: 1, Start: 0}
		var testEntities []*TestEntity
		var paginationTotal int

		dbPaginator := &dbPaginatorImpl{}
		err := dbPaginator.Paginate(db, paginationQuery, &testEntities, &paginationTotal)

		assert.Nil(t, err)
		assert.Equal(t, resultCount, paginationTotal)
	})

	t.Run("Paginate:CountError", func(t *testing.T) {
		mocket.Catcher.Reset()
		mocket.Catcher.NewMock().WithQuery("SELECT *").WithReply(entityReply)
		mocket.Catcher.NewMock().WithQuery("SELECT count(*)").WithError(errors.New(""))

		paginationQuery := &models.PaginationQuery{Limit: 10, Page: 1, Start: 0}
		var testEntities []*TestEntity
		var paginationTotal int

		dbPaginator := &dbPaginatorImpl{}
		err := dbPaginator.Paginate(db, paginationQuery, &testEntities, &paginationTotal)

		assert.Error(t, err)
	})

	t.Run("Paginate:SelectError", func(t *testing.T) {
		mocket.Catcher.Reset()
		mocket.Catcher.NewMock().OneTime().WithQuery("SELECT *").WithError(errors.New(""))
		mocket.Catcher.NewMock().WithQuery("SELECT count(*)").WithReply(countReply)

		paginationQuery := &models.PaginationQuery{Limit: 10, Page: 1, Start: 0}
		var testEntities []*TestEntity
		var paginationTotal int

		dbPaginator := &dbPaginatorImpl{}
		err := dbPaginator.Paginate(db, paginationQuery, &testEntities, &paginationTotal)

		assert.Error(t, err)
	})
}
