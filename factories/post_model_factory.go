package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	postModelFactoryImpl struct {
	}
)

func NewPostModelFactory() contracts.PostModelFactory {
	return new(postModelFactoryImpl)
}

func (*postModelFactoryImpl) CreatePost() *models.Post {
	return new(models.Post)
}

func (*postModelFactoryImpl) CreatePostCreate() *models.PostCreate {
	return new(models.PostCreate)
}

func (*postModelFactoryImpl) CreatePostUpdate() *models.PostUpdate {
	return new(models.PostUpdate)
}
