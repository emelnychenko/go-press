package user

import (
	"../common"
	"../common_contract"
	"../user_contract"
	"../user_domain"
)

type (
	userServiceImpl struct {
		hasher     common_contract.Hasher
		repository user_contract.UserRepository
	}
)

func NewUserService(hasher common_contract.Hasher, userRepository user_contract.UserRepository) *userServiceImpl {
	return &userServiceImpl{hasher: hasher, repository: userRepository}
}

func (c *userServiceImpl) ListUsers() ([]*user_domain.UserEntity, common.Error) {
	return c.repository.ListUsers()
}

func (c *userServiceImpl) CreateUser(data *user_domain.UserCreate) (*user_domain.UserEntity, common.Error) {
	userEntity := user_domain.NewUserEntity()
	userEntity.FirstName = data.FirstName
	userEntity.LastName = data.LastName
	userEntity.Email = data.Email

	if data.Password != "" {
		hash, err := c.hasher.Make(data.Password)

		if nil != err {
			return nil, err
		}

		userEntity.Password = hash
	}

	if err := c.repository.SaveUser(userEntity); err != nil {
		return nil, err
	}

	return userEntity, nil
}

func (c *userServiceImpl) GetUser(userId *user_domain.UserId) (*user_domain.UserEntity, common.Error) {
	return c.repository.GetUser(userId)
}

func (c *userServiceImpl) LookupUser(entityIdentity string) (*user_domain.UserEntity, common.Error) {
	return c.repository.LookupUser(entityIdentity)
}

func (c *userServiceImpl) ChallengeUser(userEntity *user_domain.UserEntity, password string) common.Error {
	if userEntity.Password == "" {
		return common.ServerError("password is not registered")
	}

	return c.hasher.Check(userEntity.Password, password)
}

func (c *userServiceImpl) UpdateUser(userEntity *user_domain.UserEntity, data *user_domain.UserUpdate) common.Error {
	userEntity.FirstName = data.FirstName
	userEntity.LastName = data.LastName

	return c.repository.SaveUser(userEntity)
}

func (c *userServiceImpl) VerifyUser(entity *user_domain.UserEntity) common.Error {
	entity.Verified = true

	return c.repository.SaveUser(entity)
}

func (c *userServiceImpl) ChangeUserIdentity(userEntity *user_domain.UserEntity, data *user_domain.UserChangeIdentity) common.Error {
	userEntity.Email = data.Email

	return c.repository.SaveUser(userEntity)
}

func (c *userServiceImpl) ChangeUserPassword(userEntity *user_domain.UserEntity, data *user_domain.UserChangePassword) common.Error {
	if data.NewPassword != "" {
		hash, err := c.hasher.Make(data.NewPassword)

		if nil != err {
			return err
		}

		userEntity.Password = string(hash)
	} else {
		userEntity.Password = ""
	}

	return c.repository.SaveUser(userEntity)
}

func (c *userServiceImpl) DeleteUser(userEntity *user_domain.UserEntity) common.Error {
	return c.repository.RemoveUser(userEntity)
}
