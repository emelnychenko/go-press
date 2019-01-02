package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostRepository interface {
		ListPosts(
			postPaginationQuery *models.PostPaginationQuery,
		) (paginationResult *models.PaginationResult, err common.Error)
		GetScheduledPosts() (postEntities []*entities.PostEntity, err common.Error)
		GetPost(postId *models.PostId) (postEntity *entities.PostEntity, err common.Error)
		SavePost(postEntity *entities.PostEntity) (err common.Error)
		RemovePost(postEntity *entities.PostEntity) (err common.Error)
	}
)
