package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostService interface {
		ListPosts(postPaginationQuery *models.PostPaginationQuery) (*models.PaginationResult, errors.Error)
		GetScheduledPosts() (postEntities []*entities.PostEntity, err errors.Error)
		GetPost(postId *models.PostId) (postEntity *entities.PostEntity, err errors.Error)
		CreatePost(postAuthor models.Subject, data *models.PostCreate) (postEntity *entities.PostEntity, err errors.Error)
		UpdatePost(postEntity *entities.PostEntity, data *models.PostUpdate) (err errors.Error)
		DeletePost(postEntity *entities.PostEntity) (err errors.Error)
	}
)
