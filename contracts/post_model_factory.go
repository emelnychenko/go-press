package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	PostModelFactory interface {
		CreatePostPaginationQuery() *models.PostPaginationQuery
		CreatePost() *models.Post
		CreatePostCreate() *models.PostCreate
		CreatePostUpdate() *models.PostUpdate
	}
)
