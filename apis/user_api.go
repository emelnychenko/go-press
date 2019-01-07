package apis

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
)

type (
	userApiImpl struct {
		eventDispatcher  contracts.EventDispatcher
		userEventFactory contracts.UserEventFactory
		userService      contracts.UserService
		userAggregator   contracts.UserAggregator
	}
)

func NewUserApi(
	eventDispatcher contracts.EventDispatcher,
	userEventFactory contracts.UserEventFactory,
	userService contracts.UserService,
	userAggregator contracts.UserAggregator,
) (userApi contracts.UserApi) {
	return &userApiImpl{
		eventDispatcher,
		userEventFactory,
		userService,
		userAggregator,
	}
}

func (a *userApiImpl) ListUsers(
	userPaginationQuery *models.UserPaginationQuery,
) (paginationResult *models.PaginationResult, err errors.Error) {
	entityPaginationResult, err := a.userService.ListUsers(userPaginationQuery)

	if nil != err {
		return
	}

	paginationResult = a.userAggregator.AggregatePaginationResult(entityPaginationResult)
	return
}

func (a *userApiImpl) GetUser(userId *models.UserId) (user *models.User, err errors.Error) {
	userEntity, err := a.userService.GetUser(userId)

	if nil != err {
		return
	}

	user = a.userAggregator.AggregateUser(userEntity)
	return
}

func (a *userApiImpl) CreateUser(data *models.UserCreate) (user *models.User, err errors.Error) {
	userEntity, err := a.userService.CreateUser(data)

	if nil != err {
		return
	}

	userEvent := a.userEventFactory.CreateUserCreatedEvent(userEntity)
	a.eventDispatcher.Dispatch(userEvent)

	user = a.userAggregator.AggregateUser(userEntity)
	return
}

func (a *userApiImpl) UpdateUser(userId *models.UserId, data *models.UserUpdate) (err errors.Error) {
	userService := a.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return
	}

	err = userService.UpdateUser(userEntity, data)

	if nil != err {
		return
	}

	userEvent := a.userEventFactory.CreateUserUpdatedEvent(userEntity)
	a.eventDispatcher.Dispatch(userEvent)

	return
}

func (a *userApiImpl) VerifyUser(userId *models.UserId) (err errors.Error) {
	userService := a.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return
	}

	err = userService.VerifyUser(userEntity)

	if nil != err {
		return
	}

	userEvent := a.userEventFactory.CreateUserVerifiedEvent(userEntity)
	a.eventDispatcher.Dispatch(userEvent)

	return
}

func (a *userApiImpl) ChangeUserIdentity(userId *models.UserId, data *models.UserChangeIdentity) (err errors.Error) {
	userService := a.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return
	}

	err = userService.ChangeUserIdentity(userEntity, data)

	if nil != err {
		return
	}

	userEvent := a.userEventFactory.CreateUserIdentityChangedEvent(userEntity)
	a.eventDispatcher.Dispatch(userEvent)

	return
}

func (a *userApiImpl) ChangeUserPassword(userId *models.UserId, input *models.UserChangePassword) (err errors.Error) {
	userService := a.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return
	}

	if userEntity.Password != "" {
		if err = userService.ChallengeUser(userEntity, input.OldPassword); err != nil {
			return
		}
	}

	err = userService.ChangeUserPassword(userEntity, input)

	if nil != err {
		return
	}

	userEvent := a.userEventFactory.CreateUserPasswordChangedEvent(userEntity)
	a.eventDispatcher.Dispatch(userEvent)

	return
}

func (a *userApiImpl) DeleteUser(userId *models.UserId) (err errors.Error) {
	userService := a.userService
	userEntity, err := userService.GetUser(userId)

	if nil != err {
		return
	}

	err = userService.DeleteUser(userEntity)

	if nil != err {
		return
	}

	userEvent := a.userEventFactory.CreateUserDeletedEvent(userEntity)
	a.eventDispatcher.Dispatch(userEvent)

	return
}
