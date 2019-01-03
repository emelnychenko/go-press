package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	commentModelFactoryImpl struct {
		paginationModelFactory contracts.PaginationModelFactory
	}
)

func NewCommentModelFactory(paginationModelFactory contracts.PaginationModelFactory) contracts.CommentModelFactory {
	return &commentModelFactoryImpl{paginationModelFactory}
}

func (f *commentModelFactoryImpl) CreateCommentPaginationQuery() *models.CommentPaginationQuery {
	paginationQuery := f.paginationModelFactory.CreatePaginationQuery()
	return &models.CommentPaginationQuery{PaginationQuery: paginationQuery}
}

func (*commentModelFactoryImpl) CreateComment() *models.Comment {
	return new(models.Comment)
}

func (*commentModelFactoryImpl) CreateCommentCreate() *models.CommentCreate {
	return new(models.CommentCreate)
}

func (*commentModelFactoryImpl) CreateCommentUpdate() *models.CommentUpdate {
	return new(models.CommentUpdate)
}
