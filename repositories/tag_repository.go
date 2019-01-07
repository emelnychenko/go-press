package repositories

import (
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

//NewTagRepository
func NewTagRepository(db *gorm.DB, dbPaginator contracts.DbPaginator) contracts.TagRepository {
	return &tagRepositoryImpl{db, dbPaginator}
}

//ListTags
func (r *tagRepositoryImpl) ListTags(
	tagPaginationQuery *models.TagPaginationQuery,
) (paginationResult *models.PaginationResult, err errors.Error) {
	paginationTotal, tagEntities := 0, make([]*entities.TagEntity, tagPaginationQuery.Limit)
	db := r.db.Model(&tagEntities).Order("created desc")
	err = r.dbPaginator.Paginate(db, tagPaginationQuery.PaginationQuery, &tagEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: tagEntities}
	return
}

//GetTag
func (r *tagRepositoryImpl) GetTag(tagId *models.TagId) (tagEntity *entities.TagEntity, err errors.Error) {
	tagEntity = new(entities.TagEntity)

	if gormErr := r.db.First(tagEntity, "id = ?", tagId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewTagByIdNotFoundError(tagId)
		} else {
			err = errors.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

//SaveTag
func (r *tagRepositoryImpl) SaveTag(tagEntity *entities.TagEntity) (err errors.Error) {
	if gormErr := r.db.Save(tagEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//RemoveTag
func (r *tagRepositoryImpl) RemoveTag(tagEntity *entities.TagEntity) (err errors.Error) {
	if gormErr := r.db.Delete(tagEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//GetTagXrefs 
func (r *tagRepositoryImpl) GetTagXrefs(tagEntity *entities.TagEntity) (
	tagXrefEntities []*entities.TagXrefEntity, err errors.Error,
) {
	gormErr := r.db.Where("tag_id = ?", tagEntity.Id).
		Order("created asc").
		Find(&tagXrefEntities).Error

	if gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//GetTagObjectXrefs 
func (r *tagRepositoryImpl) GetTagObjectXrefs(tagObject models.Object) (
	tagXrefEntities []*entities.TagXrefEntity, err errors.Error,
) {
	gormErr := r.db.Where("object_type = ?", tagObject.ObjectType()).
		Where("object_id = ?", tagObject.ObjectId()).
		Order("created asc").
		Find(&tagXrefEntities).Error

	if gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//GetTagXref
func (r *tagRepositoryImpl) GetTagXref(
	tagEntity *entities.TagEntity, tagObject models.Object,
) (tagXrefEntity *entities.TagXrefEntity, err errors.Error) {
	tagXrefEntity = new(entities.TagXrefEntity)

	gormErr := r.db.Where("tag_id = ?", tagEntity.Id).
		Where("object_type = ?", tagObject.ObjectType()).
		Where("object_id = ?", tagObject.ObjectId()).
		First(tagXrefEntity).Error

	if gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			tagObjectType := string(tagObject.ObjectType())
			err = errors.NewTagXrefNotFoundByReferenceError(
				tagEntity.Id, tagObjectType, tagObject.ObjectId())
		} else {
			err = errors.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

//SaveTagXref
func (r *tagRepositoryImpl) SaveTagXref(tagXrefEntity *entities.TagXrefEntity) (err errors.Error) {
	if gormErr := r.db.Save(tagXrefEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

//RemoveTagXref
func (r *tagRepositoryImpl) RemoveTagXref(tagXrefEntity *entities.TagXrefEntity) (err errors.Error) {
	if gormErr := r.db.Delete(tagXrefEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
