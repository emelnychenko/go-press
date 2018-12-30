package repositories

import (
	"errors"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbPaginator := mocks.NewMockDbPaginator(ctrl)
	db, _ := gorm.Open(mocket.DriverName, "")
	userRepository, isUserRepository := NewUserRepository(db, dbPaginator).(*userRepositoryImpl)

	assert.True(t, isUserRepository)
	assert.Equal(t, db, userRepository.db)
	assert.Equal(t, dbPaginator, userRepository.dbPaginator)

	userId := common.NewModelId()
	userIdString := userId.String()
	commonReply := []map[string]interface{}{{
		"id":        userIdString,
		"firstName": userIdString,
		"lastName":  userIdString,
	}}

	t.Run("ListUsers", func(t *testing.T) {
		userPaginationQuery := &models.UserPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), userPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(nil)

		paginationResult, err := userRepository.ListUsers(userPaginationQuery)
		assert.IsType(t, []*entities.UserEntity{}, paginationResult.Data)
		assert.Nil(t, err)
	})

	t.Run("ListUsers:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		userPaginationQuery := &models.UserPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = errors.New("")
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), userPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		userEntities, err := userRepository.ListUsers(userPaginationQuery)
		assert.Nil(t, userEntities)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetUser", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		result, err := userRepository.GetUser(userId)
		assert.IsType(t, &entities.UserEntity{}, result)
		assert.Nil(t, err)
	})

	t.Run("GetUser:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		userEntity, err := userRepository.GetUser(userId)
		assert.NotNil(t, userEntity)
		assert.Error(t, err)
	})

	t.Run("GetUser:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		userEntity, err := userRepository.GetUser(userId)
		assert.NotNil(t, userEntity)
		assert.Error(t, err, common.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("LookupUser", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		result, err := userRepository.LookupUser("")
		assert.IsType(t, &entities.UserEntity{}, result)
		assert.Nil(t, err)
	})

	t.Run("LookupUser:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		userEntity, err := userRepository.LookupUser("")
		assert.NotNil(t, userEntity)
		assert.Error(t, err)
	})

	t.Run("LookupUser:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		userEntity, err := userRepository.LookupUser("")
		assert.NotNil(t, userEntity)
		assert.Error(t, err, common.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("SaveUser", func(t *testing.T) {
		mocket.Catcher.Reset()

		err := userRepository.SaveUser(entities.NewUserEntity())
		assert.Nil(t, err)
	})

	t.Run("SaveUser:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		err := userRepository.SaveUser(entities.NewUserEntity())
		assert.Error(t, err)
	})

	t.Run("RemoveUser", func(t *testing.T) {
		mocket.Catcher.Reset()

		err := userRepository.RemoveUser(entities.NewUserEntity())
		assert.Nil(t, err)
	})

	t.Run("RemoveUser:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = errors.New("")

		err := userRepository.RemoveUser(entities.NewUserEntity())
		assert.Error(t, err)
	})
}
