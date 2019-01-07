package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
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

func (s *userServiceImpl) ListUsers(
	userPaginationQuery *models.UserPaginationQuery,
) (*models.PaginationResult, errors.Error) {
	return s.userRepository.ListUsers(userPaginationQuery)
}

func (s *userServiceImpl) CreateUser(data *models.UserCreate) (*entities.UserEntity, errors.Error) {
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

func (s *userServiceImpl) GetUser(userId *models.UserId) (*entities.UserEntity, errors.Error) {
	return s.userRepository.GetUser(userId)
}

func (s *userServiceImpl) LookupUser(entityIdentity string) (*entities.UserEntity, errors.Error) {
	return s.userRepository.LookupUser(entityIdentity)
}

func (s *userServiceImpl) ChallengeUser(userEntity *entities.UserEntity, password string) errors.Error {
	if "" == userEntity.Password {
		return errors.NewSystemError("UserEntity.Password is empty")
	}

	return s.hasher.Check(userEntity.Password, password)
}

func (s *userServiceImpl) UpdateUser(userEntity *entities.UserEntity, data *models.UserUpdate) errors.Error {
	userEntity.FirstName = data.FirstName
	userEntity.LastName = data.LastName

	updated := time.Now().UTC()
	userEntity.Updated = &updated

	return s.userRepository.SaveUser(userEntity)
}

func (s *userServiceImpl) VerifyUser(entity *entities.UserEntity) errors.Error {
	entity.Verified = true

	return s.userRepository.SaveUser(entity)
}

func (s *userServiceImpl) ChangeUserIdentity(userEntity *entities.UserEntity, data *models.UserChangeIdentity) errors.Error {
	userEntity.Email = data.Email

	return s.userRepository.SaveUser(userEntity)
}

func (s *userServiceImpl) ChangeUserPassword(userEntity *entities.UserEntity, data *models.UserChangePassword) errors.Error {
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

func (s *userServiceImpl) DeleteUser(userEntity *entities.UserEntity) errors.Error {
	return s.userRepository.RemoveUser(userEntity)
}
