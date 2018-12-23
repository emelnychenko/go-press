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
	postRepositoryImpl struct {
		db *gorm.DB
	}
)

func NewPostRepository(db *gorm.DB) (postRepository contracts.PostRepository) {
	return &postRepositoryImpl{db}
}

func (c *postRepositoryImpl) ListPosts() (postEntities []*entities.PostEntity, err common.Error) {
	if gormErr := c.db.Find(&postEntities).Error; nil != gormErr {
		err = common.NewSystemError(gormErr)
	}

	return
}

func (c *postRepositoryImpl) GetPost(postId *models.PostId) (postEntity *entities.PostEntity, err common.Error) {
	postEntity = new(entities.PostEntity)

	if gormErr := c.db.First(postEntity, "id = ?", postId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewPostByIdNotFoundError(postId)
		} else {
			err = common.NewSystemError(gormErr)
		}
	}

	return
}

func (c *postRepositoryImpl) SavePost(postEntity *entities.PostEntity) (err common.Error) {
	if gormErr := c.db.Save(postEntity).Error; gormErr != nil {
		err = common.NewSystemError(gormErr)
	}

	return
}

func (c *postRepositoryImpl) RemovePost(postEntity *entities.PostEntity) (err common.Error) {
	if gormErr := c.db.Delete(postEntity).Error; gormErr != nil {
		err = common.NewSystemError(gormErr)
	}

	return
}
