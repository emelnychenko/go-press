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
		db *gorm.DB
	}
)

func NewFileRepository(db *gorm.DB) (fileRepository contracts.FileRepository) {
	return &fileRepositoryImpl{db}
}

func (r *fileRepositoryImpl) ListFiles() (fileEntities []*entities.FileEntity, err common.Error) {
	if gormErr := r.db.Find(&fileEntities).Error; nil != gormErr {
		err = common.NewSystemError(gormErr)
	}

	return
}

func (r *fileRepositoryImpl) GetFile(fileId *models.FileId) (fileEntity *entities.FileEntity, err common.Error) {
	fileEntity = new(entities.FileEntity)

	if gormErr := r.db.First(fileEntity, "id = ?", fileId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewFileByIdNotFoundError(fileId)
		} else {
			err = common.NewSystemError(gormErr)
		}
	}

	return
}

func (r *fileRepositoryImpl) SaveFile(fileEntity *entities.FileEntity) (err common.Error) {
	if gormErr := r.db.Save(fileEntity).Error; gormErr != nil {
		err = common.NewSystemError(gormErr)
	}

	return
}

func (r *fileRepositoryImpl) RemoveFile(fileEntity *entities.FileEntity) (err common.Error) {
	if gormErr := r.db.Delete(fileEntity).Error; gormErr != nil {
		err = common.NewSystemError(gormErr)
	}

	return
}
