package services

import (
	"github.com/emelnychenko/go-press/contracts"
	"github.com/emelnychenko/go-press/entities"
	"github.com/emelnychenko/go-press/errors"
	"github.com/emelnychenko/go-press/models"
	"time"
)

type (
	bannerServiceImpl struct {
		bannerEntityFactory contracts.BannerEntityFactory
		bannerRepository    contracts.BannerRepository
	}
)

func NewBannerService(
	bannerEntityFactory contracts.BannerEntityFactory,
	bannerRepository contracts.BannerRepository,
) (bannerService contracts.BannerService) {
	return &bannerServiceImpl{
		bannerEntityFactory,
		bannerRepository,
	}
}

func (s *bannerServiceImpl) ListBanners(
	bannerPaginationQuery *models.BannerPaginationQuery,
) (*models.PaginationResult, errors.Error) {
	return s.bannerRepository.ListBanners(bannerPaginationQuery)
}

func (s *bannerServiceImpl) GetBanner(bannerId *models.BannerId) (*entities.BannerEntity, errors.Error) {
	return s.bannerRepository.GetBanner(bannerId)
}

func (s *bannerServiceImpl) CreateBanner(data *models.BannerCreate) (
	bannerEntity *entities.BannerEntity, err errors.Error,
) {
	bannerEntity = s.bannerEntityFactory.CreateBannerEntity()
	bannerEntity.Title = data.Title
	bannerEntity.Key = data.Key

	err = s.bannerRepository.SaveBanner(bannerEntity)
	return
}

func (s *bannerServiceImpl) UpdateBanner(bannerEntity *entities.BannerEntity, data *models.BannerUpdate) errors.Error {
	bannerEntity.Title = data.Title
	bannerEntity.Key = data.Key

	updated := time.Now().UTC()
	bannerEntity.Updated = &updated

	return s.bannerRepository.SaveBanner(bannerEntity)
}

func (s *bannerServiceImpl) DeleteBanner(bannerEntity *entities.BannerEntity) errors.Error {
	return s.bannerRepository.RemoveBanner(bannerEntity)
}
