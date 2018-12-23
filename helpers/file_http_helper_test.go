package helpers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileHttpHelper(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var fileId = common.NewModelId()

	t.Run("NewFileHttpHelper", func(t *testing.T) {
		_, isFileParamParser := NewFileHttpHelper().(*fileHttpHelperImpl)
		assert.True(t, isFileParamParser)
	})

	t.Run("ParseFileId", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Parameter(FileIdParameterName).Return(fileId.String())

		fileHttpHelper := &fileHttpHelperImpl{}
		parsedFileId, err := fileHttpHelper.ParseFileId(httpContext)
		assert.Equal(t, fileId.String(), parsedFileId.String())
		assert.Nil(t, err)
	})

	t.Run("GetFileHeader", func(t *testing.T) {
		httpContext := mocks.NewMockHttpContext(ctrl)
		fileHeader := &multipart.FileHeader{}
		httpContext.EXPECT().FormFile(FileFormFileName).Return(fileHeader, nil)

		fileHttpHelper := &fileHttpHelperImpl{}
		response, err := fileHttpHelper.GetFileHeader(httpContext)
		assert.Equal(t, fileHeader, response)
		assert.Nil(t, err)
	})

	t.Run("OpenFormFile", func(t *testing.T) {
		fileHeader := &multipart.FileHeader{}

		fileHttpHelper := &fileHttpHelperImpl{}
		formFile, err := fileHttpHelper.OpenFormFile(fileHeader)
		assert.Nil(t, formFile)
		assert.Error(t, err)
	})

	t.Run("PrepareFileDestination", func(t *testing.T) {
		responseWriter := httptest.NewRecorder()

		response := &echo.Response{Writer: responseWriter}
		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().Response().Return(response)

		fileType := "application/test"
		var fileSize int64 = 99
		file := &models.File{
			Size: fileSize,
			Type: fileType,
		}

		fileHttpHelper := &fileHttpHelperImpl{}
		prepareFileDestination := fileHttpHelper.PrepareFileDestination(httpContext)
		writer := prepareFileDestination(file)

		assert.Equal(t, response, writer)

		assert.Equal(t, http.StatusOK, responseWriter.Code)
		assert.Equal(t, fileType, responseWriter.Header().Get(echo.HeaderContentType))
		assert.Equal(t, string(fileSize), responseWriter.Header().Get(echo.HeaderContentLength))
	})
}
