package contracts

import (
	"github.com/emelnychenko/go-press/models"
)

type (
	PostModelFactory interface {
		CreatePost() *models.Post
		CreatePostCreate() *models.PostCreate
		CreatePostUpdate() *models.PostUpdate
	}
)
