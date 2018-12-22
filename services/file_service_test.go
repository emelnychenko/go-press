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
		fileEntityFactory := mocks.NewMockFileEntityFactory(ctrl)
		fileRepository := mocks.NewMockFileRepository(ctrl)
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		fileService, isFileService := NewFileService(fileEntityFactory, fileRepository, storageProvider).(*fileServiceImpl)

		assert.True(t, isFileService)
		assert.Equal(t, fileService.fileEntityFactory, fileEntityFactory)
		assert.Equal(t, fileService.fileRepository, fileRepository)
		assert.Equal(t, fileService.storageProvider, storageProvider)
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
		fileEntity := new(entities.FileEntity)
		fileEntityFactory := mocks.NewMockFileEntityFactory(ctrl)
		fileEntityFactory.EXPECT().CreateFileEntity().Return(fileEntity)

		fileRepository := mocks.NewMockFileRepository(ctrl)
		fileRepository.EXPECT().SaveFile(fileEntity).Return(nil)

		fileSource := bytes.NewBufferString("test")
		filePath := ""
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		storageProvider.EXPECT().UploadFile(fileEntity, fileSource).Return(filePath, nil)

		data := &models.FileUpload{
			Name: "0",
			Size: 10,
			Type: "1",
		}
		fileService := &fileServiceImpl{
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

	t.Run("UploadFile:Error", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)
		fileEntityFactory := mocks.NewMockFileEntityFactory(ctrl)
		fileEntityFactory.EXPECT().CreateFileEntity().Return(fileEntity)

		fileSource := bytes.NewBufferString("test")
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		storageProvider.EXPECT().UploadFile(fileEntity, fileSource).Return("", common.ServerError(""))

		data := new(models.FileUpload)
		fileService := &fileServiceImpl{fileEntityFactory: fileEntityFactory, storageProvider: storageProvider}
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
