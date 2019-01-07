package repositories

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"github.com/jinzhu/gorm"
)

type (
	categoryRepositoryImpl struct {
		db                       *gorm.DB
		dbPaginator              contracts.DbPaginator
		categoryTreeBuilder      contracts.CategoryTreeBuilder
		categoryNestedSetBuilder contracts.CategoryNestedSetBuilder
	}
)

//NewCategoryRepository
func NewCategoryRepository(
	db *gorm.DB, dbPaginator contracts.DbPaginator,
	categoryTreeBuilder contracts.CategoryTreeBuilder,
	categoryNestedSetBuilder contracts.CategoryNestedSetBuilder,
) contracts.CategoryRepository {
	return &categoryRepositoryImpl{
		db,
		dbPaginator,
		categoryTreeBuilder,
		categoryNestedSetBuilder}
}

//ListCategories
func (r *categoryRepositoryImpl) ListCategories(
	categoryPaginationQuery *models.CategoryPaginationQuery,
) (paginationResult *models.PaginationResult, err common.Error) {
	paginationTotal, categoryEntities := 0, make([]*entities.CategoryEntity, categoryPaginationQuery.Limit)

	db := r.db.Model(&categoryEntities).Order("created desc")
	err = r.dbPaginator.Paginate(db, categoryPaginationQuery.PaginationQuery, &categoryEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: categoryEntities}
	return
}

//GetCategories
func (r *categoryRepositoryImpl) GetCategories() (categoryEntities []*entities.CategoryEntity, err common.Error) {
	if gormErr := r.db.Order("created asc").Find(&categoryEntities).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//GetCategoriesExcept
func (r *categoryRepositoryImpl) GetCategoriesExcept(categoryEntity *entities.CategoryEntity) (
	categoryEntities []*entities.CategoryEntity, err common.Error,
) {
	gormErr := r.db.Where("id != ?", categoryEntity.Id).Order("created asc").Find(&categoryEntities).Error

	if gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//GetCategoriesTree
func (r *categoryRepositoryImpl) GetCategoriesTree() (
	categoryEntityTree *entities.CategoryEntityTree, err common.Error,
) {
	var categoryEntities []*entities.CategoryEntity

	if gormErr := r.db.Order("created asc").Find(&categoryEntities).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
		return
	}

	categoryEntityTree, err = r.categoryTreeBuilder.BuildCategoryEntityTree(categoryEntities)
	return
}

//GetCategory
func (r *categoryRepositoryImpl) GetCategory(categoryId *models.CategoryId) (categoryEntity *entities.CategoryEntity, err common.Error) {
	categoryEntity = new(entities.CategoryEntity)

	if gormErr := r.db.First(categoryEntity, "id = ?", categoryId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewCategoryByIdNotFoundError(categoryId)
		} else {
			err = common.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

//GetCategoryTree
func (r *categoryRepositoryImpl) GetCategoryTree(categoryId *models.CategoryId) (
	categoryEntityTree *entities.CategoryEntityTree, err common.Error,
) {
	categoryEntity, err := r.GetCategory(categoryId)

	if nil != err {
		return
	}

	var categoryEntities []*entities.CategoryEntity

	gormErr := r.db.Where("left >= ?", categoryEntity.Left).
		Where("right <= ?", categoryEntity.Right).
		Order("created asc").
		Find(&categoryEntities).Error

	if gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
		return
	}

	categoryEntityTree, err = r.categoryTreeBuilder.BuildCategoryEntityTree(categoryEntities)
	return
}

//SaveCategory
func (r *categoryRepositoryImpl) SaveCategory(categoryEntity *entities.CategoryEntity) (err common.Error) {
	if gormErr := r.db.Save(categoryEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//RemoveCategory
func (r *categoryRepositoryImpl) RemoveCategory(categoryEntity *entities.CategoryEntity) (err common.Error) {
	if gormErr := r.db.Delete(categoryEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
