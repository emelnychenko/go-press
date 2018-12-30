package controllers

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewUserController", func(t *testing.T) {
		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userApi := mocks.NewMockUserApi(ctrl)
		userController, isUserController := NewUserController(userHttpHelper, userModelFactory, userApi).(*userControllerImpl)

		assert.True(t, isUserController)
		assert.Equal(t, userHttpHelper, userController.userHttpHelper)
		assert.Equal(t, userModelFactory, userController.userModelFactory)
		assert.Equal(t, userApi, userController.userApi)
	})

	t.Run("ListUsers", func(t *testing.T) {
		userPaginationQuery := new(models.UserPaginationQuery)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserPaginationQuery().Return(userPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(userPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(userPaginationQuery).Return(nil)

		var paginationResult *models.PaginationResult
		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().ListUsers(userPaginationQuery).Return(paginationResult, nil)

		userController := &userControllerImpl{
			userModelFactory: userModelFactory,
			userApi: userApi,
		}
		response, err := userController.ListUsers(httpContext)

		assert.Equal(t, paginationResult, response)
		assert.Nil(t, err)
	})

	t.Run("ListUsers:BindPaginationQueryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		userPaginationQuery := new(models.UserPaginationQuery)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserPaginationQuery().Return(userPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(userPaginationQuery.PaginationQuery).Return(systemErr)

		userController := &userControllerImpl{
			userModelFactory: userModelFactory,
		}
		response, err := userController.ListUsers(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ListUsers:BindUserPaginationQueryError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		userPaginationQuery := new(models.UserPaginationQuery)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserPaginationQuery().Return(userPaginationQuery)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(userPaginationQuery.PaginationQuery).Return(nil)
		httpContext.EXPECT().BindModel(userPaginationQuery).Return(systemErr)

		userController := &userControllerImpl{
			userModelFactory: userModelFactory,
		}
		response, err := userController.ListUsers(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetUser", func(t *testing.T) {
		userId := new(models.UserId)
		httpContext := mocks.NewMockHttpContext(ctrl)

		var user *models.User
		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().GetUser(userId).Return(user, nil)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{userHttpHelper: userHttpHelper, userApi: userApi}
		response, err := userController.GetUser(httpContext)

		assert.Equal(t, user, response)
		assert.Nil(t, err)
	})

	t.Run("GetUser:ParserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(nil, systemErr)

		userController := &userControllerImpl{userHttpHelper: userHttpHelper}
		response, err := userController.GetUser(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetUser:ApiError", func(t *testing.T) {
		userId := new(models.UserId)
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().GetUser(userId).Return(nil, systemErr)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{userHttpHelper: userHttpHelper, userApi: userApi}
		response, err := userController.GetUser(httpContext)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateUser", func(t *testing.T) {
		user := new(models.User)
		data := new(models.UserCreate)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserCreate().Return(data)

		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().CreateUser(data).Return(user, nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		userController := &userControllerImpl{
			userModelFactory: userModelFactory,
			userApi:          userApi,
		}
		response, err := userController.CreateUser(httpContext)

		assert.Equal(t, user, response)
		assert.Nil(t, err)
	})

	t.Run("CreateUser:BindUserUpdateError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.UserCreate)

		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserCreate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		userController := &userControllerImpl{
			userModelFactory: userModelFactory,
		}
		_, err := userController.CreateUser(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateUser:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.UserCreate)

		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserCreate().Return(data)

		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().CreateUser(data).Return(nil, systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		userController := &userControllerImpl{
			userModelFactory: userModelFactory,
			userApi:          userApi,
		}
		_, err := userController.CreateUser(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		userId := new(models.UserId)
		data := new(models.UserUpdate)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserUpdate().Return(data)

		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().UpdateUser(userId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{
			userHttpHelper:   userHttpHelper,
			userModelFactory: userModelFactory,
			userApi:          userApi,
		}
		_, err := userController.UpdateUser(httpContext)

		assert.Nil(t, err)
	})

	t.Run("UpdateUser:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(nil, systemErr)

		userController := &userControllerImpl{userHttpHelper: userHttpHelper}
		_, err := userController.UpdateUser(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateUser:BindUserUpdateError", func(t *testing.T) {
		userId := new(models.UserId)
		systemErr := common.NewUnknownError()
		data := new(models.UserUpdate)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserUpdate().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{
			userHttpHelper:   userHttpHelper,
			userModelFactory: userModelFactory,
		}
		_, err := userController.UpdateUser(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserIdentity", func(t *testing.T) {
		userId := new(models.UserId)
		data := new(models.UserChangeIdentity)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserChangeIdentity().Return(data)

		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().ChangeUserIdentity(userId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{
			userHttpHelper:   userHttpHelper,
			userModelFactory: userModelFactory,
			userApi:          userApi,
		}
		_, err := userController.ChangeUserIdentity(httpContext)

		assert.Nil(t, err)
	})

	t.Run("ChangeUserIdentity:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(nil, systemErr)

		userController := &userControllerImpl{userHttpHelper: userHttpHelper}
		_, err := userController.ChangeUserIdentity(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserIdentity:BindUserUpdateError", func(t *testing.T) {
		userId := new(models.UserId)
		systemErr := common.NewUnknownError()
		data := new(models.UserChangeIdentity)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserChangeIdentity().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{
			userHttpHelper:   userHttpHelper,
			userModelFactory: userModelFactory,
		}
		_, err := userController.ChangeUserIdentity(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserIdentity:ApiError", func(t *testing.T) {
		userId := new(models.UserId)
		systemErr := common.NewUnknownError()

		data := new(models.UserChangeIdentity)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserChangeIdentity().Return(data)

		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().ChangeUserIdentity(userId, data).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{
			userHttpHelper:   userHttpHelper,
			userModelFactory: userModelFactory,
			userApi:          userApi,
		}
		_, err := userController.ChangeUserIdentity(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserPassword", func(t *testing.T) {
		userId := new(models.UserId)
		data := new(models.UserChangePassword)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserChangePassword().Return(data)

		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().ChangeUserPassword(userId, data).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{
			userHttpHelper:   userHttpHelper,
			userModelFactory: userModelFactory,
			userApi:          userApi,
		}
		_, err := userController.ChangeUserPassword(httpContext)

		assert.Nil(t, err)
	})

	t.Run("ChangeUserPassword:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		httpContext := mocks.NewMockHttpContext(ctrl)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(nil, systemErr)

		userController := &userControllerImpl{userHttpHelper: userHttpHelper}
		_, err := userController.ChangeUserPassword(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserPassword:BindUserUpdateError", func(t *testing.T) {
		userId := new(models.UserId)
		systemErr := common.NewUnknownError()
		data := new(models.UserChangePassword)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserChangePassword().Return(data)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(systemErr)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{
			userHttpHelper:   userHttpHelper,
			userModelFactory: userModelFactory,
		}
		_, err := userController.ChangeUserPassword(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserPassword:ApiError", func(t *testing.T) {
		userId := new(models.UserId)
		systemErr := common.NewUnknownError()

		data := new(models.UserChangePassword)
		userModelFactory := mocks.NewMockUserModelFactory(ctrl)
		userModelFactory.EXPECT().CreateUserChangePassword().Return(data)

		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().ChangeUserPassword(userId, data).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)
		httpContext.EXPECT().BindModel(data).Return(nil)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{
			userHttpHelper:   userHttpHelper,
			userModelFactory: userModelFactory,
			userApi:          userApi,
		}
		_, err := userController.ChangeUserPassword(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		userId := new(models.UserId)

		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().DeleteUser(userId).Return(nil)

		httpContext := mocks.NewMockHttpContext(ctrl)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{userHttpHelper: userHttpHelper, userApi: userApi}
		_, err := userController.DeleteUser(httpContext)

		assert.Nil(t, err)
	})

	t.Run("DeleteUser:ParseError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		httpContext := mocks.NewMockHttpContext(ctrl)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(nil, systemErr)

		userController := &userControllerImpl{userHttpHelper: userHttpHelper}
		_, err := userController.DeleteUser(httpContext)

		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteUser:ApiError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		userId := new(models.UserId)

		userApi := mocks.NewMockUserApi(ctrl)
		userApi.EXPECT().DeleteUser(userId).Return(systemErr)

		httpContext := mocks.NewMockHttpContext(ctrl)

		userHttpHelper := mocks.NewMockUserHttpHelper(ctrl)
		userHttpHelper.EXPECT().ParseUserId(httpContext).Return(userId, nil)

		userController := &userControllerImpl{userHttpHelper: userHttpHelper, userApi: userApi}
		_, err := userController.DeleteUser(httpContext)

		assert.Equal(t, systemErr, err)
	})
}
