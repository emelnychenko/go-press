package repositories

import (
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbPaginator := mocks.NewMockDbPaginator(ctrl)
	db, _ := gorm.Open(mocket.DriverName, "")
	fileRepository, isFileRepository := NewFileRepository(db, dbPaginator).(*fileRepositoryImpl)

	assert.True(t, isFileRepository)
	assert.Equal(t, db, fileRepository.db)
	assert.Equal(t, dbPaginator, fileRepository.dbPaginator)

	fileId := models.NewModelId()
	commonReply := []map[string]interface{}{{
		"id": fileId.String(),
	}}

	t.Run("ListFiles", func(t *testing.T) {
		filePaginationQuery := &models.FilePaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), filePaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(nil)

		paginationResult, err := fileRepository.ListFiles(filePaginationQuery)
		assert.IsType(t, []*entities.FileEntity{}, paginationResult.Data)
		assert.Nil(t, err)
	})

	t.Run("ListFiles:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		filePaginationQuery := &models.FilePaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), filePaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		fileEntities, err := fileRepository.ListFiles(filePaginationQuery)
		assert.Nil(t, fileEntities)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetFile", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		fileEntity, err := fileRepository.GetFile(fileId)
		assert.IsType(t, new(entities.FileEntity), fileEntity)
		assert.Nil(t, err)
	})

	t.Run("GetFile:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		fileEntity, err := fileRepository.GetFile(fileId)
		assert.NotNil(t, fileEntity)
		assert.Error(t, err)
	})

	t.Run("GetFile:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		fileEntity, err := fileRepository.GetFile(fileId)
		assert.NotNil(t, fileEntity)
		assert.Error(t, err, errors.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("SaveFile", func(t *testing.T) {
		mocket.Catcher.Reset()

		fileEntity := entities.NewFileEntity()
		assert.Nil(t, fileRepository.SaveFile(fileEntity))
	})

	t.Run("SaveFile:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		fileEntity := new(entities.FileEntity)
		assert.Error(t, fileRepository.SaveFile(fileEntity))
	})

	t.Run("RemoveFile", func(t *testing.T) {
		mocket.Catcher.Reset()

		fileEntity := new(entities.FileEntity)
		assert.Nil(t, fileRepository.RemoveFile(fileEntity))
	})

	t.Run("RemoveFile:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		fileEntity := new(entities.FileEntity)
		assert.Error(t, fileRepository.RemoveFile(fileEntity))
	})
}
