package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
)

type (
	userPictureServiceImpl struct {
		userRepository contracts.UserRepository
	}
)

func NewUserPictureService(userRepository contracts.UserRepository) contracts.UserPictureService {
	return &userPictureServiceImpl{userRepository: userRepository}
}

func (s *userPictureServiceImpl) ChangeUserPicture(userEntity *entities.UserEntity, userPicture *entities.FileEntity) common.Error {
	userEntity.SetPicture(userPicture)

	return s.userRepository.SaveUser(userEntity)
}

func (s *userPictureServiceImpl) RemoveUserPicture(userEntity *entities.UserEntity) common.Error {
	userEntity.RemovePicture()

	return s.userRepository.SaveUser(userEntity)
}
