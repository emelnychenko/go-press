package repositories

import (
	"errors"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func SetupUserRepository() contracts.UserRepository {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	db, _ := gorm.Open(mocket.DriverName, "")

	return NewUserRepository(db)
}

func TestUserRepository(t *testing.T) {
	repository := SetupUserRepository()

	userId := common.NewModelId()
	userIdString := userId.String()
	commonReply := []map[string]interface{}{{
		"id":        userIdString,
		"firstName": userIdString,
		"lastName":  userIdString,
	}}

	t.Run("ListUsers", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		results, err := repository.ListUsers()
		assert.IsType(t, []*entities.UserEntity{}, results)
		assert.Nil(t, err)
	})

	t.Run("ListUsers:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		results, err := repository.ListUsers()
		assert.NotNil(t, results)
		assert.Error(t, err)
	})

	t.Run("GetUser", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		result, err := repository.GetUser(userId)
		assert.IsType(t, &entities.UserEntity{}, result)
		assert.Nil(t, err)
	})

	t.Run("GetUser:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		userEntity, err := repository.GetUser(userId)
		assert.NotNil(t, userEntity)
		assert.Error(t, err)
	})

	t.Run("GetUser:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		userEntity, err := repository.GetUser(userId)
		assert.NotNil(t, userEntity)
		assert.Error(t, err, common.NewSystemError(gorm.ErrInvalidSQL))
	})

	t.Run("LookupUser", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		result, err := repository.LookupUser("")
		assert.IsType(t, &entities.UserEntity{}, result)
		assert.Nil(t, err)
	})

	t.Run("LookupUser:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		userEntity, err := repository.LookupUser("")
		assert.NotNil(t, userEntity)
		assert.Error(t, err)
	})

	t.Run("LookupUser:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		userEntity, err := repository.LookupUser("")
		assert.NotNil(t, userEntity)
		assert.Error(t, err, common.NewSystemError(gorm.ErrInvalidSQL))
	})

	t.Run("SaveUser", func(t *testing.T) {
		mocket.Catcher.Reset()

		err := repository.SaveUser(entities.NewUserEntity())
		assert.Nil(t, err)
	})

	t.Run("SaveUser:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		err := repository.SaveUser(entities.NewUserEntity())
		assert.Error(t, err)
	})

	t.Run("RemoveUser", func(t *testing.T) {
		mocket.Catcher.Reset()

		err := repository.RemoveUser(entities.NewUserEntity())
		assert.Nil(t, err)
	})

	t.Run("RemoveUser:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		err := repository.RemoveUser(entities.NewUserEntity())
		assert.Error(t, err)
	})
}
