package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
)

type (
	userPictureServiceImpl struct {
		userRepository contracts.UserRepository
	}
)

func NewUserPictureService(userRepository contracts.UserRepository) contracts.UserPictureService {
	return &userPictureServiceImpl{userRepository: userRepository}
}

func (s *userPictureServiceImpl) ChangeUserPicture(
	userEntity *entities.UserEntity, userPictureEntity *entities.FileEntity,
) errors.Error {
	userEntity.SetPicture(userPictureEntity)

	return s.userRepository.SaveUser(userEntity)
}

func (s *userPictureServiceImpl) RemoveUserPicture(userEntity *entities.UserEntity) errors.Error {
	userEntity.RemovePicture()

	return s.userRepository.SaveUser(userEntity)
}
