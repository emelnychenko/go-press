package services

import (
	"bytes"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewFileService", func(t *testing.T) {
		filePathStrategy := mocks.NewMockFilePathStrategy(ctrl)
		fileEntityFactory := mocks.NewMockFileEntityFactory(ctrl)
		fileRepository := mocks.NewMockFileRepository(ctrl)
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		fileService, isFileService := NewFileService(filePathStrategy, fileEntityFactory, fileRepository, storageProvider).(*fileServiceImpl)

		assert.True(t, isFileService)
		assert.Equal(t, filePathStrategy, fileService.filePathStrategy)
		assert.Equal(t, fileEntityFactory, fileService.fileEntityFactory)
		assert.Equal(t, fileRepository, fileService.fileRepository)
		assert.Equal(t, storageProvider, fileService.storageProvider)
	})

	t.Run("ListFiles", func(t *testing.T) {
		var fileEntities []*entities.FileEntity
		fileRepository := mocks.NewMockFileRepository(ctrl)
		fileRepository.EXPECT().ListFiles().Return(fileEntities, nil)

		fileService := &fileServiceImpl{fileRepository: fileRepository}
		response, err := fileService.ListFiles()

		assert.Equal(t, fileEntities, response)
		assert.Nil(t, err)
	})

	t.Run("GetFile", func(t *testing.T) {
		fileId := new(models.FileId)
		fileEntity := new(entities.FileEntity)
		fileRepository := mocks.NewMockFileRepository(ctrl)
		fileRepository.EXPECT().GetFile(fileId).Return(fileEntity, nil)

		fileService := &fileServiceImpl{fileRepository: fileRepository}
		response, err := fileService.GetFile(fileId)

		assert.Equal(t, fileEntity, response)
		assert.Nil(t, err)
	})

	t.Run("UploadFile", func(t *testing.T) {
		filePath := "path/to/file"
		fileEntity := new(entities.FileEntity)

		filePathStrategy := mocks.NewMockFilePathStrategy(ctrl)
		filePathStrategy.EXPECT().BuildPath(fileEntity).Return(filePath, nil)

		fileEntityFactory := mocks.NewMockFileEntityFactory(ctrl)
		fileEntityFactory.EXPECT().CreateFileEntity().Return(fileEntity)

		fileRepository := mocks.NewMockFileRepository(ctrl)
		fileRepository.EXPECT().SaveFile(fileEntity).Return(nil)

		fileSource := bytes.NewBufferString("test")
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		storageProvider.EXPECT().UploadFile(fileEntity, fileSource).Return(nil)

		data := &models.FileUpload{
			Name: "0",
			Size: 10,
			Type: "1",
		}
		fileService := &fileServiceImpl{
			filePathStrategy:  filePathStrategy,
			fileEntityFactory: fileEntityFactory,
			fileRepository:    fileRepository,
			storageProvider:   storageProvider,
		}
		response, err := fileService.UploadFile(fileSource, data)

		assert.Equal(t, fileEntity, response)
		assert.Nil(t, err)
		assert.Equal(t, data.Name, fileEntity.Name)
		assert.Equal(t, data.Size, fileEntity.Size)
		assert.Equal(t, data.Type, fileEntity.Type)
		assert.Equal(t, filePath, fileEntity.Path)
	})

	t.Run("UploadFile:BuildPathError", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)

		filePathStrategy := mocks.NewMockFilePathStrategy(ctrl)
		filePathStrategy.EXPECT().BuildPath(fileEntity).Return("", common.ServerError(""))

		fileEntityFactory := mocks.NewMockFileEntityFactory(ctrl)
		fileEntityFactory.EXPECT().CreateFileEntity().Return(fileEntity)

		fileSource := bytes.NewBufferString("test")

		data := new(models.FileUpload)
		fileService := &fileServiceImpl{
			filePathStrategy:  filePathStrategy,
			fileEntityFactory: fileEntityFactory,
		}

		response, err := fileService.UploadFile(fileSource, data)
		assert.Equal(t, fileEntity, response)
		assert.Error(t, err)
	})

	t.Run("UploadFile:StorageError", func(t *testing.T) {
		filePath := "path/to/file"
		fileEntity := new(entities.FileEntity)

		filePathStrategy := mocks.NewMockFilePathStrategy(ctrl)
		filePathStrategy.EXPECT().BuildPath(fileEntity).Return(filePath, nil)

		fileEntityFactory := mocks.NewMockFileEntityFactory(ctrl)
		fileEntityFactory.EXPECT().CreateFileEntity().Return(fileEntity)

		fileSource := bytes.NewBufferString("test")
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		storageProvider.EXPECT().UploadFile(fileEntity, fileSource).Return(common.ServerError(""))

		data := new(models.FileUpload)
		fileService := &fileServiceImpl{
			filePathStrategy:  filePathStrategy,
			fileEntityFactory: fileEntityFactory,
			storageProvider: storageProvider,
		}
		response, err := fileService.UploadFile(fileSource, data)

		assert.Equal(t, fileEntity, response)
		assert.Error(t, err)
	})

	t.Run("DownloadFile", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileDestination := bytes.NewBuffer(nil)
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		storageProvider.EXPECT().DownloadFile(fileEntity, fileDestination).Return(nil)

		fileService := &fileServiceImpl{storageProvider: storageProvider}
		assert.Nil(t, fileService.DownloadFile(fileEntity, fileDestination))
	})

	t.Run("UpdateFile", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileRepository := mocks.NewMockFileRepository(ctrl)
		fileRepository.EXPECT().SaveFile(fileEntity).Return(nil)

		data := &models.FileUpdate{Name: "0"}
		fileService := &fileServiceImpl{fileRepository: fileRepository}

		assert.Nil(t, fileService.UpdateFile(fileEntity, data))
		assert.Equal(t, data.Name, fileEntity.Name)
		assert.NotNil(t, fileEntity.Updated)
	})

	t.Run("DeleteFile", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileRepository := mocks.NewMockFileRepository(ctrl)
		fileRepository.EXPECT().RemoveFile(fileEntity).Return(nil)

		fileService := &fileServiceImpl{fileRepository: fileRepository}
		assert.Nil(t, fileService.DeleteFile(fileEntity))
	})
}
