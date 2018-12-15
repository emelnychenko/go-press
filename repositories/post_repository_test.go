package repositories

import (
	"errors"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	errors2 "github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostRepository(t *testing.T) {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	db, _ := gorm.Open(mocket.DriverName, "")
	postRepository := NewPostRepository(db)

	postId := models.NewModelId()
	commonReply := []map[string]interface{}{{
		"id": postId.String(),
	}}

	t.Run("ListPosts", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		postEntities, err := postRepository.ListPosts()
		assert.IsType(t, []*entities.PostEntity{}, postEntities)
		assert.Nil(t, err)
	})

	t.Run("ListPosts:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		postEntities, err := postRepository.ListPosts()
		assert.Nil(t, postEntities)
		assert.Error(t, err)
	})

	t.Run("GetPost", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		postEntity, err := postRepository.GetPost(postId)
		assert.IsType(t, new(entities.PostEntity), postEntity)
		assert.Nil(t, err)
	})

	t.Run("GetPost:PostNotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		postEntity, err := postRepository.GetPost(postId)
		assert.Nil(t, postEntity)
		assert.Error(t, err, errors2.PostNotFoundError{})
	})

	t.Run("GetPost:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		postEntity, err := postRepository.GetPost(postId)
		assert.Nil(t, postEntity)
		assert.Error(t, err, common.NewServerError(gorm.ErrInvalidSQL))
	})

	t.Run("SavePost", func(t *testing.T) {
		mocket.Catcher.Reset()

		postEntity := entities.NewPostEntity()
		assert.Nil(t, postRepository.SavePost(postEntity))
	})

	t.Run("SavePost:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		postEntity := new(entities.PostEntity)
		assert.Error(t, postRepository.SavePost(postEntity))
	})

	t.Run("RemovePost", func(t *testing.T) {
		mocket.Catcher.Reset()

		postEntity := new(entities.PostEntity)
		assert.Nil(t, postRepository.RemovePost(postEntity))
	})

	t.Run("RemovePost:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		postEntity := new(entities.PostEntity)
		assert.Error(t, postRepository.RemovePost(postEntity))
	})
}
