package repositories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"github.com/jinzhu/gorm"
)

type (
	commentRepositoryImpl struct {
		db          *gorm.DB
		dbPaginator contracts.DbPaginator
	}
)

func NewCommentRepository(db *gorm.DB, dbPaginator contracts.DbPaginator) contracts.CommentRepository {
	return &commentRepositoryImpl{db, dbPaginator}
}

func (r *commentRepositoryImpl) ListComments(
	commentPaginationQuery *models.CommentPaginationQuery,
) (paginationResult *models.PaginationResult, err errors.Error) {
	paginationTotal, commentEntities := 0, make([]*entities.CommentEntity, commentPaginationQuery.Limit)
	db := r.db.Model(&commentEntities).Order("created desc")
	err = r.dbPaginator.Paginate(db, commentPaginationQuery.PaginationQuery, &commentEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: commentEntities}
	return
}

func (r *commentRepositoryImpl) GetComment(commentId *models.CommentId) (
	commentEntity *entities.CommentEntity, err errors.Error,
) {
	commentEntity = new(entities.CommentEntity)

	if gormErr := r.db.First(commentEntity, "id = ?", commentId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewCommentByIdNotFoundError(commentId)
		} else {
			err = errors.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

func (r *commentRepositoryImpl) SaveComment(commentEntity *entities.CommentEntity) (err errors.Error) {
	if gormErr := r.db.Save(commentEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

func (r *commentRepositoryImpl) RemoveComment(commentEntity *entities.CommentEntity) (err errors.Error) {
	if gormErr := r.db.Delete(commentEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
