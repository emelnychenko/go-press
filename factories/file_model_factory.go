package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	fileModelFactoryImpl struct {
		paginationModelFactory contracts.PaginationModelFactory
	}
)

func NewFileModelFactory(paginationModelFactory contracts.PaginationModelFactory) contracts.FileModelFactory {
	return &fileModelFactoryImpl{paginationModelFactory}
}

func (f *fileModelFactoryImpl) CreateFilePaginationQuery() *models.FilePaginationQuery {
	paginationQuery := f.paginationModelFactory.CreatePaginationQuery()
	return &models.FilePaginationQuery{PaginationQuery: paginationQuery}
}

func (*fileModelFactoryImpl) CreateFile() *models.File {
	return new(models.File)
}

func (*fileModelFactoryImpl) CreateFileUpload() *models.FileUpload {
	return new(models.FileUpload)
}

func (*fileModelFactoryImpl) CreateFileUpdate() *models.FileUpdate {
	return new(models.FileUpdate)
}
