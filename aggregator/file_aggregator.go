package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type fileAggregatorImpl struct {
	fileModelFactory contracts.FileModelFactory
}

func NewFileAggregator(fileModelFactory contracts.FileModelFactory) contracts.FileAggregator {
	return &fileAggregatorImpl{fileModelFactory}
}

func (a *fileAggregatorImpl) AggregateFile(fileEntity *entities.FileEntity) (file *models.File) {
	file = a.fileModelFactory.CreateFile()
	file.Id = fileEntity.Id
	file.Name = fileEntity.Name
	file.Size = fileEntity.Size
	file.Type = fileEntity.Type
	file.Path = fileEntity.Path
	file.Created = fileEntity.Created

	return
}

func (a *fileAggregatorImpl) AggregateFiles(fileEntities []*entities.FileEntity) (files []*models.File) {
	files = make([]*models.File, len(fileEntities))

	for k, postEntity := range fileEntities {
		files[k] = a.AggregateFile(postEntity)
	}

	return
}
