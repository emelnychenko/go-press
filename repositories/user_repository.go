package repositories

import (
	"github.com/emelnychenko/go-press/common"
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"github.com/jinzhu/gorm"
)

type (
	userRepositoryImpl struct {
		db          *gorm.DB
		dbPaginator contracts.DbPaginator
	}
)

func NewUserRepository(db *gorm.DB, dbPaginator contracts.DbPaginator) contracts.UserRepository {
	return &userRepositoryImpl{db, dbPaginator}
}

func (r *userRepositoryImpl) ListUsers(
	userPaginationQuery *models.UserPaginationQuery,
) (paginationResult *models.PaginationResult, err common.Error) {
	paginationTotal, userEntities := 0, make([]*entities.UserEntity, userPaginationQuery.Limit)
	db := r.db.Model(&userEntities).Order("created desc")

	err = r.dbPaginator.Paginate(db, userPaginationQuery.PaginationQuery, &userEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: userEntities}
	return
}

func (r *userRepositoryImpl) GetUser(userId *models.UserId) (userEntity *entities.UserEntity, err common.Error) {
	userEntity = new(entities.UserEntity)

	if gormErr := r.db.First(userEntity, "id = ?", userId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewUserByIdNotFoundError(userId)
		} else {
			err = common.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

func (r *userRepositoryImpl) LookupUser(userIdentity string) (userEntity *entities.UserEntity, err common.Error) {
	userEntity = new(entities.UserEntity)

	if gormErr := r.db.First(userEntity, "email = ?", userIdentity).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewUserNotFoundError(userIdentity)
		} else {
			err = common.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

func (r *userRepositoryImpl) SaveUser(userEntity *entities.UserEntity) (err common.Error) {
	if gormErr := r.db.Save(userEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

func (r *userRepositoryImpl) RemoveUser(userEntity *entities.UserEntity) (err common.Error) {
	if gormErr := r.db.Delete(userEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
