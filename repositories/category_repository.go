package repositories

import (
	"fmt"
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
) (paginationResult *models.PaginationResult, err errors.Error) {
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
func (r *categoryRepositoryImpl) GetCategories() (categoryEntities []*entities.CategoryEntity, err errors.Error) {
	if gormErr := r.db.Order("created asc").Find(&categoryEntities).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//GetCategoriesExcept
func (r *categoryRepositoryImpl) GetCategoriesExcept(categoryEntity *entities.CategoryEntity) (
	categoryEntities []*entities.CategoryEntity, err errors.Error,
) {
	gormErr := r.db.Where("id != ?", categoryEntity.Id).Order("created asc").Find(&categoryEntities).Error

	if gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//GetCategoriesTree
func (r *categoryRepositoryImpl) GetCategoriesTree() (
	categoryEntityTree *entities.CategoryEntityTree, err errors.Error,
) {
	var categoryEntities []*entities.CategoryEntity

	if gormErr := r.db.Order("created asc").Find(&categoryEntities).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
		return
	}

	categoryEntityTree, err = r.categoryTreeBuilder.BuildCategoryEntityTree(categoryEntities)
	return
}

//GetCategory
func (r *categoryRepositoryImpl) GetCategory(categoryId *models.CategoryId) (categoryEntity *entities.CategoryEntity, err errors.Error) {
	categoryEntity = new(entities.CategoryEntity)

	if gormErr := r.db.First(categoryEntity, "id = ?", categoryId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewCategoryByIdNotFoundError(categoryId)
		} else {
			err = errors.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

//GetCategoryTree
func (r *categoryRepositoryImpl) GetCategoryTree(categoryId *models.CategoryId) (
	categoryEntityTree *entities.CategoryEntityTree, err errors.Error,
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
		err = errors.NewSystemErrorFromBuiltin(gormErr)
		return
	}

	categoryEntityTree, err = r.categoryTreeBuilder.BuildCategoryEntityTree(categoryEntities)
	return
}

//SaveCategory
func (r *categoryRepositoryImpl) SaveCategory(categoryEntity *entities.CategoryEntity) (err errors.Error) {
	if gormErr := r.db.Save(categoryEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//RemoveCategory
func (r *categoryRepositoryImpl) RemoveCategory(categoryEntity *entities.CategoryEntity) (err errors.Error) {
	if gormErr := r.db.Delete(categoryEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//GetCategoryXrefs
func (r *categoryRepositoryImpl) GetCategoryXrefs(categoryEntity *entities.CategoryEntity) (
	categoryXrefEntities []*entities.CategoryXrefEntity, err errors.Error,
) {
	gormErr := r.db.Where("category_id = ?", categoryEntity.Id).
		Order("created asc").
		Find(&categoryXrefEntities).Error

	if gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//GetCategoryObjectXrefs
func (r *categoryRepositoryImpl) GetCategoryObjectXrefs(categoryObject models.Object) (
	categoryXrefEntities []*entities.CategoryXrefEntity, err errors.Error,
) {
	gormErr := r.db.Where("object_type = ?", categoryObject.ObjectType()).
		Where("object_id = ?", categoryObject.ObjectId()).
		Order("created asc").
		Find(&categoryXrefEntities).Error

	if gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//GetCategoryXref
func (r *categoryRepositoryImpl) GetCategoryXref(
	categoryEntity *entities.CategoryEntity, categoryObject models.Object,
) (categoryXrefEntity *entities.CategoryXrefEntity, err errors.Error) {
	categoryXrefEntity = new(entities.CategoryXrefEntity)

	gormErr := r.db.Where("category_id = ?", categoryEntity.Id).
		Where("object_type = ?", categoryObject.ObjectType()).
		Where("object_id = ?", categoryObject.ObjectId()).
		First(categoryXrefEntity).Error

	if gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			categoryObjectType := string(categoryObject.ObjectType())
			err = errors.NewCategoryXrefNotFoundByReferenceError(
				categoryEntity.Id, categoryObjectType, categoryObject.ObjectId())
		} else {
			err = errors.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

//SaveCategoryXref
func (r *categoryRepositoryImpl) SaveCategoryXref(categoryXrefEntity *entities.CategoryXrefEntity) (err errors.Error) {
	if gormErr := r.db.Save(categoryXrefEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//RemoveCategoryXref
func (r *categoryRepositoryImpl) RemoveCategoryXref(categoryXrefEntity *entities.CategoryXrefEntity) (err errors.Error) {
	if gormErr := r.db.Delete(categoryXrefEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//ListObjectCategories
func (r *categoryRepositoryImpl) ListObjectCategories(
	categoryObject models.Object, categoryPaginationQuery *models.CategoryPaginationQuery,
) (paginationResult *models.PaginationResult, err errors.Error) {
	paginationTotal, categoryEntities := 0, make([]*entities.CategoryEntity, categoryPaginationQuery.Limit)

	db := r.db.Model(&categoryEntities).
		Joins(fmt.Sprintf(
			"inner join %s on %s.id = %s.category_id",
			entities.CategoryXrefTableName,
			entities.CategoryTableName,
			entities.CategoryXrefTableName,
		)).
		Where(fmt.Sprintf("%s.object_type = ?", entities.CategoryXrefTableName), categoryObject.ObjectType()).
		Where(fmt.Sprintf("%s.object_id = ?", entities.CategoryXrefTableName), categoryObject.ObjectId()).
		Order(fmt.Sprintf("%s.created desc", entities.CategoryTableName))

	err = r.dbPaginator.Paginate(db, categoryPaginationQuery.PaginationQuery, &categoryEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: categoryEntities}
	return
}
