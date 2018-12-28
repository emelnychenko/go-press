package repositories

import (
	"errors"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileRepository(t *testing.T) {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	db, _ := gorm.Open(mocket.DriverName, "")
	fileRepository := NewFileRepository(db)

	fileId := common.NewModelId()
	commonReply := []map[string]interface{}{{
		"id": fileId.String(),
	}}

	t.Run("ListFiles", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		fileEntities, err := fileRepository.ListFiles()
		assert.IsType(t, []*entities.FileEntity{}, fileEntities)
		assert.Nil(t, err)
	})

	t.Run("ListFiles:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		fileEntities, err := fileRepository.ListFiles()
		assert.NotNil(t, fileEntities)
		assert.Error(t, err)
	})

	t.Run("GetFile", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		fileEntity, err := fileRepository.GetFile(fileId)
		assert.IsType(t, new(entities.FileEntity), fileEntity)
		assert.Nil(t, err)
	})

	t.Run("GetFile:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		fileEntity, err := fileRepository.GetFile(fileId)
		assert.NotNil(t, fileEntity)
		assert.Error(t, err)
	})

	t.Run("GetFile:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		fileEntity, err := fileRepository.GetFile(fileId)
		assert.NotNil(t, fileEntity)
		assert.Error(t, err, common.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("SaveFile", func(t *testing.T) {
		mocket.Catcher.Reset()

		fileEntity := entities.NewFileEntity()
		assert.Nil(t, fileRepository.SaveFile(fileEntity))
	})

	t.Run("SaveFile:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		fileEntity := new(entities.FileEntity)
		assert.Error(t, fileRepository.SaveFile(fileEntity))
	})

	t.Run("RemoveFile", func(t *testing.T) {
		mocket.Catcher.Reset()

		fileEntity := new(entities.FileEntity)
		assert.Nil(t, fileRepository.RemoveFile(fileEntity))
	})

	t.Run("RemoveFile:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		fileEntity := new(entities.FileEntity)
		assert.Error(t, fileRepository.RemoveFile(fileEntity))
	})
}
