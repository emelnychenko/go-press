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

	fileId := models.NewModelId()

	t.Run("NewFileService", func(t *testing.T) {
		fileRepository := mocks.NewMockFileRepository(ctrl)
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		fileService, isFileService := NewFileService(fileRepository, storageProvider).(*fileServiceImpl)

		assert.True(t, isFileService)
		assert.Equal(t, fileService.fileRepository, fileRepository)
		assert.Equal(t, fileService.storageProvider, storageProvider)
	})

	t.Run("ListFiles", func(t *testing.T) {
		fileRepository := mocks.NewMockFileRepository(ctrl)
		fileService := &fileServiceImpl{fileRepository: fileRepository}
		var replies []*entities.FileEntity

		fileRepository.EXPECT().ListFiles().Return(replies, nil)
		fileEntities, err := fileService.ListFiles()

		assert.Equal(t, replies, fileEntities)
		assert.Nil(t, err)
	})

	t.Run("GetFile", func(t *testing.T) {
		fileRepository := mocks.NewMockFileRepository(ctrl)
		fileService := &fileServiceImpl{fileRepository: fileRepository}
		var reply *entities.FileEntity

		fileRepository.EXPECT().GetFile(fileId).Return(reply, nil)
		fileEntity, err := fileService.GetFile(fileId)

		assert.Equal(t, reply, fileEntity)
		assert.Nil(t, err)
	})

	t.Run("UploadFile", func(t *testing.T) {
		fileRepository := mocks.NewMockFileRepository(ctrl)
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		fileService := &fileServiceImpl{fileRepository: fileRepository, storageProvider: storageProvider}
		fileSource := bytes.NewBufferString("test")
		data := &models.FileUpload{
			Name: "0",
			Size: 10,
			Type: "1",
		}
		filePath := ""

		fileRepository.EXPECT().SaveFile(gomock.Any()).Return(nil)
		storageProvider.EXPECT().UploadFile(gomock.Any(), fileSource).Return(filePath, nil)
		fileEntity, err := fileService.UploadFile(fileSource, data)

		assert.IsType(t, new(entities.FileEntity), fileEntity)
		assert.Nil(t, err)
		assert.Equal(t, data.Name, fileEntity.Name)
		assert.Equal(t, data.Size, fileEntity.Size)
		assert.Equal(t, data.Type, fileEntity.Type)
		assert.Equal(t, filePath, fileEntity.Path)
	})

	t.Run("UploadFile:Error", func(t *testing.T) {
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		fileService := &fileServiceImpl{storageProvider: storageProvider}
		fileSource := bytes.NewBufferString("test")
		data := new(models.FileUpload)

		storageProvider.EXPECT().UploadFile(gomock.Any(), fileSource).Return("", common.ServerError(""))
		fileEntity, err := fileService.UploadFile(fileSource, data)

		assert.IsType(t, new(entities.FileEntity), fileEntity)
		assert.Error(t, err)
	})

	t.Run("DownloadFile", func(t *testing.T) {
		storageProvider := mocks.NewMockStorageProvider(ctrl)
		fileService := &fileServiceImpl{storageProvider: storageProvider}
		fileEntity := new(entities.FileEntity)
		fileDestination := bytes.NewBuffer(nil)

		storageProvider.EXPECT().DownloadFile(fileEntity, fileDestination).Return(nil)
		err := fileService.DownloadFile(fileEntity, fileDestination)

		assert.Nil(t, err)
	})

	t.Run("UpdateFile", func(t *testing.T) {
		fileRepository := mocks.NewMockFileRepository(ctrl)
		fileService := &fileServiceImpl{fileRepository: fileRepository}
		data := &models.FileUpdate{
			Name: "0",
		}
		fileEntity := new(entities.FileEntity)

		fileRepository.EXPECT().SaveFile(fileEntity).Return(nil)
		assert.Nil(t, fileService.UpdateFile(fileEntity, data))

		assert.Equal(t, data.Name, fileEntity.Name)
		assert.NotNil(t, fileEntity.Updated)
	})

	t.Run("DeleteFile", func(t *testing.T) {
		fileRepository := mocks.NewMockFileRepository(ctrl)
		fileService := &fileServiceImpl{fileRepository: fileRepository}
		fileEntity := new(entities.FileEntity)

		fileRepository.EXPECT().RemoveFile(fileEntity).Return(nil)
		assert.Nil(t, fileService.DeleteFile(fileEntity))
	})
}
