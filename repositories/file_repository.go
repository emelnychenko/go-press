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
	fileRepositoryImpl struct {
		db          *gorm.DB
		dbPaginator contracts.DbPaginator
	}
)

func NewFileRepository(db *gorm.DB, dbPaginator contracts.DbPaginator) contracts.FileRepository {
	return &fileRepositoryImpl{db, dbPaginator}
}

func (r *fileRepositoryImpl) ListFiles(
	filePaginationQuery *models.FilePaginationQuery,
) (paginationResult *models.PaginationResult, err common.Error) {
	paginationTotal, fileEntities := 0, make([]*entities.FileEntity, filePaginationQuery.Limit)
	db := r.db.Model(&fileEntities).Order("created desc")
	err = r.dbPaginator.Paginate(db, filePaginationQuery.PaginationQuery, &fileEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: fileEntities}
	return
}

func (r *fileRepositoryImpl) GetFile(fileId *models.FileId) (fileEntity *entities.FileEntity, err common.Error) {
	fileEntity = new(entities.FileEntity)

	if gormErr := r.db.First(fileEntity, "id = ?", fileId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewFileByIdNotFoundError(fileId)
		} else {
			err = common.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

func (r *fileRepositoryImpl) SaveFile(fileEntity *entities.FileEntity) (err common.Error) {
	if gormErr := r.db.Save(fileEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

func (r *fileRepositoryImpl) RemoveFile(fileEntity *entities.FileEntity) (err common.Error) {
	if gormErr := r.db.Delete(fileEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
