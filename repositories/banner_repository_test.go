package repositories

import (
	mocket "github.com/Selvatico/go-mocket"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/mocks"
	"github.com/emelnychenko/go-press/models"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBannerRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbPaginator := mocks.NewMockDbPaginator(ctrl)
	db, _ := gorm.Open(mocket.DriverName, "")
	bannerRepository, isBannerRepository := NewBannerRepository(db, dbPaginator).(*bannerRepositoryImpl)

	assert.True(t, isBannerRepository)
	assert.Equal(t, db, bannerRepository.db)
	assert.Equal(t, dbPaginator, bannerRepository.dbPaginator)

	bannerId := models.NewModelId()
	commonReply := []map[string]interface{}{{
		"id": bannerId.String(),
	}}

	t.Run("ListBanners", func(t *testing.T) {
		bannerPaginationQuery := &models.BannerPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), bannerPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(nil)

		paginationResult, err := bannerRepository.ListBanners(bannerPaginationQuery)
		assert.IsType(t, []*entities.BannerEntity{}, paginationResult.Data)
		assert.Nil(t, err)
	})

	t.Run("ListBanners:Error", func(t *testing.T) {
		systemErr := errors.NewUnknownError()
		bannerPaginationQuery := &models.BannerPaginationQuery{
			PaginationQuery: &models.PaginationQuery{Limit: 20},
		}
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL
		dbPaginator.EXPECT().Paginate(
			gomock.Any(), bannerPaginationQuery.PaginationQuery, gomock.Any(), gomock.Any(),
		).Return(systemErr)

		bannerEntities, err := bannerRepository.ListBanners(bannerPaginationQuery)
		assert.Nil(t, bannerEntities)
		assert.Equal(t, systemErr, err)
	})

	t.Run("GetBanner", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery("SELECT *").WithReply(commonReply)

		bannerEntity, err := bannerRepository.GetBanner(bannerId)
		assert.IsType(t, new(entities.BannerEntity), bannerEntity)
		assert.Nil(t, err)
	})

	t.Run("GetBanner:NotFoundError", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrRecordNotFound

		bannerEntity, err := bannerRepository.GetBanner(bannerId)
		assert.NotNil(t, bannerEntity)
		assert.Error(t, err)
	})

	t.Run("GetBanner:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		bannerEntity, err := bannerRepository.GetBanner(bannerId)
		assert.NotNil(t, bannerEntity)
		assert.Error(t, err, errors.NewSystemErrorFromBuiltin(gorm.ErrInvalidSQL))
	})

	t.Run("SaveBanner", func(t *testing.T) {
		mocket.Catcher.Reset()

		bannerEntity := entities.NewBannerEntity()
		assert.Nil(t, bannerRepository.SaveBanner(bannerEntity))
	})

	t.Run("SaveBanner:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		bannerEntity := new(entities.BannerEntity)
		assert.Error(t, bannerRepository.SaveBanner(bannerEntity))
	})

	t.Run("RemoveBanner", func(t *testing.T) {
		mocket.Catcher.Reset()

		bannerEntity := new(entities.BannerEntity)
		assert.Nil(t, bannerRepository.RemoveBanner(bannerEntity))
	})

	t.Run("RemoveBanner:Error", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().Error = gorm.ErrInvalidSQL

		bannerEntity := new(entities.BannerEntity)
		assert.Error(t, bannerRepository.RemoveBanner(bannerEntity))
	})
}
