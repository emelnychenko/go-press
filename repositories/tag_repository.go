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
	tagRepositoryImpl struct {
		db          *gorm.DB
		dbPaginator contracts.DbPaginator
	}
)

func NewTagRepository(db *gorm.DB, dbPaginator contracts.DbPaginator) contracts.TagRepository {
	return &tagRepositoryImpl{db, dbPaginator}
}

func (r *tagRepositoryImpl) ListTags(
	tagPaginationQuery *models.TagPaginationQuery,
) (paginationResult *models.PaginationResult, err common.Error) {
	paginationTotal, tagEntities := 0, make([]*entities.TagEntity, tagPaginationQuery.Limit)
	db := r.db.Model(&tagEntities).Order("created desc")
	err = r.dbPaginator.Paginate(db, tagPaginationQuery.PaginationQuery, &tagEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: tagEntities}
	return
}

func (r *tagRepositoryImpl) GetTag(tagId *models.TagId) (tagEntity *entities.TagEntity, err common.Error) {
	tagEntity = new(entities.TagEntity)

	if gormErr := r.db.First(tagEntity, "id = ?", tagId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewTagByIdNotFoundError(tagId)
		} else {
			err = common.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

func (r *tagRepositoryImpl) SaveTag(tagEntity *entities.TagEntity) (err common.Error) {
	if gormErr := r.db.Save(tagEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

func (r *tagRepositoryImpl) RemoveTag(tagEntity *entities.TagEntity) (err common.Error) {
	if gormErr := r.db.Delete(tagEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
