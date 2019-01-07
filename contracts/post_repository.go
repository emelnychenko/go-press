package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostRepository interface {
		ListPosts(postPaginationQuery *models.PostPaginationQuery) (*models.PaginationResult, errors.Error)
		GetScheduledPosts() (postEntities []*entities.PostEntity, err errors.Error)
		GetPost(postId *models.PostId) (postEntity *entities.PostEntity, err errors.Error)
		SavePost(postEntity *entities.PostEntity) (err errors.Error)
		RemovePost(postEntity *entities.PostEntity) (err errors.Error)
	}
)
