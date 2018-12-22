package services

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	userServiceImpl struct {
		hasher            contracts.Hasher
		userEntityFactory contracts.UserEntityFactory
		userRepository    contracts.UserRepository
	}
)

func NewUserService(
	hasher contracts.Hasher,
	userEntityFactory contracts.UserEntityFactory,
	userRepository contracts.UserRepository,
) contracts.UserService {
	return &userServiceImpl{hasher: hasher, userEntityFactory: userEntityFactory, userRepository: userRepository}
}

func (s *userServiceImpl) ListUsers() ([]*entities.UserEntity, common.Error) {
	return s.userRepository.ListUsers()
}

func (s *userServiceImpl) CreateUser(data *models.UserCreate) (*entities.UserEntity, common.Error) {
	userEntity := s.userEntityFactory.CreateUserEntity()
	userEntity.FirstName = data.FirstName
	userEntity.LastName = data.LastName
	userEntity.Email = data.Email

	if data.Password != "" {
		hash, err := s.hasher.Make(data.Password)

		if nil != err {
			return nil, err
		}

		userEntity.Password = hash
	}

	if err := s.userRepository.SaveUser(userEntity); err != nil {
		return nil, err
	}

	return userEntity, nil
}

func (s *userServiceImpl) GetUser(userId *models.UserId) (*entities.UserEntity, common.Error) {
	return s.userRepository.GetUser(userId)
}

func (s *userServiceImpl) LookupUser(entityIdentity string) (*entities.UserEntity, common.Error) {
	return s.userRepository.LookupUser(entityIdentity)
}

func (s *userServiceImpl) ChallengeUser(userEntity *entities.UserEntity, password string) common.Error {
	if userEntity.Password == "" {
		return common.ServerError("password is not registered")
	}

	return s.hasher.Check(userEntity.Password, password)
}

func (s *userServiceImpl) UpdateUser(userEntity *entities.UserEntity, data *models.UserUpdate) common.Error {
	userEntity.FirstName = data.FirstName
	userEntity.LastName = data.LastName

	updated := time.Now().UTC()
	userEntity.Updated = &updated

	return s.userRepository.SaveUser(userEntity)
}

func (s *userServiceImpl) VerifyUser(entity *entities.UserEntity) common.Error {
	entity.Verified = true

	return s.userRepository.SaveUser(entity)
}

func (s *userServiceImpl) ChangeUserIdentity(userEntity *entities.UserEntity, data *models.UserChangeIdentity) common.Error {
	userEntity.Email = data.Email

	return s.userRepository.SaveUser(userEntity)
}

func (s *userServiceImpl) ChangeUserPassword(userEntity *entities.UserEntity, data *models.UserChangePassword) common.Error {
	if data.NewPassword != "" {
		hash, err := s.hasher.Make(data.NewPassword)

		if nil != err {
			return err
		}

		userEntity.Password = string(hash)
	} else {
		userEntity.Password = ""
	}

	return s.userRepository.SaveUser(userEntity)
}

func (s *userServiceImpl) DeleteUser(userEntity *entities.UserEntity) common.Error {
	return s.userRepository.RemoveUser(userEntity)
}
