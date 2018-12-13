package user

import (
	. "../common"
	"../user_domain"
	"errors"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

var DB *gorm.DB

func SetupRepository() *userRepositoryImpl {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	db, _ := gorm.Open(mocket.DriverName, "connection_string") // Can be any connection string
	DB = db

	return NewUserRepository(db)
}

func TestUserRepository(t *testing.T) {
	repository := SetupRepository()

	userId := NewModelId()
	userIdString := userId.String()
	commonReply := []map[string]interface{}{{
		"id": userIdString,
		"firstName": userIdString,
		"lastName": userIdString,
	}}

	t.Run("GetDb()", func(t *testing.T) {
		db := repository.db
		assert.IsType(t, DB, db)
	})

	t.Run("ListUsers() thrown ServerError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		results, err := repository.ListUsers()
		assert.Nil(t, results)
		assert.Error(t, err)
	})

	t.Run("ListUsers()", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		results, err := repository.ListUsers()
		assert.IsType(t, []*user_domain.UserEntity{}, results)
		assert.Nil(t, err)
	})

	t.Run("GetUser(UUID) thrown UserNotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset()

		result, err := repository.GetUser(userId)
		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("GetUser(UUID)", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		result, err := repository.GetUser(userId)
		assert.IsType(t, &user_domain.UserEntity{}, result)
		assert.Nil(t, err)
	})

	t.Run("LookupUser(UUID) thrown UserNotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset()

		result, err := repository.LookupUser("")
		assert.Nil(t, result)
		assert.Error(t, err)
	})

	t.Run("LookupUser(UUID)", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		result, err := repository.LookupUser("")
		assert.IsType(t, &user_domain.UserEntity{}, result)
		assert.Nil(t, err)
	})

	t.Run("SaveUser(UUID,UserUpdate) thrown Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		err := repository.SaveUser(user_domain.NewUserEntity())
		assert.Error(t, err)
	})

	t.Run("SaveUser(UUID, UserUpdate)", func(t *testing.T) {
		mocket.Catcher.Reset()

		err := repository.SaveUser(user_domain.NewUserEntity())
		assert.Nil(t, err)
	})

	t.Run("RemoveUser(UUID) thrown Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		err := repository.RemoveUser(user_domain.NewUserEntity())
		assert.Error(t, err)
	})

	t.Run("RemoveUser(UUID)", func(t *testing.T) {
		mocket.Catcher.Reset()

		err := repository.RemoveUser(user_domain.NewUserEntity())
		assert.Nil(t, err)
	})
}