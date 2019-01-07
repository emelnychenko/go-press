package repositories

import (
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCategoryRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbPaginator := mocks.NewMockDbPaginator(ctrl)
	categoryTreeBuilder := mocks.NewMockCategoryTreeBuilder(ctrl)
	categoryNestedSetBuilder := mocks.NewMockCategoryNestedSetBuilder(ctrl)
	db, _ := gorm.Open(mocket.DriverName, "")

	categoryRepository, isCategoryRepository := NewCategoryRepository(
		db, dbPaginator, categoryTreeBuilder, categoryNestedSetBuilder).(*categoryRepositoryImpl)

	assert.True(t, isCategoryRepository)
	assert.Equal(t, db, categoryRepository.db)
	assert.Equal(t, dbPaginator, categoryRepository.dbPaginator)
	assert.Equal(t, categoryTreeBuilder, categoryRepository.categoryTreeBuilder)
	assert.Equal(t, categoryNestedSetBuilder, categoryRepository.categoryNestedSetBuilder)

	categoryId := models.NewModelId()
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
		systemErr := errors.NewUnknownError()
		categoryPaginationQuery := &models.CategoryPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), categoryPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		categoryEntities, err := categoryRepository.ListCategories(categoryPaginationQuery)
		assert.Nil(t, categoryEntities)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetCategories", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		results, err := categoryRepository.GetCategories()
		assert.NotNil(t, results)
		assert.Nil(t, err)
	})

	t.Run("GetCategories:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT *`).WithError(gorm.ErrInvalidSQL)

		results, err := categoryRepository.GetCategories()
		assert.NotNil(t, results)
		assert.Error(t, err)
	})

	t.Run("GetCategoriesExcept", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		categoryEntity := &entities.CategoryEntity{Id: new(models.CategoryId)}
		results, err := categoryRepository.GetCategoriesExcept(categoryEntity)
		assert.NotNil(t, results)
		assert.Nil(t, err)
	})

	t.Run("GetCategoriesExcept:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT *`).WithError(gorm.ErrInvalidSQL)

		categoryEntity := &entities.CategoryEntity{Id: new(models.CategoryId)}
		results, err := categoryRepository.GetCategoriesExcept(categoryEntity)
		assert.NotNil(t, results)
		assert.Error(t, err)
	})

	t.Run("GetCategoriesTree", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		categoryEntityTree := new(entities.CategoryEntityTree)
		categoryTreeBuilder.EXPECT().BuildCategoryEntityTree(gomock.Any()).Return(categoryEntityTree, nil)

		result, err := categoryRepository.GetCategoriesTree()
		assert.Equal(t, categoryEntityTree, result)
		assert.Nil(t, err)
	})

	t.Run("GetCategoriesTree:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT *`).WithError(gorm.ErrInvalidSQL)

		result, err := categoryRepository.GetCategoriesTree()
		assert.Nil(t, result)
		assert.Error(t, err)
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
		assert.Error(t, err, errors.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("GetCategoryTree", func(t *testing.T) {
		mocket.Catcher.Reset()
		mocket.Catcher.NewMock().WithQuery("SELECT *").WithReply(commonReply)
		mocket.Catcher.NewMock().WithQuery("SELECT *").WithReply(commonReply)

		categoryEntityTree := new(entities.CategoryEntityTree)
		categoryTreeBuilder.EXPECT().BuildCategoryEntityTree(gomock.Any()).Return(categoryEntityTree, nil)

		result, err := categoryRepository.GetCategoryTree(categoryId)
		assert.Equal(t, categoryEntityTree, result)
		assert.Nil(t, err)
	})

	t.Run("GetCategoryTree:GetCategoryError", func(t *testing.T) {
		mocket.Catcher.Reset()
		mocket.Catcher.NewMock().WithQuery("SELECT *").WithError(gorm.ErrInvalidSQL)
		mocket.Catcher.NewMock().WithQuery("SELECT *").WithReply(commonReply)

		result, err := categoryRepository.GetCategoryTree(categoryId)
		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("GetCategoryTree:GormError", func(t *testing.T) {
		mocket.Catcher.Reset()
		mocket.Catcher.NewMock().WithQuery(`SELECT * FROM "categories"  WHERE (id`).WithReply(commonReply)
		mocket.Catcher.NewMock().WithQuery(`SELECT * FROM "categories"  WHERE (left`).WithError(gorm.ErrInvalidSQL)

		result, err := categoryRepository.GetCategoryTree(categoryId)
		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("SaveCategory", func(t *testing.T) {
		mocket.Catcher.Reset()

		categoryEntity := entities.NewCategoryEntity()
		assert.Nil(t, categoryRepository.SaveCategory(categoryEntity))
	})

	t.Run("SaveCategory:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		categoryEntity := new(entities.CategoryEntity)
		assert.Error(t, categoryRepository.SaveCategory(categoryEntity))
	})

	t.Run("RemoveCategory", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("DELETE")

		categoryEntity := new(entities.CategoryEntity)
		assert.Nil(t, categoryRepository.RemoveCategory(categoryEntity))
	})

	t.Run("RemoveCategory:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		categoryEntity := new(entities.CategoryEntity)
		assert.Error(t, categoryRepository.RemoveCategory(categoryEntity))
	})

	t.Run("GetCategoryXrefs", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		categoryEntity := &entities.CategoryEntity{Id: new(models.CategoryId)}

		results, err := categoryRepository.GetCategoryXrefs(categoryEntity)
		assert.NotNil(t, results)
		assert.Nil(t, err)
	})

	t.Run("GetCategoryXrefs:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT *`).WithError(gorm.ErrInvalidSQL)
		categoryEntity := &entities.CategoryEntity{Id: new(models.CategoryId)}

		results, err := categoryRepository.GetCategoryXrefs(categoryEntity)
		assert.NotNil(t, results)
		assert.Error(t, err)
	})

	t.Run("GetCategoryObjectXrefs", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		postEntity := &entities.PostEntity{Id: new(models.PostId)}

		results, err := categoryRepository.GetCategoryObjectXrefs(postEntity)
		assert.NotNil(t, results)
		assert.Nil(t, err)
	})

	t.Run("GetCategoryObjectXrefs:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT *`).WithError(gorm.ErrInvalidSQL)
		postEntity := &entities.PostEntity{Id: new(models.PostId)}

		results, err := categoryRepository.GetCategoryObjectXrefs(postEntity)
		assert.NotNil(t, results)
		assert.Error(t, err)
	})

	t.Run("GetCategoryXref", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		categoryId := new(models.CategoryId)
		categoryEntity := &entities.CategoryEntity{Id: categoryId}
		postId := new(models.PostId)
		postEntity := &entities.PostEntity{Id: postId}

		categoryXrefEntity, err := categoryRepository.GetCategoryXref(categoryEntity, postEntity)
		assert.IsType(t, new(entities.CategoryXrefEntity), categoryXrefEntity)
		assert.Nil(t, err)
	})

	t.Run("GetCategoryXref:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		categoryId := new(models.CategoryId)
		categoryEntity := &entities.CategoryEntity{Id: categoryId}
		postId := new(models.PostId)
		postEntity := &entities.PostEntity{Id: postId}

		categoryXrefEntity, err := categoryRepository.GetCategoryXref(categoryEntity, postEntity)
		assert.NotNil(t, categoryXrefEntity)
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotFound, err.Code())
	})

	t.Run("GetCategoryXref:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		categoryId := new(models.CategoryId)
		categoryEntity := &entities.CategoryEntity{Id: categoryId}
		postId := new(models.PostId)
		postEntity := &entities.PostEntity{Id: postId}

		categoryXrefEntity, err := categoryRepository.GetCategoryXref(categoryEntity, postEntity)
		assert.NotNil(t, categoryXrefEntity)
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, err.Code())
	})

	t.Run("SaveCategoryXref", func(t *testing.T) {
		mocket.Catcher.Reset()

		categoryXrefEntity := &entities.CategoryXrefEntity{
			CategoryId: new(models.CategoryId), ObjectId: new(models.ObjectId),
		}
		assert.Nil(t, categoryRepository.SaveCategoryXref(categoryXrefEntity))
	})

	t.Run("SaveCategoryXref:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		categoryXrefEntity := new(entities.CategoryXrefEntity)
		assert.Error(t, categoryRepository.SaveCategoryXref(categoryXrefEntity))
	})

	t.Run("RemoveCategoryXref", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("DELETE")

		categoryXrefEntity := new(entities.CategoryXrefEntity)
		assert.Nil(t, categoryRepository.RemoveCategoryXref(categoryXrefEntity))
	})

	t.Run("RemoveCategoryXref:GormError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		categoryXrefEntity := new(entities.CategoryXrefEntity)
		assert.Error(t, categoryRepository.RemoveCategoryXref(categoryXrefEntity))
	})

	t.Run("ListObjectCategories", func(t *testing.T) {
		categoryObject := mocks.NewMockObject(ctrl)
		categoryObject.EXPECT().ObjectType().Return(models.ObjectType("object"))
		categoryObject.EXPECT().ObjectId().Return(new(models.ObjectId))

		categoryPaginationQuery := &models.CategoryPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), categoryPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(nil)

		paginationResult, err := categoryRepository.ListObjectCategories(categoryObject, categoryPaginationQuery)
		assert.IsType(t, []*entities.CategoryEntity{}, paginationResult.Data)
		assert.Nil(t, err)
	})

	t.Run("ListObjectCategories:GormError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		categoryObject := mocks.NewMockObject(ctrl)
		categoryObject.EXPECT().ObjectType().Return(models.ObjectType("object"))
		categoryObject.EXPECT().ObjectId().Return(new(models.ObjectId))

		categoryPaginationQuery := &models.CategoryPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), categoryPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		categoryEntities, err := categoryRepository.ListObjectCategories(categoryObject, categoryPaginationQuery)
		assert.Nil(t, categoryEntities)
		assert.Equal(t, systemErr, err)
	})
}
