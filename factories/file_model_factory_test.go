package factories

import (
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileModelFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewFileModelFactory", func(t *testing.T) {
		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		fileModelFactory, isFileModelFactory := NewFileModelFactory(paginationModelFactory).(*fileModelFactoryImpl)

		assert.True(t, isFileModelFactory)
		assert.Equal(t, paginationModelFactory, fileModelFactory.paginationModelFactory)
	})

	t.Run("CreateFilePaginationQuery", func(t *testing.T) {
		paginationQuery := new(models.PaginationQuery)

		paginationModelFactory := mocks.NewMockPaginationModelFactory(ctrl)
		paginationModelFactory.EXPECT().CreatePaginationQuery().Return(paginationQuery)

		fileModelFactory := &fileModelFactoryImpl{paginationModelFactory: paginationModelFactory}
		filePaginationQuery := fileModelFactory.CreateFilePaginationQuery()

		assert.Equal(t, paginationQuery, filePaginationQuery.PaginationQuery)
	})

	t.Run("CreateFile", func(t *testing.T) {
		fileModelFactory := new(fileModelFactoryImpl)
		assert.NotNil(t, fileModelFactory.CreateFile())
	})

	t.Run("CreateFileUpload", func(t *testing.T) {
		fileModelFactory := new(fileModelFactoryImpl)
		assert.NotNil(t, fileModelFactory.CreateFileUpload())
	})

	t.Run("CreateFileUpdate", func(t *testing.T) {
		fileModelFactory := new(fileModelFactoryImpl)
		assert.NotNil(t, fileModelFactory.CreateFileUpdate())
	})
}
