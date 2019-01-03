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
		db          *gorm.DB
		dbPaginator contracts.DbPaginator
	}
)

func NewCategoryRepository(db *gorm.DB, dbPaginator contracts.DbPaginator) contracts.CategoryRepository {
	return &categoryRepositoryImpl{db, dbPaginator}
}

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

func (r *categoryRepositoryImpl) SaveCategory(categoryEntity *entities.CategoryEntity) (err common.Error) {
	if gormErr := r.db.Save(categoryEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

func (r *categoryRepositoryImpl) RemoveCategory(categoryEntity *entities.CategoryEntity) (err common.Error) {
	if gormErr := r.db.Delete(categoryEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
