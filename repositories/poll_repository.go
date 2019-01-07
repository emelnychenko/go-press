package repositories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"github.com/jinzhu/gorm"
)

type (
	pollRepositoryImpl struct {
		db          *gorm.DB
		dbPaginator contracts.DbPaginator
	}
)

func NewPollRepository(db *gorm.DB, dbPaginator contracts.DbPaginator) contracts.PollRepository {
	return &pollRepositoryImpl{db, dbPaginator}
}

func (r *pollRepositoryImpl) ListPolls(
	pollPaginationQuery *models.PollPaginationQuery,
) (paginationResult *models.PaginationResult, err errors.Error) {
	paginationTotal, pollEntities := 0, make([]*entities.PollEntity, pollPaginationQuery.Limit)
	db := r.db.Model(&pollEntities).Order("created desc")
	err = r.dbPaginator.Paginate(db, pollPaginationQuery.PaginationQuery, &pollEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: pollEntities}
	return
}

func (r *pollRepositoryImpl) GetPoll(pollId *models.PollId) (pollEntity *entities.PollEntity, err errors.Error) {
	pollEntity = new(entities.PollEntity)

	if gormErr := r.db.First(pollEntity, "id = ?", pollId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewPollByIdNotFoundError(pollId)
		} else {
			err = errors.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

func (r *pollRepositoryImpl) SavePoll(pollEntity *entities.PollEntity) (err errors.Error) {
	if gormErr := r.db.Save(pollEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

func (r *pollRepositoryImpl) RemovePoll(pollEntity *entities.PollEntity) (err errors.Error) {
	if gormErr := r.db.Delete(pollEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
