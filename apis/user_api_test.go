package apis

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/events"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NewUserApi", func(t *testing.T) {
		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		userEventFactory := mocks.NewMockUserEventFactory(ctrl)
		userService := mocks.NewMockUserService(ctrl)
		userAggregator := mocks.NewMockUserAggregator(ctrl)

		userApi, isUserApi := NewUserApi(
			eventDispatcher, userEventFactory, userService, userAggregator,
		).(*userApiImpl)

		assert.True(t, isUserApi)
		assert.Equal(t, eventDispatcher, userApi.eventDispatcher)
		assert.Equal(t, userEventFactory, userApi.userEventFactory)
		assert.Equal(t, userService, userApi.userService)
		assert.Equal(t, userAggregator, userApi.userAggregator)
	})

	t.Run("ListUsers", func(t *testing.T) {
		var userEntities []*entities.UserEntity
		var users []*models.User

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().ListUsers().Return(userEntities, nil)

		userAggregator := mocks.NewMockUserAggregator(ctrl)
		userAggregator.EXPECT().AggregateUsers(userEntities).Return(users)

		userApi := &userApiImpl{userService: userService, userAggregator: userAggregator}
		response, err := userApi.ListUsers()

		assert.Equal(t, users, response)
		assert.Nil(t, err)
	})

	t.Run("ListUsers:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().ListUsers().Return(nil, systemErr)

		userApi := &userApiImpl{userService: userService}
		response, err := userApi.ListUsers()

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetUser", func(t *testing.T) {
		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)
		user := new(models.User)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)

		userAggregator := mocks.NewMockUserAggregator(ctrl)
		userAggregator.EXPECT().AggregateUser(userEntity).Return(user)

		userApi := &userApiImpl{userService: userService, userAggregator: userAggregator}
		response, err := userApi.GetUser(userId)

		assert.Equal(t, user, response)
		assert.Nil(t, err)
	})

	t.Run("GetUser:GetUserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		userId := new(models.UserId)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(nil, systemErr)

		userApi := &userApiImpl{userService: userService}
		response, err := userApi.GetUser(userId)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("CreateUser", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		user := new(models.User)
		data := new(models.UserCreate)

		userEvent := new(events.UserEvent)
		userEventFactory := mocks.NewMockUserEventFactory(ctrl)
		userEventFactory.EXPECT().CreateUserCreatedEvent(userEntity).Return(userEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(userEvent)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().CreateUser(data).Return(userEntity, nil)

		userAggregator := mocks.NewMockUserAggregator(ctrl)
		userAggregator.EXPECT().AggregateUser(userEntity).Return(user)

		userApi := &userApiImpl{
			eventDispatcher: eventDispatcher,
			userEventFactory: userEventFactory,
			userService: userService,
			userAggregator: userAggregator,
		}
		response, err := userApi.CreateUser(data)

		assert.Equal(t, user, response)
		assert.Nil(t, err)
	})

	t.Run("CreateUser:Error", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		data := new(models.UserCreate)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().CreateUser(data).Return(nil, systemErr)

		userApi := &userApiImpl{userService: userService}
		response, err := userApi.CreateUser(data)

		assert.Nil(t, response)
		assert.Equal(t, systemErr, err)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)
		data := new(models.UserUpdate)

		userEvent := new(events.UserEvent)
		userEventFactory := mocks.NewMockUserEventFactory(ctrl)
		userEventFactory.EXPECT().CreateUserUpdatedEvent(userEntity).Return(userEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(userEvent)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)
		userService.EXPECT().UpdateUser(userEntity, data).Return(nil)

		userApi := &userApiImpl{
			eventDispatcher: eventDispatcher,
			userEventFactory: userEventFactory,
			userService: userService,
		}
		assert.Nil(t, userApi.UpdateUser(userId, data))
	})

	t.Run("UpdateUser:GetUserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		userId := new(models.UserId)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(nil, systemErr)

		userApi := &userApiImpl{userService: userService}
		err := userApi.UpdateUser(userId, nil)
		assert.Equal(t, systemErr, err)
	})

	t.Run("VerifyUser", func(t *testing.T) {
		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)

		userEvent := new(events.UserEvent)
		userEventFactory := mocks.NewMockUserEventFactory(ctrl)
		userEventFactory.EXPECT().CreateUserVerifiedEvent(userEntity).Return(userEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(userEvent)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)
		userService.EXPECT().VerifyUser(userEntity).Return(nil)

		userApi := &userApiImpl{
			eventDispatcher: eventDispatcher,
			userEventFactory: userEventFactory,
			userService: userService,
		}
		assert.Nil(t, userApi.VerifyUser(userId))
	})

	t.Run("VerifyUser:GetUserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		userId := new(models.UserId)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(nil, systemErr)

		userApi := &userApiImpl{userService: userService}
		err := userApi.VerifyUser(userId)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserIdentity", func(t *testing.T) {
		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)
		data := new(models.UserChangeIdentity)

		userEvent := new(events.UserEvent)
		userEventFactory := mocks.NewMockUserEventFactory(ctrl)
		userEventFactory.EXPECT().CreateUserIdentityChangedEvent(userEntity).Return(userEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(userEvent)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)
		userService.EXPECT().ChangeUserIdentity(userEntity, data).Return(nil)

		userApi := &userApiImpl{
			eventDispatcher: eventDispatcher,
			userEventFactory: userEventFactory,
			userService: userService,
		}
		assert.Nil(t, userApi.ChangeUserIdentity(userId, data))
	})

	t.Run("ChangeUserIdentity:GetUserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		userId := new(models.UserId)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(nil, systemErr)

		userApi := &userApiImpl{userService: userService}
		err := userApi.ChangeUserIdentity(userId, nil)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserPassword", func(t *testing.T) {
		userId := new(models.UserId)
		password := "pass0"
		userEntity := &entities.UserEntity{Password: password}
		data := &models.UserChangePassword{OldPassword: password, NewPassword: ""}

		userEvent := new(events.UserEvent)
		userEventFactory := mocks.NewMockUserEventFactory(ctrl)
		userEventFactory.EXPECT().CreateUserPasswordChangedEvent(userEntity).Return(userEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(userEvent)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)
		userService.EXPECT().ChallengeUser(userEntity, data.OldPassword).Return(nil)
		userService.EXPECT().ChangeUserPassword(userEntity, data).Return(nil)

		userApi := &userApiImpl{
			eventDispatcher: eventDispatcher,
			userEventFactory: userEventFactory,
			userService: userService,
		}
		assert.Nil(t, userApi.ChangeUserPassword(userId, data))
	})

	t.Run("ChangeUserPassword:NoPassword", func(t *testing.T) {
		userId := new(models.UserId)
		data := &models.UserChangePassword{}
		userEntity := new(entities.UserEntity)

		userEvent := new(events.UserEvent)
		userEventFactory := mocks.NewMockUserEventFactory(ctrl)
		userEventFactory.EXPECT().CreateUserPasswordChangedEvent(userEntity).Return(userEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(userEvent)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)
		userService.EXPECT().ChangeUserPassword(userEntity, data).Return(nil)

		userApi := &userApiImpl{
			eventDispatcher: eventDispatcher,
			userEventFactory: userEventFactory,
			userService: userService,
		}
		assert.Nil(t, userApi.ChangeUserPassword(userId, data))
	})

	t.Run("ChangeUserPassword:GetUserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		userId := new(models.UserId)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(nil, systemErr)

		userApi := &userApiImpl{userService: userService}
		err := userApi.ChangeUserPassword(userId, nil)
		assert.Equal(t, systemErr, err)
	})

	t.Run("ChangeUserPassword:CheckError", func(t *testing.T) {
		systemErr := common.NewUnknownError()

		userId := new(models.UserId)
		password := "pass0"
		userEntity := &entities.UserEntity{Password: password}
		data := &models.UserChangePassword{OldPassword: password, NewPassword: ""}

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)
		userService.EXPECT().ChallengeUser(userEntity, data.OldPassword).Return(systemErr)

		userApi := &userApiImpl{userService: userService}
		err := userApi.ChangeUserPassword(userId, data)
		assert.Equal(t, systemErr, err)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)

		userEvent := new(events.UserEvent)
		userEventFactory := mocks.NewMockUserEventFactory(ctrl)
		userEventFactory.EXPECT().CreateUserDeletedEvent(userEntity).Return(userEvent)

		eventDispatcher := mocks.NewMockEventDispatcher(ctrl)
		eventDispatcher.EXPECT().Dispatch(userEvent)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(userEntity, nil)
		userService.EXPECT().DeleteUser(userEntity).Return(nil)

		userApi := &userApiImpl{
			eventDispatcher: eventDispatcher,
			userEventFactory: userEventFactory,
			userService: userService,
		}
		assert.Nil(t, userApi.DeleteUser(userId))
	})

	t.Run("DeleteUser:GetUserError", func(t *testing.T) {
		systemErr := common.NewUnknownError()
		userId := new(models.UserId)

		userService := mocks.NewMockUserService(ctrl)
		userService.EXPECT().GetUser(userId).Return(nil, systemErr)

		userApi := &userApiImpl{userService: userService}
		err := userApi.DeleteUser(userId)
		assert.Equal(t, systemErr, err)
	})
}
