package contracts

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type (
	FileAggregator interface {
		AggregateFile(fileEntity *entities.FileEntity) (file *models.File)
		AggregateFiles(fileEntities []*entities.FileEntity) (files []*models.File)
		AggregatePaginationResult(
			entityPaginationResult *models.PaginationResult,
		) (paginationResult *models.PaginationResult)
	}
)
