package contracts

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/models"
)

type (
	PostApi interface {
		ListPosts() ([]*models.Post, common.Error)
		GetPost(postId models.PostId) (*models.Post, common.Error)
	}
)
