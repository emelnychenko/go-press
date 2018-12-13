package user

import (
	"../common"
	"../user_domain"
	"../user_mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userId := new(user_domain.UserId)
	testPass := "pass0";
	testErr := common.ServerError("err0");

	t.Run("ListUsers", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userAggregator := user_mock.NewMockUserAggregator(ctrl)
		userApi := NewUserApi(userRepository, userAggregator)
		var userEntities []*user_domain.UserEntity

		userRepository.EXPECT().ListUsers().Return(userEntities, nil)
		userAggregator.EXPECT().AggregateCollection(userEntities).Return([]*user_domain.User{})
		models, err := userApi.ListUsers()

		assert.IsType(t, []*user_domain.User{}, models)
		assert.Nil(t, err)
	})

	t.Run("ListUsers: Error", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)

		userRepository.EXPECT().ListUsers().Return(nil, testErr)
		models, err := userApi.ListUsers()

		assert.Nil(t, models)
		assert.Error(t, err)
	})

	t.Run("CreateUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userAggregator := user_mock.NewMockUserAggregator(ctrl)
		userApi := NewUserApi(userRepository, userAggregator)
		userEntity := &user_domain.UserEntity{}
		data := &user_domain.UserCreate{}

		userRepository.EXPECT().CreateUser(data).Return(userEntity, nil)
		userAggregator.EXPECT().AggregateObject(userEntity).Return(&user_domain.User{})
		model, err := userApi.CreateUser(data)

		assert.IsType(t, &user_domain.User{}, model)
		assert.Nil(t, err)
	})

	t.Run("CreateUser:Error", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)
		data := &user_domain.UserCreate{}

		userRepository.EXPECT().CreateUser(data).Return(nil, testErr)
		model, err := userApi.CreateUser(data)

		assert.Nil(t, model)
		assert.Error(t, err)
	})

	t.Run("GetUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userAggregator := user_mock.NewMockUserAggregator(ctrl)
		userApi := NewUserApi(userRepository, userAggregator)
		userEntity := &user_domain.UserEntity{}

		userRepository.EXPECT().GetUser(userId).Return(userEntity, nil)
		userAggregator.EXPECT().AggregateObject(userEntity).Return(&user_domain.User{})
		model, err := userApi.GetUser(userId)

		assert.IsType(t, &user_domain.User{}, model)
		assert.Nil(t, err)
	})

	t.Run("GetUser: NoUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)

		userRepository.EXPECT().GetUser(userId).Return(nil, testErr)
		model, err := userApi.GetUser(userId)
		assert.Nil(t, model)
		assert.Error(t, err)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)
		userEntity := &user_domain.UserEntity{}
		form := &user_domain.UserUpdate{}

		userRepository.EXPECT().GetUser(userId).Return(userEntity, nil)
		userRepository.EXPECT().UpdateUser(userEntity, form).Return(nil)
		err := userApi.UpdateUser(userId, form)
		assert.Nil(t, err)
	})

	t.Run("UpdateUser: NoUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)

		userRepository.EXPECT().GetUser(userId).Return(nil, testErr)
		err := userApi.UpdateUser(userId, nil)
		assert.Error(t, err)
	})

	t.Run("ChangeUserIduserEntity", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)
		userEntity := &user_domain.UserEntity{}
		model := &user_domain.UserChangeIdentity{}

		userRepository.EXPECT().GetUser(userId).Return(userEntity, nil)
		userRepository.EXPECT().ChangeUserIdentity(userEntity, model).Return(nil)
		err := userApi.ChangeUserIdentity(userId, model)
		assert.Nil(t, err)
	})

	t.Run("ChangeUserIduserEntity: NoUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)

		userRepository.EXPECT().GetUser(userId).Return(nil, testErr)
		err := userApi.ChangeUserIdentity(userId, nil)
		assert.Error(t, err)
	})

	t.Run("ChangeUserPassword: NoUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)

		userRepository.EXPECT().GetUser(userId).Return(nil, testErr)
		err := userApi.ChangeUserPassword(userId, nil)
		assert.Error(t, err)
	})

	t.Run("ChangeUserPassword: NoPassword", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)
		userEntity := &user_domain.UserEntity{}
		formData := &user_domain.UserChangePassword{}

		userRepository.EXPECT().GetUser(userId).Return(userEntity, nil)
		userRepository.EXPECT().ChangeUserPassword(userEntity, formData).Return(nil)
		err := userApi.ChangeUserPassword(userId, formData)

		assert.Nil(t, err)
	})

	t.Run("ChangeUserPassword: ChallengeUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)
		userEntity := &user_domain.UserEntity{Password: testPass}
		formData := &user_domain.UserChangePassword{OldPassword: testPass, NewPassword: ""}

		userRepository.EXPECT().GetUser(userId).Return(userEntity, nil)
		userRepository.EXPECT().ChallengeUser(userEntity, formData.OldPassword).Return(nil)
		userRepository.EXPECT().ChangeUserPassword(userEntity, formData).Return(nil)
		err := userApi.ChangeUserPassword(userId, formData)

		assert.Nil(t, err)
	})

	t.Run("ChangeUserPassword: FailChallenge", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)
		userEntity := &user_domain.UserEntity{Password: testPass}
		formData := &user_domain.UserChangePassword{OldPassword: testPass, NewPassword: ""}

		userRepository.EXPECT().GetUser(userId).Return(userEntity, nil)
		userRepository.EXPECT().ChallengeUser(userEntity, formData.OldPassword).Return(testErr)
		err := userApi.ChangeUserPassword(userId, formData)

		assert.Error(t, err)
	})

	t.Run("VerifyUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)
		userEntity := &user_domain.UserEntity{}

		userRepository.EXPECT().GetUser(userId).Return(userEntity, nil)
		userRepository.EXPECT().VerifyUser(userEntity).Return(nil)
		err := userApi.VerifyUser(userId)
		assert.Nil(t, err)
	})

	t.Run("VerifyUser: NoUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)

		userRepository.EXPECT().GetUser(userId).Return(nil, testErr)
		err := userApi.VerifyUser(userId)
		assert.Error(t, err)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)
		userEntity := &user_domain.UserEntity{}

		userRepository.EXPECT().GetUser(userId).Return(userEntity, nil)
		userRepository.EXPECT().DeleteUser(userEntity).Return(nil)
		err := userApi.DeleteUser(userId)
		assert.Nil(t, err)
	})

	t.Run("DeleteUser: NoUser", func(t *testing.T) {
		userRepository := user_mock.NewMockUserService(ctrl)
		userApi := NewUserApi(userRepository, nil)

		userRepository.EXPECT().GetUser(userId).Return(nil, testErr)
		err := userApi.DeleteUser(userId)
		assert.Error(t, err)
	})
}