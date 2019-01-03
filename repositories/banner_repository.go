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
	bannerRepositoryImpl struct {
		db          *gorm.DB
		dbPaginator contracts.DbPaginator
	}
)

func NewBannerRepository(db *gorm.DB, dbPaginator contracts.DbPaginator) contracts.BannerRepository {
	return &bannerRepositoryImpl{db, dbPaginator}
}

func (r *bannerRepositoryImpl) ListBanners(
	bannerPaginationQuery *models.BannerPaginationQuery,
) (paginationResult *models.PaginationResult, err common.Error) {
	paginationTotal, bannerEntities := 0, make([]*entities.BannerEntity, bannerPaginationQuery.Limit)
	db := r.db.Model(&bannerEntities).Order("created desc")
	err = r.dbPaginator.Paginate(db, bannerPaginationQuery.PaginationQuery, &bannerEntities, &paginationTotal)

	if nil != err {
		return
	}

	paginationResult = &models.PaginationResult{Total: paginationTotal, Data: bannerEntities}
	return
}

func (r *bannerRepositoryImpl) GetBanner(bannerId *models.BannerId) (
	bannerEntity *entities.BannerEntity, err common.Error,
) {
	bannerEntity = new(entities.BannerEntity)

	if gormErr := r.db.First(bannerEntity, "id = ?", bannerId).Error; gormErr != nil {
		if gorm.IsRecordNotFoundError(gormErr) {
			err = errors.NewBannerByIdNotFoundError(bannerId)
		} else {
			err = common.NewSystemErrorFromBuiltin(gormErr)
		}
	}

	return
}

func (r *bannerRepositoryImpl) SaveBanner(bannerEntity *entities.BannerEntity) (err common.Error) {
	if gormErr := r.db.Save(bannerEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}

func (r *bannerRepositoryImpl) RemoveBanner(bannerEntity *entities.BannerEntity) (err common.Error) {
	if gormErr := r.db.Delete(bannerEntity).Error; gormErr != nil {
		err = common.NewSystemErrorFromBuiltin(gormErr)
	}

	return
}
