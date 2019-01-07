package repositories

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/enums"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"github.com/jinzhu/gorm"
	"time"
)

type (
	postRepositoryImpl struct {
		db          *gorm.DB
		dbPaginator contracts.DbPaginator
	}
)

func NewPostRepository(db *gorm.DB, dbPaginator contracts.DbPaginator) contracts.PostRepository {
	return &postRepositoryImpl{db, dbPaginator}
}

func (r *postRepositoryImpl) ListPosts(
	postPaginationQuery *models.PostPaginationQuery,
) (paginationResult *models.PaginationResult, err errors.Error) {
	paginationTotal, postEntities := 0, make([]*entities.PostEntity, postPaginationQuery.Limit)
	db := r.db.Model(&postEntities).Order("created desc")

	if "" != postPaginationQuery.Status {
		db = db.Where("status = ?", postPaginationQuery.Status)
	}

	if "" != postPaginationQuery.Privacy {
		db = db.Where("privacy = ?", postPaginationQuery.Privacy)
	}

	if "" != postPaginationQuery.Author {
		db = db.Where("author_id = ?", postPaginationQuery.Author)
	}

	err = r.dbPaginator.Paginate(db, postPaginationQuery.PaginationQuery, &postEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: postEntities}
	return
}

func (r *postRepositoryImpl) GetScheduledPosts() (postEntities []*entities.PostEntity, err errors.Error) {
	postEntities = []*entities.PostEntity{}

	gormErr := r.db.Where("status = ?", enums.PostScheduledStatus).
		Where("published < ?", time.Now().UTC().Format("2006-01-02 15:04:05")).
		Order("published desc").
		Limit(100).
		Find(&postEntities).Error

	if gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

func (r *postRepositoryImpl) GetPost(postId *models.PostId) (postEntity *entities.PostEntity, err errors.Error) {
	postEntity = new(entities.PostEntity)

	if gormErr := r.db.First(postEntity, "id = ?", postId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewPostByIdNotFoundError(postId)
		} else {
			err = errors.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

func (r *postRepositoryImpl) SavePost(postEntity *entities.PostEntity) (err errors.Error) {
	if gormErr := r.db.Save(postEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

func (r *postRepositoryImpl) RemovePost(postEntity *entities.PostEntity) (err errors.Error) {
	if gormErr := r.db.Delete(postEntity).Error; gormErr != nil {
		err = errors.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
