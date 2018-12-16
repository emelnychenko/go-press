package apis

import (
	"bytes"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestFileApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileId := models.NewModelId()
	testErr := common.ServerError("err0")

	t.Run("NewFileApi", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileAggregator := mocks.NewMockFileAggregator(ctrl)
		fileApi, isFileApi := NewFileApi(fileService, fileAggregator).(*fileApiImpl)

		assert.True(t, isFileApi)
		assert.Equal(t, fileService, fileApi.fileService)
		assert.Equal(t, fileAggregator, fileApi.fileAggregator)
	})

	t.Run("ListFiles", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileAggregator := mocks.NewMockFileAggregator(ctrl)
		fileApi := &fileApiImpl{fileService: fileService, fileAggregator: fileAggregator}
		var fileEntities []*entities.FileEntity
		var replies []*models.File

		fileService.EXPECT().ListFiles().Return(fileEntities, nil)
		fileAggregator.EXPECT().AggregateFiles(fileEntities).Return(replies)
		files, err := fileApi.ListFiles()

		assert.Equal(t, replies, files)
		assert.Nil(t, err)
	})

	t.Run("ListFiles:Error", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}

		fileService.EXPECT().ListFiles().Return(nil, testErr)
		files, err := fileApi.ListFiles()

		assert.Nil(t, files)
		assert.Equal(t, testErr, err)
	})

	t.Run("GetFile", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileAggregator := mocks.NewMockFileAggregator(ctrl)
		fileApi := &fileApiImpl{fileService: fileService, fileAggregator: fileAggregator}
		var fileEntity *entities.FileEntity
		var reply *models.File

		fileService.EXPECT().GetFile(fileId).Return(fileEntity, nil)
		fileAggregator.EXPECT().AggregateFile(fileEntity).Return(reply)
		file, err := fileApi.GetFile(fileId)

		assert.Equal(t, reply, file)
		assert.Nil(t, err)
	})

	t.Run("GetFile:Error", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}

		fileService.EXPECT().GetFile(fileId).Return(nil, testErr)
		file, err := fileApi.GetFile(fileId)

		assert.Nil(t, file)
		assert.Equal(t, testErr, err)
	})

	t.Run("UploadFile", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileAggregator := mocks.NewMockFileAggregator(ctrl)
		fileApi := &fileApiImpl{fileService: fileService, fileAggregator: fileAggregator}
		fileSource := bytes.NewBufferString("src")
		data := new(models.FileUpload)
		var fileEntity *entities.FileEntity
		var reply *models.File

		fileService.EXPECT().UploadFile(fileSource, data).Return(fileEntity, nil)
		fileAggregator.EXPECT().AggregateFile(fileEntity).Return(reply)
		file, err := fileApi.UploadFile(fileSource, data)

		assert.Equal(t, reply, file)
		assert.Nil(t, err)
	})

	t.Run("UploadFile:Error", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}
		data := new(models.FileUpload)
		fileSource := bytes.NewBufferString("src")

		fileService.EXPECT().UploadFile(fileSource, data).Return(nil, testErr)
		file, err := fileApi.UploadFile(fileSource, data)

		assert.Nil(t, file)
		assert.Equal(t, testErr, err)
	})

	t.Run("DownloadFile", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileAggregator := mocks.NewMockFileAggregator(ctrl)
		fileApi := &fileApiImpl{fileService: fileService, fileAggregator: fileAggregator}
		fileDestination := bytes.NewBufferString("src")
		var fileEntity *entities.FileEntity
		var reply *models.File

		fileService.EXPECT().GetFile(fileId).Return(fileEntity, nil)
		fileAggregator.EXPECT().AggregateFile(fileEntity).Return(reply)
		fileService.EXPECT().DownloadFile(fileEntity, fileDestination).Return(nil)
		err := fileApi.DownloadFile(fileId, func(file *models.File) io.Writer {
			assert.Equal(t, reply, file)
			return fileDestination
		})

		assert.Nil(t, err)
	})

	t.Run("DownloadFile:Error", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}
		fileDestination := bytes.NewBufferString("test")

		fileService.EXPECT().GetFile(fileId).Return(nil, testErr)
		err := fileApi.DownloadFile(fileId, func(file *models.File) io.Writer {
			return fileDestination
		})

		assert.Equal(t, testErr, err)
	})

	t.Run("DownloadFile:NoDestination", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileAggregator := mocks.NewMockFileAggregator(ctrl)
		fileApi := &fileApiImpl{fileService: fileService, fileAggregator: fileAggregator}
		var fileEntity *entities.FileEntity
		var reply *models.File

		fileService.EXPECT().GetFile(fileId).Return(fileEntity, nil)
		fileAggregator.EXPECT().AggregateFile(fileEntity).Return(reply)
		err := fileApi.DownloadFile(fileId, func(file *models.File) io.Writer {
			return nil
		})

		assert.Nil(t, err)
	})

	t.Run("UpdateFile", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}
		data := new(models.FileUpdate)
		var fileEntity *entities.FileEntity

		fileService.EXPECT().GetFile(fileId).Return(fileEntity, nil)
		fileService.EXPECT().UpdateFile(fileEntity, data).Return(nil)
		err := fileApi.UpdateFile(fileId, data)

		assert.Nil(t, err)
	})

	t.Run("UpdateFile:Error", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}
		data := new(models.FileUpdate)

		fileService.EXPECT().GetFile(fileId).Return(nil, testErr)
		err := fileApi.UpdateFile(fileId, data)

		assert.Equal(t, testErr, err)
	})

	t.Run("DeleteFile", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}
		var fileEntity *entities.FileEntity

		fileService.EXPECT().GetFile(fileId).Return(fileEntity, nil)
		fileService.EXPECT().DeleteFile(fileEntity).Return(nil)
		err := fileApi.DeleteFile(fileId)

		assert.Nil(t, err)
	})

	t.Run("DeleteFile:Error", func(t *testing.T) {
		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}

		fileService.EXPECT().GetFile(fileId).Return(nil, testErr)
		err := fileApi.DeleteFile(fileId)

		assert.Equal(t, testErr, err)
	})
}
