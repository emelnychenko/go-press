package user

import (
	"../common"
	"../user_contract"
	"github.com/jinzhu/gorm"
)

type userContainer struct {
	UserRepository user_contract.UserRepository
	UserService    user_contract.UserService
	UserApi        user_contract.UserApi
}

func NewUserContainer(db *gorm.DB) *userContainer {
	userRepository := NewUserRepository(db)
	hasher := common.NewBCryptHasher()
	userService := NewUserService(hasher, userRepository)
	userAggregator := NewUserAggregator()
	userApi := NewUserApi(userService, userAggregator)

	return &userContainer{
		UserRepository: userRepository,
		UserService:    userService,
		UserApi:        userApi,
	}
}
