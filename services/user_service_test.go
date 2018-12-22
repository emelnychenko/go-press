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

	testErr := common.ServerError("err0")

	t.Run("NewUserService", func(t *testing.T) {
		hasher := mocks.NewMockHasher(ctrl)
		userEntityFactory := mocks.NewMockUserEntityFactory(ctrl)
		userRepository := mocks.NewMockUserRepository(ctrl)

		userService, isUserService := NewUserService(hasher, userEntityFactory, userRepository).(*userServiceImpl)

		assert.True(t, isUserService)
		assert.Equal(t, hasher, userService.hasher)
		assert.Equal(t, userEntityFactory, userService.userEntityFactory)
		assert.Equal(t, userRepository, userService.userRepository)
	})

	t.Run("ListUsers", func(t *testing.T) {
		var userEntities []*entities.UserEntity
		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().ListUsers().Return(userEntities, nil)

		userService := &userServiceImpl{userRepository: userRepository}
		response, err := userService.ListUsers()

		assert.Equal(t, userEntities, response)
		assert.Nil(t, err)
	})

	t.Run("CreateUser:WithoutPassword", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEntityFactory := mocks.NewMockUserEntityFactory(ctrl)
		userEntityFactory.EXPECT().CreateUserEntity().Return(userEntity)

		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().SaveUser(userEntity).Return(nil)

		data := &models.UserCreate{FirstName: "foo", LastName: "bar", Email: "foo@bar"}
		userService := &userServiceImpl{userEntityFactory: userEntityFactory, userRepository: userRepository}
		response, err := userService.CreateUser(data)

		assert.Equal(t, userEntity, response)
		assert.Nil(t, err)
		assert.Equal(t, data.FirstName, userEntity.FirstName)
		assert.Equal(t, data.LastName, userEntity.LastName)
		assert.Equal(t, data.Email, userEntity.Email)
		assert.Empty(t, userEntity.Password)
	})

	t.Run("CreateUser:SaveError", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userEntityFactory := mocks.NewMockUserEntityFactory(ctrl)
		userEntityFactory.EXPECT().CreateUserEntity().Return(userEntity)

		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().SaveUser(userEntity).Return(testErr)

		data := new(models.UserCreate)
		userService := &userServiceImpl{userEntityFactory: userEntityFactory, userRepository: userRepository}
		response, err := userService.CreateUser(data)

		assert.Nil(t, response)
		assert.Error(t, err)
	})

	t.Run("CreateUser:WithPassword", func(t *testing.T) {
		password := "pass0";
		hashedPassword := "<hash>pass0"
		hasher := mocks.NewMockHasher(ctrl)
		hasher.EXPECT().Make(password).Return(hashedPassword, nil)

		userEntity := new(entities.UserEntity)
		userEntityFactory := mocks.NewMockUserEntityFactory(ctrl)
		userEntityFactory.EXPECT().CreateUserEntity().Return(userEntity)

		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().SaveUser(userEntity).Return(nil)

		data := &models.UserCreate{Password: password}
		userService := &userServiceImpl{
			hasher:            hasher,
			userEntityFactory: userEntityFactory,
			userRepository:    userRepository,
		}
		response, _ := userService.CreateUser(data)

		assert.Equal(t, hashedPassword, response.Password)
	})

	t.Run("CreateUser:HasherError", func(t *testing.T) {
		password := "pass0"
		hasher := mocks.NewMockHasher(ctrl)
		hasher.EXPECT().Make(password).Return("", testErr)

		userEntity := new(entities.UserEntity)
		userEntityFactory := mocks.NewMockUserEntityFactory(ctrl)
		userEntityFactory.EXPECT().CreateUserEntity().Return(userEntity)

		userService := &userServiceImpl{hasher: hasher, userEntityFactory: userEntityFactory}
		data := &models.UserCreate{Password: password}
		_, err := userService.CreateUser(data)

		assert.Error(t, err)
	})

	t.Run("GetUser", func(t *testing.T) {
		userId := new(models.UserId)
		userEntity := new(entities.UserEntity)
		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().GetUser(userId).Return(userEntity, nil)

		userService := &userServiceImpl{userRepository: userRepository}
		response, err := userService.GetUser(userId)

		assert.Equal(t, userEntity, response)
		assert.Nil(t, err)
	})

	t.Run("LookupUser", func(t *testing.T) {
		userIdentity := "u_identity";
		userEntity := new(entities.UserEntity)
		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().LookupUser(userIdentity).Return(userEntity, nil)

		userService := &userServiceImpl{userRepository: userRepository}
		response, err := userService.LookupUser(userIdentity)

		assert.Equal(t, userEntity, response)
		assert.Nil(t, err)
	})

	t.Run("ChallengeUser", func(t *testing.T) {
		password := "pass0";
		hashedPassword := "<hash>pass0"
		hasher := mocks.NewMockHasher(ctrl)
		hasher.EXPECT().Check(hashedPassword, password).Return(nil)

		userEntity := &entities.UserEntity{Password: hashedPassword}
		userService := &userServiceImpl{hasher: hasher}
		assert.Nil(t, userService.ChallengeUser(userEntity, password))
	})

	t.Run("ChallengeUser:NoPasswordError", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userService := &userServiceImpl{}
		assert.Error(t, userService.ChallengeUser(userEntity, ""))
	})

	t.Run("UpdateUser", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().SaveUser(userEntity).Return(nil)

		data := &models.UserUpdate{FirstName: "foo", LastName: "bar"}
		userService := &userServiceImpl{userRepository: userRepository}
		assert.Nil(t, userService.UpdateUser(userEntity, data))
		assert.Equal(t, data.FirstName, userEntity.FirstName)
		assert.Equal(t, data.LastName, userEntity.LastName)
		assert.NotNil(t, userEntity.Updated)
	})

	t.Run("VerifyUser", func(t *testing.T) {
		userEntity := &entities.UserEntity{Verified: false}
		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().SaveUser(userEntity).Return(nil)

		userService := &userServiceImpl{userRepository: userRepository}
		assert.Nil(t, userService.VerifyUser(userEntity))
		assert.True(t, userEntity.Verified)
	})

	t.Run("ChangeUserIdentity", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().SaveUser(userEntity).Return(nil)

		data := &models.UserChangeIdentity{Email: "foo@bar"}
		userService := &userServiceImpl{userRepository: userRepository}
		assert.Nil(t, userService.ChangeUserIdentity(userEntity, data))
		assert.Equal(t, data.Email, userEntity.Email)
	})

	t.Run("ChangeUserPassword:NoPassword", func(t *testing.T) {
		userEntity := &entities.UserEntity{Password: "pass0"}
		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().SaveUser(userEntity).Return(nil)

		data := new(models.UserChangePassword)
		userService := &userServiceImpl{userRepository: userRepository}
		assert.Nil(t, userService.ChangeUserPassword(userEntity, data))
		assert.Empty(t, userEntity.Password)
	})

	t.Run("ChangeUserPassword", func(t *testing.T) {
		password := "pass0";
		hashedPassword := "<hash>pass0"
		hasher := mocks.NewMockHasher(ctrl)
		hasher.EXPECT().Make(password).Return(hashedPassword, nil)

		userEntity := new(entities.UserEntity)
		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().SaveUser(userEntity).Return(nil)

		data := &models.UserChangePassword{NewPassword: password}
		userService := &userServiceImpl{hasher: hasher, userRepository: userRepository}
		assert.Nil(t, userService.ChangeUserPassword(userEntity, data))
		assert.Equal(t, hashedPassword, userEntity.Password)
	})

	t.Run("ChangeUserPassword:Error", func(t *testing.T) {
		password := "pass0"
		hasher := mocks.NewMockHasher(ctrl)
		hasher.EXPECT().Make(password).Return("", testErr)

		userEntity := new(entities.UserEntity)
		data := &models.UserChangePassword{NewPassword: password}
		userService := &userServiceImpl{hasher: hasher}
		assert.Error(t, userService.ChangeUserPassword(userEntity, data))
	})

	t.Run("DeleteUser", func(t *testing.T) {
		userEntity := new(entities.UserEntity)
		userRepository := mocks.NewMockUserRepository(ctrl)
		userRepository.EXPECT().RemoveUser(userEntity).Return(nil)

		userService := &userServiceImpl{userRepository: userRepository}
		err := userService.DeleteUser(userEntity)

		assert.Nil(t, err)
	})
}
