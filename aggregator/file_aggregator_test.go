package aggregator

import (
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileAggregator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewFileAggregator", func(t *testing.T) {
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileAggregator, isFileAggregator := NewFileAggregator(fileModelFactory).(*fileAggregatorImpl)

		assert.True(t, isFileAggregator)
		assert.Equal(t, fileModelFactory, fileAggregator.fileModelFactory)
	})

	t.Run("AggregateFile", func(t *testing.T) {
		file := new(models.File)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFile().Return(file)

		fileAggregator := &fileAggregatorImpl{fileModelFactory: fileModelFactory}
		response := fileAggregator.AggregateFile(new(entities.FileEntity))

		assert.Equal(t, file, response)
	})

	t.Run("AggregateFiles", func(t *testing.T) {
		files := new(models.File)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFile().Return(files)

		fileAggregator := &fileAggregatorImpl{fileModelFactory: fileModelFactory}
		fileEntities := []*entities.FileEntity{new(entities.FileEntity)}
		response := fileAggregator.AggregateFiles(fileEntities)

		assert.IsType(t, []*models.File{}, response)
		assert.Equal(t, len(fileEntities), len(response))
	})

	t.Run("AggregatePaginationResult", func(t *testing.T) {
		file := new(models.File)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFile().Return(file)

		fileEntities := []*entities.FileEntity{entities.NewFileEntity()}
		fileAggregator := &fileAggregatorImpl{fileModelFactory: fileModelFactory}

		entityPaginationResult := &models.PaginationResult{Data: fileEntities}
		paginationResult := fileAggregator.AggregatePaginationResult(entityPaginationResult)

		assert.IsType(t, []*models.File{}, paginationResult.Data)
		assert.Equal(t, len(fileEntities), len(paginationResult.Data.([]*models.File)))
	})
}
