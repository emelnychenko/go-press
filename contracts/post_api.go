package contracts

import (
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostApi interface {
		ListPosts(postPaginationQuery *models.PostPaginationQuery) (*models.PaginationResult, errors.Error)
		GetPost(postId *models.PostId) (post *models.Post, err errors.Error)
		CreatePost(postAuthor models.Subject, data *models.PostCreate) (post *models.Post, err errors.Error)
		UpdatePost(postId *models.PostId, data *models.PostUpdate) (err errors.Error)
		DeletePost(postId *models.PostId) (err errors.Error)
	}
)
