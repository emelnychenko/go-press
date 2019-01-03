package repositories

import (
	"errors"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbPaginator := mocks.NewMockDbPaginator(ctrl)
	db, _ := gorm.Open(mocket.DriverName, "")
	categoryRepository, isCategoryRepository := NewCategoryRepository(db, dbPaginator).(*categoryRepositoryImpl)

	assert.True(t, isCategoryRepository)
	assert.Equal(t, db, categoryRepository.db)
	assert.Equal(t, dbPaginator, categoryRepository.dbPaginator)

	categoryId := common.NewModelId()
	commonReply := []map[string]interface{}{{
		"id": categoryId.String(),
	}}

	t.Run("ListCategories", func(t *testing.T) {
		categoryPaginationQuery := &models.CategoryPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), categoryPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(nil)

		paginationResult, err := categoryRepository.ListCategories(categoryPaginationQuery)
		assert.IsType(t, []*entities.CategoryEntity{}, paginationResult.Data)
		assert.Nil(t, err)
	})

	t.Run("ListCategories:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		categoryPaginationQuery := &models.CategoryPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = errors.New("")
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), categoryPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		categoryEntities, err := categoryRepository.ListCategories(categoryPaginationQuery)
		assert.Nil(t, categoryEntities)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetCategory", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		categoryEntity, err := categoryRepository.GetCategory(categoryId)
		assert.IsType(t, new(entities.CategoryEntity), categoryEntity)
		assert.Nil(t, err)
	})

	t.Run("GetCategory:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		categoryEntity, err := categoryRepository.GetCategory(categoryId)
		assert.NotNil(t, categoryEntity)
		assert.Error(t, err)
	})

	t.Run("GetCategory:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		categoryEntity, err := categoryRepository.GetCategory(categoryId)
		assert.NotNil(t, categoryEntity)
		assert.Error(t, err, common.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("SaveCategory", func(t *testing.T) {
		mocket.Catcher.Reset()

		categoryEntity := entities.NewCategoryEntity()
		assert.Nil(t, categoryRepository.SaveCategory(categoryEntity))
	})

	t.Run("SaveCategory:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		categoryEntity := new(entities.CategoryEntity)
		assert.Error(t, categoryRepository.SaveCategory(categoryEntity))
	})

	t.Run("RemoveCategory", func(t *testing.T) {
		mocket.Catcher.Reset()

		categoryEntity := new(entities.CategoryEntity)
		assert.Nil(t, categoryRepository.RemoveCategory(categoryEntity))
	})

	t.Run("RemoveCategory:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		categoryEntity := new(entities.CategoryEntity)
		assert.Error(t, categoryRepository.RemoveCategory(categoryEntity))
	})
}
