package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"io"
	"time"
)

type (
	fileServiceImpl struct {
		filePathStrategy  contracts.FilePathStrategy
		fileEntityFactory contracts.FileEntityFactory
		fileRepository    contracts.FileRepository
		storageProvider   contracts.StorageProvider
	}
)

func NewFileService(
	filePathStrategy contracts.FilePathStrategy,
	fileEntityFactory contracts.FileEntityFactory,
	fileRepository contracts.FileRepository,
	storageProvider contracts.StorageProvider,
) (fileService contracts.FileService) {
	return &fileServiceImpl{
		filePathStrategy,
		fileEntityFactory,
		fileRepository,
		storageProvider,
	}
}

func (s *fileServiceImpl) ListFiles(
	filePaginationQuery *models.FilePaginationQuery,
) (*models.PaginationResult, errors.Error) {
	return s.fileRepository.ListFiles(filePaginationQuery)
}

func (s *fileServiceImpl) GetFile(fileId *models.FileId) (*entities.FileEntity, errors.Error) {
	return s.fileRepository.GetFile(fileId)
}

func (s *fileServiceImpl) UploadFile(fileSource io.Reader, data *models.FileUpload) (fileEntity *entities.FileEntity, err errors.Error) {
	fileEntity = s.fileEntityFactory.CreateFileEntity()
	fileEntity.Name = data.Name
	fileEntity.Type = data.Type
	fileEntity.Size = data.Size
	fileEntity.Path, err = s.filePathStrategy.BuildPath(fileEntity)

	if nil != err {
		return
	}

	if err = s.storageProvider.UploadFile(fileEntity, fileSource); nil != err {
		return
	}

	err = s.fileRepository.SaveFile(fileEntity)
	return
}

func (s *fileServiceImpl) DownloadFile(fileEntity *entities.FileEntity, fileDestination io.Writer) errors.Error {
	return s.storageProvider.DownloadFile(fileEntity, fileDestination)
}

func (s *fileServiceImpl) UpdateFile(fileEntity *entities.FileEntity, data *models.FileUpdate) errors.Error {
	fileEntity.Name = data.Name

	updated := time.Now().UTC()
	fileEntity.Updated = &updated

	return s.fileRepository.SaveFile(fileEntity)
}

func (s *fileServiceImpl) DeleteFile(fileEntity *entities.FileEntity) (err errors.Error) {
	err = s.fileRepository.RemoveFile(fileEntity)

	if nil != err {
		return
	}

	return s.storageProvider.DeleteFile(fileEntity)
}
