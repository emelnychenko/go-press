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

func (c *fileRepositoryImpl) ListFiles() (fileEntities []*entities.FileEntity, err common.Error) {
	if err := c.db.Find(&fileEntities).Error; nil != err {
		return nil, common.NewServerError(err)
	}

	return
}

func (c *fileRepositoryImpl) GetFile(fileId *models.FileId) (fileEntity *entities.FileEntity, err common.Error) {
	fileEntity = new(entities.FileEntity)

	if err := c.db.First(fileEntity, "id = ?", fileId).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.NewFileNotFoundError(fileId.String())
		}
		return nil, common.NewServerError(err)
	}

	return
}

func (c *fileRepositoryImpl) SaveFile(fileEntity *entities.FileEntity) (err common.Error) {
	if err := c.db.Save(fileEntity).Error; err != nil {
		return common.NewServerError(err)
	}

	return
}

func (c *fileRepositoryImpl) RemoveFile(fileEntity *entities.FileEntity) (err common.Error) {
	if err := c.db.Delete(fileEntity).Error; err != nil {
		return common.NewServerError(err)
	}

	return
}
