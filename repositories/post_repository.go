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
	if err := c.db.Find(&postEntities).Error; nil != err {
		return nil, common.NewServerError(err)
	}

	return
}

func (c *postRepositoryImpl) GetPost(postId *models.PostId) (postEntity *entities.PostEntity, err common.Error) {
	postEntity = &entities.PostEntity{}

	if err := c.db.First(postEntity, "id = ?", postId).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.NewPostNotFoundError(postId.String())
		}
		return nil, common.NewServerError(err)
	}

	return
}

func (c *postRepositoryImpl) SavePost(postEntity *entities.PostEntity) (err common.Error) {
	if err := c.db.Save(postEntity).Error; err != nil {
		return common.NewServerError(err)
	}

	return
}

func (c *postRepositoryImpl) RemovePost(postEntity *entities.PostEntity) (err common.Error) {
	if err := c.db.Delete(postEntity).Error; err != nil {
		return common.NewServerError(err)
	}

	return
}
