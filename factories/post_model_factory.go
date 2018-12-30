package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	postModelFactoryImpl struct {
		paginationModelFactory contracts.PaginationModelFactory
	}
)

func NewPostModelFactory(paginationModelFactory contracts.PaginationModelFactory) contracts.PostModelFactory {
	return &postModelFactoryImpl{paginationModelFactory}
}

func (f *postModelFactoryImpl) CreatePostPaginationQuery() *models.PostPaginationQuery {
	paginationQuery := f.paginationModelFactory.CreatePaginationQuery()
	return &models.PostPaginationQuery{PaginationQuery: paginationQuery}
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
