package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostService interface {
		ListPosts(
			postPaginationQuery *models.PostPaginationQuery,
		) (paginationResult *models.PaginationResult, err common.Error)
		GetScheduledPosts() (postEntities []*entities.PostEntity, err common.Error)
		GetPost(postId *models.PostId) (postEntity *entities.PostEntity, err common.Error)
		CreatePost(postAuthor common.Subject, data *models.PostCreate) (postEntity *entities.PostEntity, err common.Error)
		UpdatePost(postEntity *entities.PostEntity, data *models.PostUpdate) (err common.Error)
		DeletePost(postEntity *entities.PostEntity) (err common.Error)
	}
)
