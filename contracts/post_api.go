package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostApi interface {
		ListPosts() (posts []*models.Post, err common.Error)
		GetPost(postId *models.PostId) (post *models.Post, err common.Error)
		CreatePost(postAuthor common.Subject, data *models.PostCreate) (post *models.Post, err common.Error)
		UpdatePost(postId *models.PostId, data *models.PostUpdate) (err common.Error)
		DeletePost(postId *models.PostId) (err common.Error)
	}
)
