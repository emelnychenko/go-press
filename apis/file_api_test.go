package apis

import (
	"bytes"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/events"
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

	t.Run("NewFileApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		fileEventFactory := mocks.NewMockFileEventFactory(ctrl)
		fileService := mocks.NewMockFileService(ctrl)
		fileAggregator := mocks.NewMockFileAggregator(ctrl)

		fileApi, isFileApi := NewFileApi(
			eventDispatcher, fileEventFactory, fileService, fileAggregator,
		).(*fileApiImpl)

		assert.True(t, isFileApi)
		assert.Equal(t, eventDispatcher, fileApi.eventDispatcher)
		assert.Equal(t, fileEventFactory, fileApi.fileEventFactory)
		assert.Equal(t, fileService, fileApi.fileService)
		assert.Equal(t, fileAggregator, fileApi.fileAggregator)
	})

	t.Run("ListFiles", func(t *testing.T) {
		paginationQuery := new(models.FilePaginationQuery)
		entityPaginationResult := new(models.PaginationResult)
		paginationResult := new(models.PaginationResult)

		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().ListFiles(paginationQuery).Return(entityPaginationResult, nil)

		fileAggregator := mocks.NewMockFileAggregator(ctrl)
		fileAggregator.EXPECT().AggregatePaginationResult(entityPaginationResult).Return(paginationResult)

		fileApi := &fileApiImpl{fileService: fileService, fileAggregator: fileAggregator}
		response, err := fileApi.ListFiles(paginationQuery)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListFiles:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		paginationQuery := new(models.FilePaginationQuery)

		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().ListFiles(paginationQuery).Return(nil, systemErr)

		fileApi := &fileApiImpl{fileService: fileService}
		response, err := fileApi.ListFiles(paginationQuery)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
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
		systemErr := errors.NewUnknownError()

		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}

		fileService.EXPECT().GetFile(fileId).Return(nil, systemErr)
		file, err := fileApi.GetFile(fileId)

		assert.Nil(t, file)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UploadFile", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)

		fileEvent := new(events.FileEvent)
		fileEventFactory := mocks.NewMockFileEventFactory(ctrl)
		fileEventFactory.EXPECT().CreateFileUploadedEvent(fileEntity).Return(fileEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(fileEvent)

		fileService := mocks.NewMockFileService(ctrl)
		fileAggregator := mocks.NewMockFileAggregator(ctrl)
		fileApi := &fileApiImpl{
			eventDispatcher:  eventDispatcher,
			fileEventFactory: fileEventFactory,
			fileService:      fileService,
			fileAggregator:   fileAggregator,
		}

		fileSource := bytes.NewBufferString("src")
		data := new(models.FileUpload)
		var reply *models.File
		fileService.EXPECT().UploadFile(fileSource, data).Return(fileEntity, nil)
		fileAggregator.EXPECT().AggregateFile(fileEntity).Return(reply)
		file, err := fileApi.UploadFile(fileSource, data)

		assert.Equal(t, reply, file)
		assert.Nil(t, err)
	})

	t.Run("UploadFile:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}
		data := new(models.FileUpload)
		fileSource := bytes.NewBufferString("src")

		fileService.EXPECT().UploadFile(fileSource, data).Return(nil, systemErr)
		file, err := fileApi.UploadFile(fileSource, data)

		assert.Nil(t, file)
		assert.Equal(t, systemErr, err)
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
		systemErr := errors.NewUnknownError()

		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{fileService: fileService}
		fileDestination := bytes.NewBufferString("test")

		fileService.EXPECT().GetFile(fileId).Return(nil, systemErr)
		err := fileApi.DownloadFile(fileId, func(file *models.File) io.Writer {
			return fileDestination
		})

		assert.Equal(t, systemErr, err)
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
		fileEntity := new(entities.FileEntity)

		fileEvent := new(events.FileEvent)
		fileEventFactory := mocks.NewMockFileEventFactory(ctrl)
		fileEventFactory.EXPECT().CreateFileUpdatedEvent(fileEntity).Return(fileEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(fileEvent)

		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{
			eventDispatcher:  eventDispatcher,
			fileEventFactory: fileEventFactory,
			fileService:      fileService,
		}
		data := new(models.FileUpdate)

		fileService.EXPECT().GetFile(fileId).Return(fileEntity, nil)
		fileService.EXPECT().UpdateFile(fileEntity, data).Return(nil)
		err := fileApi.UpdateFile(fileId, data)

		assert.Nil(t, err)
	})

	t.Run("UpdateFile:GetFileError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(fileId).Return(nil, systemErr)

		data := new(models.FileUpdate)
		fileApi := &fileApiImpl{fileService: fileService}
		assert.Equal(t, systemErr, fileApi.UpdateFile(fileId, data))
	})

	t.Run("UpdateFile:UpdateFileError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		fileEntity := new(entities.FileEntity)
		data := new(models.FileUpdate)

		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(fileId).Return(fileEntity, nil)
		fileService.EXPECT().UpdateFile(fileEntity, data).Return(systemErr)

		fileApi := &fileApiImpl{
			fileService: fileService,
		}
		err := fileApi.UpdateFile(fileId, data)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteFile", func(t *testing.T) {
		fileEntity := new(entities.FileEntity)

		fileEvent := new(events.FileEvent)
		fileEventFactory := mocks.NewMockFileEventFactory(ctrl)
		fileEventFactory.EXPECT().CreateFileDeletedEvent(fileEntity).Return(fileEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(fileEvent)

		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{
			eventDispatcher:  eventDispatcher,
			fileEventFactory: fileEventFactory,
			fileService:      fileService,
		}

		fileService.EXPECT().GetFile(fileId).Return(fileEntity, nil)
		fileService.EXPECT().DeleteFile(fileEntity).Return(nil)
		err := fileApi.DeleteFile(fileId)

		assert.Nil(t, err)
	})

	t.Run("DeleteFile:GetFileError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()

		fileService := mocks.NewMockFileService(ctrl)
		fileService.EXPECT().GetFile(fileId).Return(nil, systemErr)

		fileApi := &fileApiImpl{fileService: fileService}
		assert.Equal(t, systemErr, fileApi.DeleteFile(fileId))
	})

	t.Run("DeleteFile:DeleteFileError", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		fileEntity := new(entities.FileEntity)

		fileService := mocks.NewMockFileService(ctrl)
		fileApi := &fileApiImpl{
			fileService: fileService,
		}

		fileService.EXPECT().GetFile(fileId).Return(fileEntity, nil)
		fileService.EXPECT().DeleteFile(fileEntity).Return(systemErr)
		err := fileApi.DeleteFile(fileId)

		assert.Equal(t, systemErr, err)
	})
}
