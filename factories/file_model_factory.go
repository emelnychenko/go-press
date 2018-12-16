package factories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
)

type (
	fileModelFactoryImpl struct {
	}
)

func NewFileModelFactory() contracts.FileModelFactory {
	return &fileModelFactoryImpl{}
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
