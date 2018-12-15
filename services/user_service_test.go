package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userId := new(models.UserId)
	testErr := common.ServerError("err0")
	testPass := "pass0"

	t.Run("ListUsers()", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userService := NewUserService(nil, userRepository)

		userRepository.EXPECT().ListUsers().Return([]*entities.UserEntity{}, nil)
		results, err := userService.ListUsers()

		assert.IsType(t, []*entities.UserEntity{}, results)
		assert.Nil(t, err)
	})

	t.Run("CreateUser(UserCreate) NoPassword", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userService := NewUserService(nil, userRepository)
		data := &models.UserCreate{
			FirstName: "foo",
			LastName: "bar",
			Email: "foo@bar",
		}

		userRepository.EXPECT().SaveUser(gomock.Any()).Return(nil)
		userEntity, err := userService.CreateUser(data)

		assert.IsType(t, &entities.UserEntity{}, userEntity)
		assert.Nil(t, err)
		assert.Equal(t, data.FirstName, userEntity.FirstName)
		assert.Equal(t, data.LastName, userEntity.LastName)
		assert.Equal(t, data.Email, userEntity.Email)
	})

	t.Run("CreateUser(UserCreate) UserRepository.SaveUser() Error", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userService := NewUserService(nil, userRepository)
		data := &models.UserCreate{}

		userRepository.EXPECT().SaveUser(gomock.Any()).Return(testErr)
		userEntity, err := userService.CreateUser(data)

		assert.Nil(t, userEntity)
		assert.Error(t, err)
	})

	t.Run("CreateUser(UserCreate) Password", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		hasher := mocks.NewMockHasher(ctrl)
		userService := NewUserService(hasher, userRepository)
		data := &models.UserCreate{
			Password: testPass,
		}

		hasher.EXPECT().Make(data.Password).Return(data.Password, nil)
		userRepository.EXPECT().SaveUser(gomock.Any()).Return(nil)
		userEntity, _ := userService.CreateUser(data)

		assert.Equal(t, data.Password, userEntity.Password)
	})

	t.Run("CreateUser(UserCreate) Password Error", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		hasher := mocks.NewMockHasher(ctrl)
		userService := NewUserService(hasher, userRepository)
		data := &models.UserCreate{
			Password: testPass,
		}

		hasher.EXPECT().Make(data.Password).Return("", testErr)
		_, err := userService.CreateUser(data)

		assert.Error(t, err)
	})

	t.Run("GetUser(UUID)", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userService := NewUserService(nil, userRepository)

		userRepository.EXPECT().GetUser(userId).Return(&entities.UserEntity{}, nil)
		results, err := userService.GetUser(userId)

		assert.IsType(t, &entities.UserEntity{}, results)
		assert.Nil(t, err)
	})

	t.Run("LookupUser(string)", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userService := NewUserService(nil, userRepository)
		userIdentity := ""

		userRepository.EXPECT().LookupUser(userIdentity).Return(&entities.UserEntity{}, nil)
		results, err := userService.LookupUser(userIdentity)

		assert.IsType(t, &entities.UserEntity{}, results)
		assert.Nil(t, err)
	})

	t.Run("ChallengeUser(string) NoPassword Error", func(t *testing.T) {
		userService := NewUserService(nil, nil)

		err := userService.ChallengeUser(&entities.UserEntity{}, "")
		assert.Error(t, err)
	})

	t.Run("ChallengeUser(string) NotMatch Error", func(t *testing.T) {
		hasher := mocks.NewMockHasher(ctrl)
		userService := NewUserService(hasher, nil)

		hasher.EXPECT().Check(testPass, testPass).Return(testErr)
		err := userService.ChallengeUser(&entities.UserEntity{Password: testPass}, testPass)
		assert.Error(t, err)
	})

	t.Run("ChallengeUser(string)", func(t *testing.T) {
		hasher := mocks.NewMockHasher(ctrl)
		userService := NewUserService(hasher, nil)

		hasher.EXPECT().Check(testPass, testPass).Return(nil)
		err := userService.ChallengeUser(&entities.UserEntity{Password: testPass}, testPass)
		assert.Nil(t, err)
	})

	t.Run("UpdateUser(UserEntity,UserUpdate)", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userService := NewUserService(nil, userRepository)
		data := &models.UserUpdate{FirstName: "foo",LastName:"bar"}
		userEntity := &entities.UserEntity{}

		userRepository.EXPECT().SaveUser(userEntity).Return(nil)
		err := userService.UpdateUser(userEntity, data)

		assert.Nil(t, err)
		assert.Equal(t, data.FirstName, userEntity.FirstName)
		assert.Equal(t, data.LastName, userEntity.LastName)
	})

	t.Run("VerifyUser(UserEntity)", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userService := NewUserService(nil, userRepository)
		userEntity := &entities.UserEntity{}

		userRepository.EXPECT().SaveUser(userEntity).Return(nil)
		err := userService.VerifyUser(userEntity)

		assert.Nil(t, err)
		assert.Equal(t, true, userEntity.Verified)
	})

	t.Run("ChangeUserIdentity(UserEntity,UserChangeIdentity)", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userService := NewUserService(nil, userRepository)
		userEntity := &entities.UserEntity{}
		data := &models.UserChangeIdentity{Email: "foo@bar"}

		userRepository.EXPECT().SaveUser(userEntity).Return(nil)
		err := userService.ChangeUserIdentity(userEntity, data)

		assert.Nil(t, err)
		assert.Equal(t, data.Email, userEntity.Email)
	})

	t.Run("ChangeUserPassword(UserEntity,UserChangePassword) NoPassword", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userService := NewUserService(nil, userRepository)
		userEntity := &entities.UserEntity{}
		data := &models.UserChangePassword{NewPassword: ""}

		userRepository.EXPECT().SaveUser(userEntity).Return(nil)
		err := userService.ChangeUserPassword(userEntity, data)

		assert.Nil(t, err)
		assert.Equal(t, data.NewPassword, userEntity.Password)
	})

	t.Run("ChangeUserPassword(UserEntity,UserChangePassword)", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		hasher := mocks.NewMockHasher(ctrl)
		userService := NewUserService(hasher, userRepository)
		userEntity := &entities.UserEntity{}
		data := &models.UserChangePassword{NewPassword: testPass}

		hasher.EXPECT().Make(testPass).Return(testPass, nil)
		userRepository.EXPECT().SaveUser(userEntity).Return(nil)
		err := userService.ChangeUserPassword(userEntity, data)

		assert.Nil(t, err)
		assert.Equal(t, data.NewPassword, userEntity.Password)
	})

	t.Run("ChangeUserPassword(UserEntity,UserChangePassword) Error", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		hasher := mocks.NewMockHasher(ctrl)
		userService := NewUserService(hasher, userRepository)
		userEntity := &entities.UserEntity{}
		data := &models.UserChangePassword{NewPassword: testPass}

		hasher.EXPECT().Make(testPass).Return("", testErr)
		err := userService.ChangeUserPassword(userEntity, data)

		assert.Error(t, err)
	})

	t.Run("DeleteUser(UserEntity)", func(t *testing.T) {
		userRepository := mocks.NewMockUserRepository(ctrl)
		userService := NewUserService(nil, userRepository)
		userEntity := &entities.UserEntity{}

		userRepository.EXPECT().RemoveUser(userEntity).Return(nil)
		err := userService.DeleteUser(userEntity)

		assert.Nil(t, err)
	})
}