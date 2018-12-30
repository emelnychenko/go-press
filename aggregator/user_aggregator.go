package aggregator

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/models"
)

type userAggregatorImpl struct {
	userModelFactory contracts.UserModelFactory
	fileApi          contracts.FileApi
}

func NewUserAggregator(userModelFactory contracts.UserModelFactory, fileApi contracts.FileApi) contracts.UserAggregator {
	return &userAggregatorImpl{userModelFactory, fileApi}
}

func (a *userAggregatorImpl) AggregateUser(userEntity *entities.UserEntity) *models.User {
	user := a.userModelFactory.CreateUser()
	user.Id = userEntity.Id
	user.FirstName = userEntity.FirstName
	user.LastName = userEntity.LastName
	user.Email = userEntity.Email
	user.Verified = userEntity.Verified

	if nil != userEntity.PictureId {
		user.Picture, _ = a.fileApi.GetFile(userEntity.PictureId)
	}

	return user
}

func (a *userAggregatorImpl) AggregateUsers(userEntities []*entities.UserEntity) []*models.User {
	users := make([]*models.User, len(userEntities))

	for k, v := range userEntities {
		users[k] = a.AggregateUser(v)
	}

	return users
}

func (a *userAggregatorImpl) AggregatePaginationResult(
	entityPaginationResult *models.PaginationResult,
) (
	paginationResult *models.PaginationResult,
) {
	userEntities := entityPaginationResult.Data.([]*entities.UserEntity)
	users := a.AggregateUsers(userEntities)
	return &models.PaginationResult{Total: entityPaginationResult.Total, Data: users}
}
