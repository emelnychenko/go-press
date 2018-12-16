package helpers

import (
	"github.com/emelnychenko/go-press/echo_mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileEchoHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var fileId = models.NewModelId()

	t.Run("NewFileEchoHelper", func(t *testing.T) {
		_, isFileParamParser := NewFileEchoHelper().(*fileEchoHelperImpl)
		assert.True(t, isFileParamParser)
	})

	t.Run("ParseId", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().Param(FileIdParam).Return(fileId.String())

		fileEchoHelper := &fileEchoHelperImpl{}
		parsedFileId, err := fileEchoHelper.ParseId(context)
		assert.Equal(t, fileId.String(), parsedFileId.String())
		assert.Nil(t, err)
	})

	t.Run("GetFileHeader", func(t *testing.T) {
		context := echo_mocks.NewMockContext(ctrl)
		fileHeader := &multipart.FileHeader{}
		context.EXPECT().FormFile(FileFormFile).Return(fileHeader, nil)

		fileEchoHelper := &fileEchoHelperImpl{}
		fileHeaderResult, err := fileEchoHelper.GetFileHeader(context)
		assert.Equal(t, fileHeader, fileHeaderResult)
		assert.Nil(t, err)
	})

	t.Run("OpenFormFile", func(t *testing.T) {
		fileHeader := &multipart.FileHeader{}

		fileEchoHelper := &fileEchoHelperImpl{}
		formFile, err := fileEchoHelper.OpenFormFile(fileHeader)
		assert.Nil(t, formFile)
		assert.Error(t, err)
	})

	t.Run("PrepareFileDestination", func(t *testing.T) {
		responseWriter := httptest.NewRecorder()

		response := &echo.Response{Writer: responseWriter}
		context := echo_mocks.NewMockContext(ctrl)
		context.EXPECT().Response().Return(response)

		fileType := "application/test"
		var fileSize int64 = 99
		file := &models.File{
			Size:fileSize,
			Type:fileType,
		}

		fileEchoHelper := &fileEchoHelperImpl{}
		prepareFileDestination := fileEchoHelper.PrepareFileDestination(context)
		writer := prepareFileDestination(file)

		assert.Equal(t, response, writer)

		assert.Equal(t, http.StatusOK, responseWriter.Code)
		assert.Equal(t, fileType, responseWriter.Header().Get(echo.HeaderContentType))
		assert.Equal(t, string(fileSize), responseWriter.Header().Get(echo.HeaderContentLength))
	})
}
