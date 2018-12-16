package controllers

import (
	"bytes"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/echo_mocks"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

// TODO: Assert failed status codes
func TestFileController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fileId := models.NewModelId()
	testErr := common.ServerError("err0")

	t.Run("NewFileController", func(t *testing.T) {
		fileParamHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileApi := mocks.NewMockFileApi(ctrl)
		fileController := NewFileController(fileParamHelper, fileModelFactory, fileApi)

		assert.Equal(t, fileParamHelper, fileController.fileEchoHelper)
		assert.Equal(t, fileModelFactory, fileController.fileModelFactory)
		assert.Equal(t, fileApi, fileController.fileApi)
	})

	t.Run("ListFiles", func(t *testing.T) {
		var replies []*models.File
		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().ListFiles().Return(replies, nil)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().JSON(http.StatusOK, replies).Return(nil)

		fileController := &FileController{fileApi: fileApi}
		assert.Nil(t, fileController.ListFiles(context))
	})

	t.Run("ListFiles:Error", func(t *testing.T) {
		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().ListFiles().Return(nil, testErr)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)

		fileController := &FileController{fileApi: fileApi}
		assert.Nil(t, fileController.ListFiles(context))
	})

	t.Run("GetFile", func(t *testing.T) {
		var reply *models.File
		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().GetFile(fileId).Return(reply, nil)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().JSON(http.StatusOK, reply).Return(nil)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(fileId, nil)

		fileController := &FileController{fileEchoHelper: fileEchoHelper, fileApi: fileApi}
		assert.Nil(t, fileController.GetFile(context))
	})

	t.Run("GetFile:ParserError", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(nil, testErr)

		fileController := &FileController{fileEchoHelper: fileEchoHelper}
		assert.Nil(t, fileController.GetFile(context))
	})

	t.Run("GetFile:ApiError", func(t *testing.T) {
		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().GetFile(fileId).Return(nil, testErr)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(fileId, nil)

		fileController := &FileController{fileEchoHelper: fileEchoHelper, fileApi: fileApi}
		assert.Nil(t, fileController.GetFile(context))
	})

	t.Run("UploadFile", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)

		var fileSize int64 = 99
		fileName := "foo"
		fileContentType := "application/test"
		fileHeader := &multipart.FileHeader{
			Size: fileSize,
			Header: map[string][]string{
				echo.HeaderContentType: []string{fileContentType},
			},
			Filename: fileName,
		}
		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().GetFileHeader(context).Return(fileHeader, nil)

		data := new(models.FileUpload)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpload().Return(data)

		fileSource := new(os.File)
		fileEchoHelper.EXPECT().OpenFormFile(fileHeader).Return(fileSource, nil)

		file := new(models.File)
		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().UploadFile(fileSource, data).Return(file, nil)

		context.EXPECT().JSON(http.StatusOK, file).Return(nil)
		fileController := &FileController{
			fileEchoHelper:   fileEchoHelper,
			fileModelFactory: fileModelFactory,
			fileApi:          fileApi,
		}
		assert.Nil(t, fileController.UploadFile(context))
		assert.Equal(t, fileSize, data.Size)
		assert.Equal(t, fileName, data.Name)
		assert.Equal(t, fileContentType, data.Type)
	})

	t.Run("UploadFile:FileHeaderError", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().GetFileHeader(context).Return(nil, testErr)

		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)
		fileController := &FileController{
			fileEchoHelper:   fileEchoHelper,
		}
		assert.Nil(t, fileController.UploadFile(context))
	})

	t.Run("UploadFile:FormFileError", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)

		fileHeader := &multipart.FileHeader{
			Header: map[string][]string{},
		}
		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().GetFileHeader(context).Return(fileHeader, nil)

		data := new(models.FileUpload)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpload().Return(data)

		fileEchoHelper.EXPECT().OpenFormFile(fileHeader).Return(nil, testErr)

		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)
		fileController := &FileController{
			fileEchoHelper:   fileEchoHelper,
			fileModelFactory: fileModelFactory,
		}
		assert.Nil(t, fileController.UploadFile(context))
	})

	t.Run("UploadFile:ApiError", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)

		fileHeader := &multipart.FileHeader{
			Header: map[string][]string{},
		}
		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().GetFileHeader(context).Return(fileHeader, nil)

		data := new(models.FileUpload)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpload().Return(data)

		fileSource := new(os.File)
		fileEchoHelper.EXPECT().OpenFormFile(fileHeader).Return(fileSource, nil)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().UploadFile(fileSource, data).Return(nil, testErr)

		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)
		fileController := &FileController{
			fileEchoHelper:   fileEchoHelper,
			fileModelFactory: fileModelFactory,
			fileApi:          fileApi,
		}
		assert.Nil(t, fileController.UploadFile(context))
	})

	t.Run("DownloadFile", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)
		var prepareFileDestination contracts.PrepareFileDestination = func(file *models.File) (fileDestination io.Writer) {
			return bytes.NewBufferString("")
		}

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(fileId, nil)
		fileEchoHelper.EXPECT().PrepareFileDestination(context).Return(prepareFileDestination)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().DownloadFile(fileId, gomock.AssignableToTypeOf(prepareFileDestination)).Return(nil)

		fileController := &FileController{
			fileEchoHelper: fileEchoHelper,
			fileApi:        fileApi,
		}
		assert.Nil(t, fileController.DownloadFile(context))
	})

	t.Run("DownloadFile:ParseIdError", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(nil, testErr)

		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)

		fileController := &FileController{
			fileEchoHelper: fileEchoHelper,
		}
		assert.Nil(t, fileController.DownloadFile(context))
	})

	t.Run("DownloadFile:ApiError", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)
		var prepareFileDestination contracts.PrepareFileDestination

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(fileId, nil)
		fileEchoHelper.EXPECT().PrepareFileDestination(context).Return(prepareFileDestination)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().DownloadFile(fileId, prepareFileDestination).Return(testErr)

		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)

		fileController := &FileController{
			fileEchoHelper: fileEchoHelper,
			fileApi:        fileApi,
		}
		assert.Nil(t, fileController.DownloadFile(context))
	})

	t.Run("UpdateFile", func(t *testing.T) {
		data := new(models.FileUpdate)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpdate().Return(data)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().UpdateFile(fileId, data).Return(nil)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().Bind(data).Return(nil)
		context.EXPECT().JSON(http.StatusOK, nil).Return(nil)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(fileId, nil)

		fileController := &FileController{
			fileEchoHelper:   fileEchoHelper,
			fileModelFactory: fileModelFactory,
			fileApi:          fileApi,
		}
		assert.Nil(t, fileController.UpdateFile(context))
	})

	t.Run("UpdateFile:ParseError", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(nil, testErr)

		fileController := &FileController{fileEchoHelper: fileEchoHelper}
		assert.Nil(t, fileController.UpdateFile(context))
	})

	t.Run("UpdateFile:BindError", func(t *testing.T) {
		data := new(models.FileUpdate)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpdate().Return(data)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().Bind(data).Return(testErr)
		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(fileId, nil)

		fileController := &FileController{
			fileEchoHelper:   fileEchoHelper,
			fileModelFactory: fileModelFactory,
		}
		assert.Nil(t, fileController.UpdateFile(context))
	})

	t.Run("UpdateFile:ApiError", func(t *testing.T) {
		data := new(models.FileUpdate)
		fileModelFactory := mocks.NewMockFileModelFactory(ctrl)
		fileModelFactory.EXPECT().CreateFileUpdate().Return(data)

		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().UpdateFile(fileId, data).Return(testErr)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().Bind(data).Return(nil)
		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(fileId, nil)

		fileController := &FileController{
			fileEchoHelper:   fileEchoHelper,
			fileModelFactory: fileModelFactory,
			fileApi:          fileApi,
		}
		assert.Nil(t, fileController.UpdateFile(context))
	})

	t.Run("DeleteFile", func(t *testing.T) {
		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().DeleteFile(fileId).Return(nil)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().JSON(http.StatusOK, nil).Return(nil)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(fileId, nil)

		fileController := &FileController{fileEchoHelper: fileEchoHelper, fileApi: fileApi}
		assert.Nil(t, fileController.DeleteFile(context))
	})

	t.Run("DeleteFile:ParseError", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(nil, testErr)

		fileController := &FileController{fileEchoHelper: fileEchoHelper}
		assert.Nil(t, fileController.DeleteFile(context))
	})

	t.Run("DeleteFile:ApiError", func(t *testing.T) {
		fileApi := mocks.NewMockFileApi(ctrl)
		fileApi.EXPECT().DeleteFile(fileId).Return(testErr)

		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().JSON(gomock.Not(http.StatusOK), testErr).Return(nil)

		fileEchoHelper := mocks.NewMockFileEchoHelper(ctrl)
		fileEchoHelper.EXPECT().ParseId(context).Return(fileId, nil)

		fileController := &FileController{fileEchoHelper: fileEchoHelper, fileApi: fileApi}
		assert.Nil(t, fileController.DeleteFile(context))
	})
}
