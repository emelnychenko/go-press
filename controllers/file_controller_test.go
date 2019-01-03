package controllers

import (
	"bytes"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"os"
	"testing"
)

func TestFileController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewFileController", func(t *testing.T) {
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileApi := mocks.NewMockFileApi(ctrl)
		fileController, isFileController := NewFileController(fileHttpHelper, fileModelFactory, fileApi).(*fileControllerImpl)

		assert.True(t, isFileController)
		assert.Equal(t, fileHttpHelper, fileController.fileHttpHelper)
		assert.Equal(t, fileModelFactory, fileController.fileModelFactory)
		assert.Equal(t, fileApi, fileController.fileApi)
	})

	t.Run("ListFiles", func(t *testing.T) {
		filePaginationQuery := new(models.FilePaginationQuery)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFilePaginationQuery().Return(filePaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(filePaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(filePaginationQuery).Return(nil)

		var paginationResult *models.PaginationResult
		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().ListFiles(filePaginationQuery).Return(paginationResult, nil)

		fileController := &fileControllerImpl{
			fileModelFactory: fileModelFactory,
			fileApi:          fileApi,
		}
		response, err := fileController.ListFiles(httpContext)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListFiles:BindPaginationQueryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		filePaginationQuery := new(models.FilePaginationQuery)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFilePaginationQuery().Return(filePaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(filePaginationQuery.PaginationQuery).Return(systemErr)

		fileController := &fileControllerImpl{
			fileModelFactory: fileModelFactory,
		}
		response, err := fileController.ListFiles(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListFiles:BindFilePaginationQueryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		filePaginationQuery := new(models.FilePaginationQuery)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFilePaginationQuery().Return(filePaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(filePaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(filePaginationQuery).Return(systemErr)

		fileController := &fileControllerImpl{
			fileModelFactory: fileModelFactory,
		}
		response, err := fileController.ListFiles(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetFile", func(t *testing.T) {
		fileId := new(models.FileId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		var file *models.File
		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().GetFile(fileId).Return(file, nil)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(fileId, nil)

		fileController := &fileControllerImpl{fileHttpHelper: fileHttpHelper, fileApi: fileApi}
		response, err := fileController.GetFile(httpContext)

		assert.Equal(t, file, response)
		assert.Nil(t, err)
	})

	t.Run("GetFile:ParserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(nil, systemErr)

		fileController := &fileControllerImpl{fileHttpHelper: fileHttpHelper}
		response, err := fileController.GetFile(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetFile:ApiError", func(t *testing.T) {
		fileId := new(models.FileId)
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().GetFile(fileId).Return(nil, systemErr)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(fileId, nil)

		fileController := &fileControllerImpl{fileHttpHelper: fileHttpHelper, fileApi: fileApi}
		response, err := fileController.GetFile(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UploadFile", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)

		var fileSize int64 = 99
		fileName := "foo"
		fileContentType := "application/test"
		fileHeader := &multipart.FileHeader{
			Size: fileSize,
			Header: map[string][]string{
				echo.HeaderContentType: {fileContentType},
			},
			Filename: fileName,
		}
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().GetFileHeader(httpContext).Return(fileHeader, nil)

		data := new(models.FileUpload)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpload().Return(data)

		fileSource := new(os.File)
		fileHttpHelper.EXPECT().OpenFormFile(fileHeader).Return(fileSource, nil)

		file := new(models.File)
		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().UploadFile(fileSource, data).Return(file, nil)

		fileController := &fileControllerImpl{
			fileHttpHelper:   fileHttpHelper,
			fileModelFactory: fileModelFactory,
			fileApi:          fileApi,
		}
		response, err := fileController.UploadFile(httpContext)

		assert.Nil(t, err)
		assert.Equal(t, file, response)
		assert.Equal(t, fileSize, data.Size)
		assert.Equal(t, fileName, data.Name)
		assert.Equal(t, fileContentType, data.Type)
	})

	t.Run("UploadFile:FileHeaderError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().GetFileHeader(httpContext).Return(nil, systemErr)

		fileController := &fileControllerImpl{
			fileHttpHelper: fileHttpHelper,
		}
		response, err := fileController.UploadFile(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UploadFile:FormFileError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		fileHeader := &multipart.FileHeader{
			Header: map[string][]string{},
		}
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().GetFileHeader(httpContext).Return(fileHeader, nil)

		data := new(models.FileUpload)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpload().Return(data)

		fileHttpHelper.EXPECT().OpenFormFile(fileHeader).Return(nil, systemErr)

		fileController := &fileControllerImpl{
			fileHttpHelper:   fileHttpHelper,
			fileModelFactory: fileModelFactory,
		}
		response, err := fileController.UploadFile(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UploadFile:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		fileHeader := &multipart.FileHeader{
			Header: map[string][]string{},
		}
		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().GetFileHeader(httpContext).Return(fileHeader, nil)

		data := new(models.FileUpload)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpload().Return(data)

		fileSource := new(os.File)
		fileHttpHelper.EXPECT().OpenFormFile(fileHeader).Return(fileSource, nil)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().UploadFile(fileSource, data).Return(nil, systemErr)

		fileController := &fileControllerImpl{
			fileHttpHelper:   fileHttpHelper,
			fileModelFactory: fileModelFactory,
			fileApi:          fileApi,
		}
		response, err := fileController.UploadFile(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("DownloadFile", func(t *testing.T) {
		fileId := new(models.FileId)

		httpContext := mocks.NewMockHttpContext(ctrl)
		var prepareFileDestination contracts.PrepareFileDestination = func(file *models.File) (fileDestination io.Writer) {
			return bytes.NewBufferString("")
		}

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(fileId, nil)
		fileHttpHelper.EXPECT().PrepareFileDestination(httpContext).Return(prepareFileDestination)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().DownloadFile(fileId, gomock.AssignableToTypeOf(prepareFileDestination)).Return(nil)

		fileController := &fileControllerImpl{
			fileHttpHelper: fileHttpHelper,
			fileApi:        fileApi,
		}
		_, err := fileController.DownloadFile(httpContext)

		assert.Nil(t, err)
	})

	t.Run("DownloadFile:ParseFileIdError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(nil, systemErr)

		fileController := &fileControllerImpl{
			fileHttpHelper: fileHttpHelper,
		}
		_, err := fileController.DownloadFile(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DownloadFile:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		fileId := new(models.FileId)

		httpContext := mocks.NewMockHttpContext(ctrl)
		var prepareFileDestination contracts.PrepareFileDestination

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(fileId, nil)
		fileHttpHelper.EXPECT().PrepareFileDestination(httpContext).Return(prepareFileDestination)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().DownloadFile(fileId, prepareFileDestination).Return(systemErr)

		fileController := &fileControllerImpl{
			fileHttpHelper: fileHttpHelper,
			fileApi:        fileApi,
		}
		_, err := fileController.DownloadFile(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateFile", func(t *testing.T) {
		fileId := new(models.FileId)
		data := new(models.FileUpdate)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpdate().Return(data)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().UpdateFile(fileId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(fileId, nil)

		fileController := &fileControllerImpl{
			fileHttpHelper:   fileHttpHelper,
			fileModelFactory: fileModelFactory,
			fileApi:          fileApi,
		}
		_, err := fileController.UpdateFile(httpContext)

		assert.Nil(t, err)
	})

	t.Run("UpdateFile:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(nil, systemErr)

		fileController := &fileControllerImpl{fileHttpHelper: fileHttpHelper}
		_, err := fileController.UpdateFile(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateFile:BindFileUpdateError", func(t *testing.T) {
		fileId := new(models.FileId)
		systemErr := common.NewUnknownError()
		data := new(models.FileUpdate)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpdate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(fileId, nil)

		fileController := &fileControllerImpl{
			fileHttpHelper:   fileHttpHelper,
			fileModelFactory: fileModelFactory,
		}
		_, err := fileController.UpdateFile(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateFile:ApiError", func(t *testing.T) {
		fileId := new(models.FileId)
		systemErr := common.NewUnknownError()

		data := new(models.FileUpdate)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpdate().Return(data)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().UpdateFile(fileId, data).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(fileId, nil)

		fileController := &fileControllerImpl{
			fileHttpHelper:   fileHttpHelper,
			fileModelFactory: fileModelFactory,
			fileApi:          fileApi,
		}
		_, err := fileController.UpdateFile(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteFile", func(t *testing.T) {
		fileId := new(models.FileId)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().DeleteFile(fileId).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(fileId, nil)

		fileController := &fileControllerImpl{fileHttpHelper: fileHttpHelper, fileApi: fileApi}
		_, err := fileController.DeleteFile(httpContext)

		assert.Nil(t, err)
	})

	t.Run("DeleteFile:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		httpContext := mocks.NewMockHttpContext(ctrl)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(nil, systemErr)

		fileController := &fileControllerImpl{fileHttpHelper: fileHttpHelper}
		_, err := fileController.DeleteFile(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteFile:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		fileId := new(models.FileId)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().DeleteFile(fileId).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)

		fileHttpHelper := mocks.NewMockFileHttpHelper(ctrl)
		fileHttpHelper.EXPECT().ParseFileId(httpContext).Return(fileId, nil)

		fileController := &fileControllerImpl{fileHttpHelper: fileHttpHelper, fileApi: fileApi}
		_, err := fileController.DeleteFile(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
