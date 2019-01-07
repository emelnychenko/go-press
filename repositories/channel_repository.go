package repositories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"github.com/jinzhu/gorm"
)

type (
	channelRepositoryImpl struct {
		db          *gorm.DB
		dbPaginator contracts.DbPaginator
	}
)

func NewChannelRepository(db *gorm.DB, dbPaginator contracts.DbPaginator) contracts.ChannelRepository {
	return &channelRepositoryImpl{db, dbPaginator}
}

func (r *channelRepositoryImpl) ListChannels(
	channelPaginationQuery *models.ChannelPaginationQuery,
) (paginationResult *models.PaginationResult, err errors.Error) {
	paginationTotal, channelEntities := 0, make([]*entities.ChannelEntity, channelPaginationQuery.Limit)
	db := r.db.Model(&channelEntities).Order("created desc")
	err = r.dbPaginator.Paginate(db, channelPaginationQuery.PaginationQuery, &channelEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: channelEntities}
	return
}

func (r *channelRepositoryImpl) GetChannel(channelId *models.ChannelId) (channelEntity *entities.ChannelEntity, err errors.Error) {
	channelEntity = new(entities.ChannelEntity)

	if gormErr := r.db.First(channelEntity, "id = ?", channelId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewChannelByIdNotFoundError(channelId)
		} else {
			err = errors.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

func (r *channelRepositoryImpl) SaveChannel(channelEntity *entities.ChannelEntity) (err errors.Error) {
	if gormErr := r.db.Save(channelEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

func (r *channelRepositoryImpl) RemoveChannel(channelEntity *entities.ChannelEntity) (err errors.Error) {
	if gormErr := r.db.Delete(channelEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
