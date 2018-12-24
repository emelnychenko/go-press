package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/models"
	"io"
)

type (
	fileApiImpl struct {
		eventDispatcher  contracts.EventDispatcher
		fileEventFactory contracts.FileEventFactory
		fileService      contracts.FileService
		fileAggregator   contracts.FileAggregator
	}
)

func NewFileApi(
	eventDispatcher contracts.EventDispatcher,
	fileEventFactory contracts.FileEventFactory,
	fileService contracts.FileService,
	fileAggregator contracts.FileAggregator,
) (fileApi contracts.FileApi) {
	return &fileApiImpl{
		eventDispatcher,
		fileEventFactory,
		fileService,
		fileAggregator,
	}
}

func (a *fileApiImpl) ListFiles() (files []*models.File, err common.Error) {
	fileEntities, err := a.fileService.ListFiles()

	if nil != err {
		return
	}

	files = a.fileAggregator.AggregateFiles(fileEntities)
	return
}

func (a *fileApiImpl) GetFile(fileId *models.FileId) (file *models.File, err common.Error) {
	fileEntity, err := a.fileService.GetFile(fileId)

	if nil != err {
		return
	}

	file = a.fileAggregator.AggregateFile(fileEntity)
	return
}

func (a *fileApiImpl) UploadFile(fileSource io.Reader, data *models.FileUpload) (file *models.File, err common.Error) {
	fileEntity, err := a.fileService.UploadFile(fileSource, data)

	if nil != err {
		return
	}

	fileEvent := a.fileEventFactory.CreateFileUploadedEvent(fileEntity)
	a.eventDispatcher.Dispatch(fileEvent)

	file = a.fileAggregator.AggregateFile(fileEntity)
	return
}

func (a *fileApiImpl) DownloadFile(fileId *models.FileId, getFileDestination contracts.PrepareFileDestination) (err common.Error) {
	fileService := a.fileService
	fileEntity, err := fileService.GetFile(fileId)

	if nil != err {
		return
	}

	file := a.fileAggregator.AggregateFile(fileEntity)
	fileDestination := getFileDestination(file)

	if nil == fileDestination {
		return
	}

	err = a.fileService.DownloadFile(fileEntity, fileDestination)
	return
}

func (a *fileApiImpl) UpdateFile(fileId *models.FileId, data *models.FileUpdate) (err common.Error) {
	fileService := a.fileService
	fileEntity, err := fileService.GetFile(fileId)

	if nil != err {
		return
	}

	err = fileService.UpdateFile(fileEntity, data)

	if nil != err {
		return
	}

	fileEvent := a.fileEventFactory.CreateFileUpdatedEvent(fileEntity)
	a.eventDispatcher.Dispatch(fileEvent)

	return
}

func (a *fileApiImpl) DeleteFile(fileId *models.FileId) (err common.Error) {
	fileService := a.fileService
	fileEntity, err := fileService.GetFile(fileId)

	if nil != err {
		return
	}

	err = fileService.DeleteFile(fileEntity)

	if nil != err {
		return
	}

	fileEvent := a.fileEventFactory.CreateFileDeletedEvent(fileEntity)
	a.eventDispatcher.Dispatch(fileEvent)

	return
}
