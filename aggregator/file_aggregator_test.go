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
		reply := new(models.File)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFile().Return(reply)

		fileAggregator := &fileAggregatorImpl{fileModelFactory: fileModelFactory}

		file := fileAggregator.AggregateFile(new(entities.FileEntity))
		assert.Equal(t, reply, file)
	})

	t.Run("AggregateCollection", func(t *testing.T) {
		reply := new(models.File)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFile().Return(reply)

		fileAggregator := &fileAggregatorImpl{fileModelFactory: fileModelFactory}
		fileEntities := []*entities.FileEntity{entities.NewFileEntity()}

		files := fileAggregator.AggregateFiles(fileEntities)
		assert.IsType(t, []*models.File{}, files)
		assert.Equal(t, len(fileEntities), len(files))
	})
}
